package service

import (
	"net/url"
)

type Link string

const (
	LinkUnknown Link = "Unknown"
	LinkYouTube Link = "YouTube"
)

func (s *service) PreviewLink(URL string) (description string, linkType Link, err error) {
	switch GetLinkType(URL) {
	case LinkYouTube:
		description, err = s.Youtube.GetVideoTitle(URL)
		if err != nil {
			s.log.Error("service - PreviewLink - youtube.GetVideoTitle: %w", err)
		}

		return description, LinkYouTube, err

	case LinkUnknown:
		s.log.Debug("service - PreviewLink - LinkUnknown")
	}

	return description, LinkUnknown, err
}

// GetLinkType gets the type of link. Example: https://www.youtube.com/watch?v=QH2-TGUlwu4 -> LinkYouTube
func GetLinkType(URL string) Link {
	if IsYouTubeLink(URL) {
		return LinkYouTube
	}

	return LinkUnknown
}

// IsYouTubeLink checks if the link is a YouTube link.
func IsYouTubeLink(URL string) bool {
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return false
	}

	switch parsedURL.Host {
	case "youtu.be", "www.youtu.be", "youtube.com", "www.youtube.com":
		return true
	default:
		return false
	}
}
