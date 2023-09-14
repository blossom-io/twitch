package tmi

import (
	"context"
	"fmt"
	"strings"

	"blossom/internal/service"
	"blossom/pkg/link"

	"github.com/gempir/go-twitch-irc/v4"
)

var msgInvite = "перейди по ссылке https://t.blsm.me/%s для приглашения в закрытый сабчат"

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

		go func() {
			imgURL, err := c.svc.Screenshot(msg.Channel)
			if err != nil {
				c.log.Error("service - CommandScreenshot: %w", err)
				return
			}

			c.log.Debug("img upload success", imgURL)
			c.TMI.Reply(msg.Channel, msg.ID, imgURL)
		}()

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
	if err != nil || linkType == service.LinkUnknown {
		c.log.Debug("tmi - CommandPreviewLink", description, linkType, err, msg)

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
		c.TMI.Reply(msg.Channel, msg.ID, fmt.Sprintf(msgInvite, msg.Channel))
		return true
	}

	return false
}

func (c *chat) CommandGPT(msg twitch.PrivateMessage) (ok bool) {
	if onCooldown := c.IsCooldown(msg.Channel, CmdGPT); onCooldown {
		return true
	}

	if after, ok := strings.CutPrefix(msg.Message, "!gpt"); ok {
		ctx, cancel := context.WithTimeout(context.Background(), c.Cfg.Bot.CmdTimeout)
		defer cancel()

		prompt := fmt.Sprintf("%s %s", after, c.Cfg.AI.CustomInstructions)

		answer, err := c.svc.Ask(ctx, prompt)
		if err != nil {
			c.log.Error(err.Error())

			return false
		}

		c.TMI.Reply(msg.Channel, msg.ID, answer)

		return true
	}

	return false
}
