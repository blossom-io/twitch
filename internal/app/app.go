package app

import (
	"context"
	"os"
	"os/signal"

	"blossom/internal/config"
	"blossom/internal/infrastructure/gpt"
	"blossom/internal/service"
	"blossom/internal/tmi"
	"blossom/pkg/ffmpeg"
	"blossom/pkg/imgur"
	"blossom/pkg/logger"
	"blossom/pkg/twitch"
	"blossom/pkg/youtube"
)

// Run injects dependencies and runs application.
func Run(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log := logger.New()
	log.Debug("starting bot...")

	log.Info("channels to connect", cfg.Channel)

	svc := service.New(log, cfg, ffmpeg.New(), imgur.New(), twitch.New(log), youtube.New(), gpt.New(cfg, log))

	chat := tmi.New(log, cfg, svc, cfg.Channel...)
	defer chat.Close()

	<-ctx.Done()
	log.Info("Gracefully shutting down...")
}
