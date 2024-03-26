package internal

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/stas727/gbot/cmd/storages"
	"github.com/stas727/gbot/cmd/storages/model"
)

type AIImpl struct {
	OpenAIStorage storages.IOpenAIStorage
}

func (c *AIImpl) NewClient(ctx context.Context, token string) *model.OpenAI {
	clientAI := openai.NewClient(token)

	openAIModel := c.OpenAIStorage.GetModel(ctx)
	openAIModel.Client = clientAI

	return openAIModel
}

func (c *AIImpl) Response(ctx context.Context, request string, model *model.OpenAI) (*string, error) {
	resp, err := model.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: request,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message.Content, nil
}

func NewAIService(OpenAIStorage storages.IOpenAIStorage) *AIImpl {
	return &AIImpl{
		OpenAIStorage: OpenAIStorage,
	}
}
