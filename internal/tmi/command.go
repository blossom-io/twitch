package tmi

import (
	"errors"
	"fmt"

	"blossom/pkg/link"

	"github.com/gempir/go-twitch-irc/v4"
)

func (c *chat) CommandPing(msg twitch.PrivateMessage) (ok bool) {
	if msg.Message == "!ping" {
		if onCooldown := c.IsCooldown(msg.Channel, CmdSS); onCooldown {
			return true
		}

		c.TMI.Reply(msg.Channel, msg.ID, "pong!")
		return true
	}

	return false
}

func (c *chat) CommandScreenshot(msg twitch.PrivateMessage) (ok bool) {
	if msg.Message == "!ss" || msg.Message == "!скриншот" {
		if onCooldown := c.IsCooldown(msg.Channel, CmdSS); onCooldown {
			return true
		}

		imgURL, err := c.svc.Screenshot(msg.Channel)
		if errors.Is(errors.New("stream is offline or channel not found"), err) {
			c.log.Debug("service - CommandScreenshot: %s", err)
			c.TMI.Reply(msg.Channel, msg.ID, "stream is offline")
			return true
		}

		if err != nil {
			c.log.Error("service - CommandScreenshot: %w", err)
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

	if onCooldown := c.IsCooldown(msg.Channel, CmdPreviewLink); onCooldown {
		return true
	}

	description, linkType, err := c.svc.PreviewLink(URL)
	if err != nil {
		c.log.Error("tmi - CommandPreviewLink: %w", err)

		return false
	}

	m := fmt.Sprint(linkType, ": ", description)
	c.TMI.Reply(msg.Channel, msg.ID, m)

	return true
}

func (c *chat) CommandInvite(msg twitch.PrivateMessage) (ok bool) {
	if onCooldown := c.IsCooldown(msg.Channel, CmdInvite); onCooldown {
		return true
	}

	if msg.Message == "!invite" {
		c.TMI.Reply(msg.Channel, msg.ID, "Просто добавь воды!")
		return true
	}

	return false
}
