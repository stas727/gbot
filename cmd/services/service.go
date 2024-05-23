package service

import (
	"context"
	"github.com/stas727/gbot/cmd/services/internal"
	"github.com/stas727/gbot/cmd/storages"
	"github.com/stas727/gbot/cmd/storages/model"
	"time"
)

type Services struct {
	AIService  IAIService
	BotService IBotService
}

type IAIService interface {
	NewClient(ctx context.Context, token string) *model.OpenAI
	Response(ctx context.Context, request string, model *model.OpenAI) (*string, error)
}

type IBotService interface {
	NewBot(ctx context.Context, token string, timeout time.Duration, url string) (*model.Telegram, error)
	Handle(ctx context.Context, handler func(message string) string, model *model.Telegram) error
	Pmetrics(ctx context.Context, payload string)
}

func NewServices(storages *storages.Storages) *Services {
	return &Services{
		AIService:  internal.NewAIService(storages.OpenAIStorage),
		BotService: internal.NewBotService(storages.TelegramStorage),
	}
}
