package tmi

import (
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
)

// IgnoreMsg checks if the user of the message is in the ignore list.
// It is good idea to ignore other bots.
func (c *chat) IgnoreMsg(msg twitch.PrivateMessage) bool {
	if _, ok := c.IgnoreChannels[msg.User.Name]; ok {
		return true
	}

	return false
}

func (c *chat) IsCommandEnabled(msg twitch.PrivateMessage) bool {
	cmd := strings.Split(msg.Message, " ")[0]

	if _, ok := c.CommandsEnabled[cmd]; ok {
		return true
	}

	return false
}
