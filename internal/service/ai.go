package service

import (
	"blossom/internal/entity"
	"context"
)

type AIer interface {
	Ask(ctx context.Context, prompt string) (answer string, err error)
	AskStream(ctx context.Context, prompt string) (delta chan entity.Delta, err error)
}

func (svc *service) Ask(ctx context.Context, prompt string) (answer string, err error) {
	return svc.gpt.Ask(ctx, prompt)
}

func (svc *service) AskStream(ctx context.Context, prompt string) (delta chan entity.Delta, err error) {
	delta = make(chan entity.Delta, 100)

	go func() {
		stream, err := svc.gpt.AskStream(ctx, prompt)
		if err != nil {
			return
		}
		defer stream.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				res, err := stream.Recv()
				if err != nil {
					delta <- entity.Delta{Content: "", Err: err}

					return
				}

				delta <- entity.Delta{Content: res.Choices[0].Delta.Content, Err: nil}
			}
		}

	}()

	return delta, nil
}
