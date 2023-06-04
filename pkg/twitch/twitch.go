package twitch

import (
	"context"

	"blossom/pkg/logger"

	"github.com/wmw64/twitchpl"
)

//go:generate mockery --name Twitcher
type Twitcher interface {
	GetSource(channel string) (source string, err error)
}

type twitch struct {
	log logger.Logger
}

type Playlist struct {
	Name       string
	Resolution string
	URL        string
	FrameRate  float64
}

func New(log logger.Logger) Twitcher {
	return &twitch{log: log}
}

func (t *twitch) GetSource(channel string) (source string, err error) {
	pl, err := twitchpl.Get(context.TODO(), channel, true)
	if err != nil {
		return "", err
	}

	return pl.AsURL(), err
}
