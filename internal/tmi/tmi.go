package tmi

import (
	"blossom/internal/config"
	"blossom/internal/service"
	"blossom/pkg/logger"

	"github.com/gempir/go-twitch-irc/v4"
)

type chat struct {
	log logger.Logger
	Cfg *config.Config
	TMI *twitch.Client // Twitch Messaging Interface
	svc service.Servicer
}

func New(log logger.Logger, cfg *config.Config, svc service.Servicer, channels ...string) *chat {
	chat := &chat{
		log: log,
		TMI: twitch.NewClient(cfg.Name, cfg.OAuth),
		svc: svc,
	}

	chat.Commands()

	chat.Join(channels...)

	chat.Connect()

	return chat
}

func (c *chat) Join(channels ...string) {
	c.TMI.Join(channels...)
}

func (c *chat) Connect() {
	err := c.TMI.Connect()
	if err != nil {
		c.log.Error("chat - Connect: %w", err)
		panic(err)
	}
}

func (c *chat) Say(channel, msg string) {
	c.TMI.Say(channel, msg)
}

func (c *chat) Close() {
	c.TMI.Disconnect()
}

func (c *chat) Commands() {
	c.TMI.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		if ok := c.CommandPing(msg); ok {
			return
		}
		if ok := c.CommandScreenshot(msg); ok {
			return
		}
		if ok := c.CommandPreviewLink(msg); ok {
			return
		}
		if ok := c.CommandInvite(msg); ok {
			return
		}
	})
}
