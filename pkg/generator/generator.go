package generator

import "fmt"

func NewCooldownKey(channel, cmd string) string {
	return fmt.Sprint(channel, ":", cmd)
}
