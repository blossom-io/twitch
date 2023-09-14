package service

import (
	"blossom/internal/infrastructure/gpt"
	"blossom/pkg/ffmpeg"
	"blossom/pkg/imgur"
	"blossom/pkg/logger"
	"blossom/pkg/twitch"
	"blossom/pkg/youtube"
)

//go:generate mockery --name Servicer
type Servicer interface {
	AIer
	Screenshot(channel string) (imgURL string, err error)
	PreviewLink(URL string) (description string, linkType Link, err error)
}

type service struct {
	log     logger.Logger
	FFMpeg  ffmpeg.FFMpeger
	Imgur   imgur.Imgurer
	Twitch  twitch.Twitcher
	Youtube youtube.Youtuber
	gpt     gpt.GPTer
}

func New(log logger.Logger, ffmpeg ffmpeg.FFMpeger, imgur imgur.Imgurer, twitch twitch.Twitcher, youtube youtube.Youtuber, gpt gpt.GPTer) Servicer {
	return &service{
		log:     log,
		FFMpeg:  ffmpeg,
		Imgur:   imgur,
		Twitch:  twitch,
		Youtube: youtube,
		gpt:     gpt,
	}
}
