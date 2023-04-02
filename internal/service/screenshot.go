package service

// Screenshot takes a screenshot of a twitch stream, uploads it to imgur, and returns the direct image URL.
func (s *service) Screenshot(channel string) (ImageURL string, err error) {
	sourceURL, err := s.Twitch.GetSource(channel)
	if err != nil {
		s.log.Error("service - Screenshot - Twitch.GetSource: %w", err)
		return "", err
	}

	img, err := s.FFMpeg.Screenshot(sourceURL)
	if err != nil {
		s.log.Error("service - Screenshot - FFMpeg.GetScreenshot: %w", err)
	}

	ImageURL, err = s.Imgur.UploadImage(img)
	if err != nil {
		s.log.Error("service - Screenshot - Imgur.UploadImage: %w", err)
	}

	return ImageURL, err
}
