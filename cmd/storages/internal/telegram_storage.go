package internal

import (
	"context"
	"github.com/stas727/gbot/cmd/storages/model"
)

type TelegramImpl struct {
}

func (c *TelegramImpl) GetTelegramModel(ctx context.Context) *model.Telegram {
	telegram := &model.Telegram{}

	return telegram
}

func NewTelegramStorage() *TelegramImpl {
	return &TelegramImpl{}
}
