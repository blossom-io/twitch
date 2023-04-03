package tmi

import (
	"time"

	"blossom/pkg/generator"
)

var (
	CmdCdDefault     = time.Second * 2
	CmdCdPing        = time.Second * 10
	CmdCdSS          = time.Second * 5
	CmdCdInvite      = time.Second * 10
	CmdCdPreviewLink = time.Second * 3
)

// IsCooldown checks if the command is on cooldown for the channel.
func (c *chat) IsCooldown(channel string, cmd Cmd) (onCooldown bool) {
	key := generator.NewCooldownKey(channel, string(cmd))

	if _, ok := c.Cooldown[key]; ok {
		if time.Since(c.Cooldown[key]) < c.CooldownDuration(cmd) {
			return true
		}
	}

	c.SetCooldown(channel, cmd)

	return false
}

func (c *chat) SetCooldown(channel string, cmd Cmd) {
	key := generator.NewCooldownKey(channel, string(cmd))
	c.Cooldown[key] = time.Now()
}

// CooldownDuration returns the duration of the cooldown for the command.
func (c *chat) CooldownDuration(cmd Cmd) time.Duration {
	switch cmd {
	case CmdPreviewLink:
		return CmdCdDefault
	case CmdInvite:
		return CmdCdInvite
	case CmdPing:
		return CmdCdPing
	case CmdSS:
		return CmdCdSS
	default:
		return CmdCdDefault
	}
}
