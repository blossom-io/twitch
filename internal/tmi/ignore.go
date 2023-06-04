package tmi

import "github.com/gempir/go-twitch-irc/v4"

// IgnoreMsg checks if the user of the message is in the ignore list.
// It is good idea to ignore other bots.
func (c *chat) IgnoreMsg(msg twitch.PrivateMessage) bool {
	if _, ok := c.Ignore[msg.User.Name]; ok {
		return true
	}

	return false
}
