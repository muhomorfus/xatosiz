package technical

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/consumers"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/kafka"
	"github.com/Shopify/sarama"
	"golang.org/x/sync/errgroup"
	"os"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/managers"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/menu"
	repository "git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/postgres"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	ErrEmptyFilename = errors.New("empty filename")
)

type App struct {
	cfg    Config
	logger *zap.SugaredLogger

	menu *menu.Menu

	alertCreationDBConsumer *consumers.AlertCreationConsumer
	alertProcessingConsumer *consumers.AlertProcessingConsumer

	alertCreationConsumerGroup   sarama.ConsumerGroup
	alertProcessingConsumerGroup sarama.ConsumerGroup
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

	producer, err := a.newAsyncProducer()
	if err != nil {
		a.logger.Fatalw("cant create async producer", "error", err)

		return fmt.Errorf("new async producer: %w", err)
	}

	a.alertCreationConsumerGroup, err = a.newConsumerGroup("alert_creation_group")
	if err != nil {
		a.logger.Fatalw("cant create consumer group", "error", err)

		return fmt.Errorf("new consumer group: %w", err)
	}

	a.alertProcessingConsumerGroup, err = a.newConsumerGroup("alert_processing_group")
	if err != nil {
		a.logger.Fatalw("cant create consumer group", "error", err)

		return fmt.Errorf("new consumer group: %w", err)
	}

	eventRepo := repository.NewEventRepository(db)
	groupRepo := repository.NewGroupRepository(db)
	traceRepo := repository.NewTraceRepository(db)

	alertRepo := repository.NewAlertRepository(db)
	alertHitRepo := repository.NewAlertHitRepository(db)
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

	a.menu = menu.New(acceptMng, showMng, alertMng, alertCfgMng)

	return nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Infow("starting application")

	var eg errgroup.Group

	eg.Go(func() error {
		if err := a.menu.Run(contextutils.WithLogger(ctx, a.logger)); err != nil {
			a.logger.Errorw("error while running application", "error", err)

			return fmt.Errorf("run app: %w", err)
		}

		return nil
	})

	eg.Go(func() error {
		for {
			if err := a.alertCreationConsumerGroup.Consume(ctx, []string{"alerts"}, a.alertCreationDBConsumer); err != nil {
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

	return eg.Wait()
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

func (a *App) newAsyncProducer() (sarama.SyncProducer, error) {
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
