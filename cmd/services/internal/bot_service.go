package internal

import (
	"context"
	"fmt"
	"github.com/stas727/gbot/cmd/storages"
	"github.com/stas727/gbot/cmd/storages/model"
	"go.opentelemetry.io/otel"
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

func (c *BotImpl) Pmetrics(ctx context.Context, payload string) {
	// Get the global MeterProvider and create a new Meter with the name "gbot_light_signal_counter"
	meter := otel.GetMeterProvider().Meter("gbot_light_signal_counter")

	// Get or create an Int64Counter instrument with the name "gbot_light_signal_<payload>"
	counter, _ := meter.Int64Counter(fmt.Sprintf("gbot_light_signal_%s", payload))

	// Add a value of 1 to the Int64Counter
	counter.Add(ctx, 1)
}

func (c *BotImpl) Handle(ctx context.Context, handler func(message string) string, model *model.Telegram) error {
	model.Client.Handle(telebot.OnText, func(m telebot.Context) error {

		payload := m.Message().Payload
		c.Pmetrics(context.Background(), payload)
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
