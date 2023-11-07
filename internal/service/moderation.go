package service

import "strings"

type Moderator interface {
	ContainsBannedWords(text string) bool
}

func (svc *service) ContainsBannedWords(text string) bool {
	text = strings.ToLower(text)

	for _, word := range svc.cfg.BannedWords {
		if strings.Contains(text, word) {
			return true
		}
	}

	return false
}
