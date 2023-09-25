package web

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/consumers"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/clickhouse"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/kafka"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/redis"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/telegram"
	"github.com/Shopify/sarama"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	redisClient "github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/openapi"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/server"
	"github.com/rs/cors"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/managers"
	repository "git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/postgres"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	clickhouseDriver "gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	ErrEmptyFilename = errors.New("empty filename")
	ErrInvalidDB     = errors.New("invalid database type")
)

type App struct {
	cfg     Config
	logger  *zap.SugaredLogger
	srv     *server.Server
	handler http.Handler

	alertCreationDBConsumer       *consumers.AlertCreationConsumer
	alertCreationTelegramConsumer *consumers.AlertCreationConsumer
	alertProcessingConsumer       *consumers.AlertProcessingConsumer

	alertCreationDBConsumerGroup       sarama.ConsumerGroup
	alertCreationTelegramConsumerGroup sarama.ConsumerGroup
	alertProcessingConsumerGroup       sarama.ConsumerGroup
}

func New() *App {
	return &App{}
}

func (a *App) Init() error {
	if err := a.readConfig(); err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	if err := a.initLogger(); err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	db, err := gorm.Open(postgres.Open(a.cfg.Postgres.DSN()), &gorm.Config{
		Logger: zapgorm2.New(a.logger.Desugar()),
	})
	if err != nil {
		a.logger.Fatalw("cant open db conn", "error", err)

		return fmt.Errorf("open gorm db: %w", err)
	}

	producer, err := a.newSyncProducer()
	if err != nil {
		a.logger.Fatalw("cant create async producer", "error", err)

		return fmt.Errorf("new async producer: %w", err)
	}

	a.alertCreationDBConsumerGroup, err = a.newConsumerGroup("alert_creation_group")
	if err != nil {
		a.logger.Fatalw("cant create consumer group", "error", err)

		return fmt.Errorf("new consumer group: %w", err)
	}

	a.alertProcessingConsumerGroup, err = a.newConsumerGroup("alert_processing_group")
	if err != nil {
		a.logger.Fatalw("cant create consumer group", "error", err)

		return fmt.Errorf("new consumer group: %w", err)
	}

	var (
		eventRepo    ports.EventRepository
		groupRepo    ports.GroupRepository
		traceRepo    ports.TraceRepository
		alertHitRepo ports.AlertHitRepository
		alertRepo    ports.AlertRepository
	)

	if a.cfg.TracesDB == TracesDBPostgres {
		eventRepo = repository.NewEventRepository(db)
		groupRepo = repository.NewGroupRepository(db)
		traceRepo = repository.NewTraceRepository(db)
		alertHitRepo = repository.NewAlertHitRepository(db)
		alertRepo = repository.NewAlertRepository(db)
	} else if a.cfg.TracesDB == TracesDBClickhouse {
		chDB, err := gorm.Open(clickhouseDriver.Open(a.cfg.Clickhouse.DSN()), &gorm.Config{
			Logger: zapgorm2.New(a.logger.Desugar()),
		})
		if err != nil {
			a.logger.Fatalw("cant open clickhouse db conn", "error", err)

			return fmt.Errorf("open gorm db: %w", err)
		}

		eventRepo = clickhouse.NewEventRepository(chDB, producer)
		groupRepo = clickhouse.NewGroupRepository(chDB, producer)
		traceRepo = clickhouse.NewTraceRepository(chDB, producer)
		alertHitRepo = clickhouse.NewAlertHitRepository(chDB, producer)

		rDB := redisClient.NewClient(&redisClient.Options{
			Addr:     a.cfg.Redis.Address,
			Username: a.cfg.Redis.User,
			Password: a.cfg.Redis.Password,
			DB:       a.cfg.Redis.DB,
		})

		ttl := time.Duration(a.cfg.Redis.TTL) * time.Millisecond

		groupRepo = redis.NewGroupRepository(rDB, groupRepo, ttl)
		traceRepo = redis.NewTraceRepository(rDB, traceRepo, ttl)
		alertHitRepo = redis.NewAlertHitRepository(rDB, alertHitRepo, ttl)
		alertRepo = redis.NewAlertRepository(rDB)
	} else {
		return fmt.Errorf("check db: %w", ErrInvalidDB)
	}

	alertConfigRepo := repository.NewAlertConfigRepository(db)

	alertCreationRepo := kafka.NewAlertCreationRepository(producer)
	alertProcessingQueue := kafka.NewAlertProcessingQueue(producer)

	alertMng := managers.NewAlertManager(alertRepo)
	alertCfgMng := managers.NewAlertConfigManager(alertConfigRepo)

	alertSaveDBMng := managers.NewAlertSaveManager(alertRepo)

	procEventSyncMng := managers.NewProcessEventSyncManager(alertHitRepo, alertConfigRepo, alertCreationRepo)
	procEventAsyncMng := managers.NewProcessEventAsyncManager(alertProcessingQueue)

	acceptMng := managers.NewAcceptManager(groupRepo, traceRepo, eventRepo, procEventAsyncMng)
	showMng := managers.NewShowManager(groupRepo, traceRepo, eventRepo)

	a.alertCreationDBConsumer = consumers.NewAlertCreationConsumer(a.logger, alertSaveDBMng)
	a.alertProcessingConsumer = consumers.NewAlertProcessingConsumer(a.logger, procEventSyncMng)

	if err := a.initTelegram(); err != nil {
		a.logger.Fatalw("cannot init telegram", "error", err)

		return fmt.Errorf("init telegram: %w", err)
	}

	a.srv = server.New(acceptMng, showMng, alertMng, alertCfgMng)

	a.handler = openapi.NewRouter(openapi.NewDefaultApiController(a.srv))
	a.handler = a.middleware(a.handler)
	a.handler = cors.AllowAll().Handler(a.handler)

	return nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Infow("starting application")

	var eg errgroup.Group

	eg.Go(func() error {
		return http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.Port), a.handler)
	})

	eg.Go(func() error {
		for {
			if err := a.alertCreationDBConsumerGroup.Consume(ctx, []string{"alerts"}, a.alertCreationDBConsumer); err != nil {
				return fmt.Errorf("consume: %w", err)
			}

			if ctx.Err() != nil {
				return nil
			}
		}
	})

	eg.Go(func() error {
		for {
			if err := a.alertProcessingConsumerGroup.Consume(ctx, []string{"alert_processing"}, a.alertProcessingConsumer); err != nil {
				return fmt.Errorf("consume: %w", err)
			}

			if ctx.Err() != nil {
				return nil
			}
		}
	})

	if a.cfg.Telegram.Enabled {
		eg.Go(func() error {
			for {
				if err := a.alertCreationTelegramConsumerGroup.Consume(ctx, []string{"alerts"}, a.alertCreationTelegramConsumer); err != nil {
					return fmt.Errorf("consume: %w", err)
				}

				if ctx.Err() != nil {
					return nil
				}
			}
		})
	}

	return eg.Wait()
}

func (a *App) initTelegram() error {
	if !a.cfg.Telegram.Enabled {
		return nil
	}

	var err error
	a.alertCreationTelegramConsumerGroup, err = a.newConsumerGroup("alert_creation_tg_group")
	if err != nil {
		return fmt.Errorf("new consumer group: %w", err)
	}

	bot, err := tgbotapi.NewBotAPI(a.cfg.Telegram.Token)
	if err != nil {
		return fmt.Errorf("init telegram: %w", err)
	}

	alertTelegramRepo := telegram.NewAlertCreationRepository(bot, a.cfg.Telegram.ChatID)
	alertSaveTelegramMng := managers.NewAlertSaveManager(alertTelegramRepo)
	a.alertCreationTelegramConsumer = consumers.NewAlertCreationConsumer(a.logger, alertSaveTelegramMng)

	return nil
}

func (a *App) readConfig() error {
	cfgFilename := flag.String("config", "", "path of config file")
	flag.Parse()

	if cfgFilename == nil {
		return fmt.Errorf("get config file: %w", ErrEmptyFilename)
	}

	cfgFile, err := os.Open(*cfgFilename)
	if err != nil {
		return fmt.Errorf("open config file: %w", err)
	}

	if err := json.NewDecoder(cfgFile).Decode(&a.cfg); err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	return nil
}

func (a *App) initLogger() error {
	level, err := zap.ParseAtomicLevel(a.cfg.Logger.Level)
	if err != nil {
		return fmt.Errorf("parse level: %w", err)
	}

	cfg := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "name",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeTime:    zapcore.RFC3339TimeEncoder,
			EncodeCaller:  zapcore.FullCallerEncoder,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
		},
		OutputPaths: []string{a.cfg.Logger.Path},
	}

	logger, err := cfg.Build()
	if err != nil {
		return fmt.Errorf("build config: %w", err)
	}

	a.logger = logger.Sugar()

	return nil
}

func (a *App) newSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(a.cfg.Kafka.Brokers, config)
	if err != nil {
		a.logger.Errorw("cant create async producer", "error", err)
		return nil, fmt.Errorf("new async producer: %w", err)
	}

	return producer, nil
}

func (a *App) newConsumerGroup(group string) (sarama.ConsumerGroup, error) {
	version, err := sarama.ParseKafkaVersion(a.cfg.Kafka.Version)
	if err != nil {
		return nil, fmt.Errorf("parse kafka version: %w", err)
	}

	config := sarama.NewConfig()
	config.Version = version

	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}

	client, err := sarama.NewConsumerGroup(a.cfg.Kafka.Brokers, group, config)
	if err != nil {
		return nil, fmt.Errorf("create consumer group: %w", err)
	}

	return client, nil
}
