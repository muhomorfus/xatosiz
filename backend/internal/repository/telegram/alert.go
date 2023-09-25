package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
)

type AlertCreationRepository struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewAlertCreationRepository(bot *tgbotapi.BotAPI, chatID int64) *AlertCreationRepository {
	return &AlertCreationRepository{bot: bot, chatID: chatID}
}

func (a *AlertCreationRepository) Create(ctx context.Context, alert *models.Alert) error {
	text := fmt.Sprintf("<b>%s</b>\nTime: %s\nCause: <a href=\"https://xatosiz.lownie.su/events/%s\">%s</a>",
		alert.Message,
		alert.Time.Format(time.RFC822),
		alert.Event.UUID,
		alert.Event.Message,
	)
	msg := tgbotapi.NewMessage(a.chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML

	if _, err := a.bot.Send(msg); err != nil {
		return fmt.Errorf("send tg message: %w", err)
	}

	return nil
}
