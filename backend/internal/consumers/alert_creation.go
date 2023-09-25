package consumers

import (
	"encoding/json"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type AlertCreationConsumer struct {
	mng    ports.AlertSaveManager
	logger *zap.SugaredLogger
}

func NewAlertCreationConsumer(logger *zap.SugaredLogger, m ports.AlertSaveManager) *AlertCreationConsumer {
	return &AlertCreationConsumer{
		mng:    m,
		logger: logger,
	}
}

func (c *AlertCreationConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AlertCreationConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AlertCreationConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			c.logger.Debugw(
				"claimed message",
				"value", string(message.Value),
				"timestamp", message.Timestamp,
				"topic", message.Topic,
			)

			var a alert
			if err := json.Unmarshal(message.Value, &a); err != nil {
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

			am, err := toAlert(a)
			if err != nil {
				c.logger.Errorw(
					"cannot convert alert",
					"error", err,
					"value", string(message.Value),
					"timestamp", message.Timestamp,
					"topic", message.Topic,
				)

				session.MarkMessage(message, "")
				continue
			}

			if err := c.mng.Save(contextutils.WithLogger(session.Context(), c.logger), am); err != nil {
				c.logger.Errorw(
					"cannot save event",
					"error", err,
					"value", string(message.Value),
					"timestamp", message.Timestamp,
					"topic", message.Topic,
				)

				session.MarkMessage(message, "")
				continue
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
