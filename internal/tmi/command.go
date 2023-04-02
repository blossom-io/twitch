package tmi

import (
	"errors"

	"blossom/pkg/link"

	"github.com/gempir/go-twitch-irc/v4"
)

func (c *chat) CommandPing(msg twitch.PrivateMessage) (ok bool) {
	if msg.Message == "!ping" {
		c.TMI.Reply(msg.Channel, msg.ID, "pong!")
		return true
	}

	return false
}

func (c *chat) CommandScreenshot(msg twitch.PrivateMessage) (ok bool) {
	if msg.Message == "!ss" || msg.Message == "!скриншот" {
		imgURL, err := c.svc.Screenshot(msg.Channel)
		if errors.Is(errors.New("stream is offline or channel not found"), err) {
			c.log.Debug("service - !ss - Screenshot: %s", err)
			c.TMI.Reply(msg.Channel, msg.ID, "stream is offline")
			return true
		}

		if err != nil {
			c.log.Error("tmi - Commands: %w", err)
			return true
		}

		c.log.Debug("img upload success", imgURL)
		c.TMI.Reply(msg.Channel, msg.ID, imgURL)
		return true
	}

	return false
}

func (c *chat) CommandPreviewLink(msg twitch.PrivateMessage) (ok bool) {
	URL, found := link.ExtractLink(msg.Message)
	if !found {
		return false
	}

	description, err := c.svc.PreviewLink(URL)
	if err != nil {
		c.log.Error("tmi - Commands: %w", err)

		return false
	}

	c.TMI.Reply(msg.Channel, msg.ID, description)

	return true
}

func (c *chat) CommandInvite(msg twitch.PrivateMessage) (ok bool) {
	if msg.Message == "!invite" {
		c.TMI.Reply(msg.Channel, msg.ID, "Просто добавь воды!")
	}

	return true
}
