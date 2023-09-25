package consumers

import (
	"encoding/json"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type AlertProcessingConsumer struct {
	mng    ports.ProcessEventManager
	logger *zap.SugaredLogger
}

func NewAlertProcessingConsumer(logger *zap.SugaredLogger, m ports.ProcessEventManager) *AlertProcessingConsumer {
	return &AlertProcessingConsumer{
		mng:    m,
		logger: logger,
	}
}

func (c *AlertProcessingConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AlertProcessingConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AlertProcessingConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			c.logger.Debugw(
				"claimed message",
				"value", string(message.Value),
				"timestamp", message.Timestamp,
				"topic", message.Topic,
			)

			var e event
			if err := json.Unmarshal(message.Value, &e); err != nil {
				c.logger.Errorw(
					"cannot unmarshal message",
					"error", err,
					"value", string(message.Value),
					"timestamp", message.Timestamp,
					"topic", message.Topic,
				)

				session.MarkMessage(message, "")
				continue
			}

			em, err := toEvent(e)
			if err != nil {
				c.logger.Errorw(
					"cannot convert event",
					"error", err,
					"value", string(message.Value),
					"timestamp", message.Timestamp,
					"topic", message.Topic,
				)

				session.MarkMessage(message, "")
				continue
			}

			if err := c.mng.ProcessEvent(contextutils.WithLogger(session.Context(), c.logger), em); err != nil {
				c.logger.Errorw(
					"cannot process event",
					"error", err,
					"value", string(message.Value),
					"timestamp", message.Timestamp,
					"topic", message.Topic,
				)

				session.MarkMessage(message, "")
				continue
			}

			c.logger.Infow("successfully ended processing")

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
