package storages

import (
	"context"
	"github.com/stas727/gbot/cmd/storages/internal"
	"github.com/stas727/gbot/cmd/storages/model"
)

type Storages struct {
	OpenAIStorage   IOpenAIStorage
	TelegramStorage ITelegramStorage
}

type IOpenAIStorage interface {
	GetModel(ctx context.Context) *model.OpenAI
}

type ITelegramStorage interface {
	GetTelegramModel(ctx context.Context) *model.Telegram
}

func NewStorages() *Storages {
	return &Storages{
		OpenAIStorage:   internal.NewOpenAIStorage(),
		TelegramStorage: internal.NewTelegramStorage(),
	}
}
