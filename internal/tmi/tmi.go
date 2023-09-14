package tmi

import (
	"fmt"
	"time"

	"blossom/internal/config"
	"blossom/internal/service"
	"blossom/pkg/logger"

	"github.com/gempir/go-twitch-irc/v4"
)

type Cmd string

const (
	CmdDefault     Cmd = "default"
	CmdPing        Cmd = "!ping"
	CmdSS          Cmd = "!ss"
	CmdInvite      Cmd = "!invite"
	CmdPreviewLink Cmd = "previewlink"
	CmdGPT         Cmd = "!gpt"
)

type chat struct {
	log      logger.Logger
	Cfg      *config.Config
	TMI      *twitch.Client // Twitch Messaging Interface
	svc      service.Servicer
	Cooldown *cooldown
	// Cooldown map[string]time.Time
	Ignore map[string]struct{}
}

func New(log logger.Logger, cfg *config.Config, svc service.Servicer, channels ...string) *chat {
	chat := &chat{
		log:      log,
		Cfg:      cfg,
		TMI:      twitch.NewClient(cfg.Name, cfg.OAuth),
		svc:      svc,
		Cooldown: &cooldown{Cooldown: make(map[string]time.Time)},
		Ignore:   make(map[string]struct{}),
	}

	chat.FillIgnoreList()

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
		fmt.Println(msg.Message)
		if ignore := c.IgnoreMsg(msg); ignore {
			return
		}
		if ok := c.CommandPing(msg); ok {
			return
		}
		if ok := c.CommandScreenshot(msg); ok {
			return
		}
		if ok := c.CommandInvite(msg); ok {
			return
		}
		if ok := c.CommandGPT(msg); ok {
			return
		}
		if ok := c.CommandPreviewLink(msg); ok {
			return
		}
	})
}

func (c *chat) FillIgnoreList() {
	c.Ignore[c.Cfg.Name] = struct{}{} // ignore ourself

	for _, channel := range c.Cfg.Ignore {
		c.Ignore[channel] = struct{}{}
	}
}
