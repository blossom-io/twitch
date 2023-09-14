package gpt

import (
	"blossom/internal/config"
	"blossom/pkg/logger"
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type GPTer interface {
	Ask(ctx context.Context, prompt string) (string, error)
	AskStream(ctx context.Context, prompt string) (stream *openai.ChatCompletionStream, err error)
}

type GPT struct {
	log    logger.Logger
	cfg    *config.Config
	client *openai.Client
}

func New(cfg *config.Config, log logger.Logger) GPTer {
	return &GPT{
		cfg:    cfg,
		log:    log,
		client: openai.NewClient(cfg.AI.OpenAiApiKey),
	}
}

func (g *GPT) Ask(ctx context.Context, prompt string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: g.cfg.AI.MaxTokens,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			},
		},
	}

	res, err := g.client.CreateChatCompletion(ctx, req)
	if err != nil {
		g.log.Error("err", err)
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}

func (g *GPT) AskStream(ctx context.Context, prompt string) (stream *openai.ChatCompletionStream, err error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: g.cfg.AI.MaxTokens,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			},
		},
		Stream: true,
	}

	stream, err = g.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		g.log.Error("err", err)
		return nil, err
	}

	return stream, nil
}
