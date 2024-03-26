package internal

import (
	"context"
	"github.com/stas727/gbot/cmd/storages/model"
)

type OpenAIImpl struct {
}

func (c *OpenAIImpl) GetModel(ctx context.Context) *model.OpenAI {
	openAIModel := &model.OpenAI{}

	return openAIModel
}

func NewOpenAIStorage() *OpenAIImpl {
	return &OpenAIImpl{}
}
