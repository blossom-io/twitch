package service

import (
	"testing"

	"blossom/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestScreenshot(t *testing.T) {
	logMock := mocks.NewLogger(t)
	twitchMock := mocks.NewTwitcher(t)
	ffmpegMock := mocks.NewFFMpeger(t)
	imgurMock := mocks.NewImgurer(t)
	youtubeMock := mocks.NewYoutuber(t)

	svc := New(logMock, ffmpegMock, imgurMock, twitchMock, youtubeMock)

	twitchMock.On("GetSource", "test-channel").Return("http://test.com/source", nil)
	ffmpegMock.On("Screenshot", "http://test.com/source").Return([]byte{1, 2, 3}, nil)
	imgurMock.On("UploadImage", []byte{1, 2, 3}).Return("http://test.com/image", nil)

	URL, err := svc.Screenshot("test-channel")

	assert.Equal(t, "http://test.com/image", URL)
	assert.Nil(t, err)

	twitchMock.AssertCalled(t, "GetSource", "test-channel")
	ffmpegMock.AssertCalled(t, "Screenshot", "http://test.com/source")
	imgurMock.AssertCalled(t, "UploadImage", []byte{1, 2, 3})
}
