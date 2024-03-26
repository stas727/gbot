package internal

import (
	"context"
	"github.com/stas727/gbot/cmd/storages"
	"github.com/stas727/gbot/cmd/storages/model"
	"gopkg.in/telebot.v3"
	"time"
)

type BotImpl struct {
	TelegramStorage storages.ITelegramStorage
}

func (c *BotImpl) NewBot(ctx context.Context, token string, timeout time.Duration, url string) (*model.Telegram, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		URL:    url,
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: timeout},
	})

	if err != nil {

		return nil, err
	}

	telegramModel := c.TelegramStorage.GetTelegramModel(ctx)
	telegramModel.Client = bot

	return telegramModel, nil
}

func (c *BotImpl) Handle(ctx context.Context, handler func(message string) string, model *model.Telegram) error {
	model.Client.Handle(telebot.OnText, func(m telebot.Context) error {
		message := handler(m.Text())

		err := m.Send(message)

		if err != nil {

			return err
		}

		return nil
	})

	return nil
}

func NewBotService(TelegramStorage storages.ITelegramStorage) *BotImpl {
	return &BotImpl{
		TelegramStorage: TelegramStorage,
	}
}
