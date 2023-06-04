package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsYouTubeLink(t *testing.T) {
	tests := []struct {
		name     string
		link     string
		expected bool
	}{
		{
			name:     "YouTube link with www",
			link:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			expected: true,
		},
		{
			name:     "YouTube link without www",
			link:     "https://youtube.com/watch?v=dQw4w9WgXcQ",
			expected: true,
		},
		{
			name:     "Shortened YouTube link with www",
			link:     "https://www.youtu.be/dQw4w9WgXcQ",
			expected: true,
		},
		{
			name:     "Shortened YouTube link without www",
			link:     "https://youtu.be/dQw4w9WgXcQ",
			expected: true,
		},
		{
			name:     "Non-YouTube link",
			link:     "https://example.com",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsYouTubeLink(tt.link)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetLinkType(t *testing.T) {
	tests := []struct {
		name     string
		link     string
		expected Link
	}{
		{
			name:     "YouTube link with www",
			link:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			expected: LinkYouTube,
		},
		{
			name:     "YouTube link without www",
			link:     "https://youtube.com/watch?v=dQw4w9WgXcQ",
			expected: LinkYouTube,
		},
		{
			name:     "Shortened YouTube link with www",
			link:     "https://www.youtu.be/dQw4w9WgXcQ",
			expected: LinkYouTube,
		},
		{
			name:     "Shortened YouTube link without www",
			link:     "https://youtu.be/dQw4w9WgXcQ",
			expected: LinkYouTube,
		},
		{
			name:     "Non-YouTube link",
			link:     "https://example.com",
			expected: LinkUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetLinkType(tt.link)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

// func TestPreviewLink(t *testing.T) {
// 	logMock := mocks.NewLogger(t)
// 	twitchMock := mocks.NewTwitcher(t)
// 	ffmpegMock := mocks.NewFFMpeger(t)
// 	imgurMock := mocks.NewImgurer(t)
// 	youtubeMock := mocks.NewYoutuber(t)

// 	svc := New(logMock, ffmpegMock, imgurMock, twitchMock, youtubeMock)

// 	twitchMock.On("GetSource", "test-channel").Return("http://test.com/source", nil)
// 	ffmpegMock.On("Screenshot", "http://test.com/source").Return([]byte{1, 2, 3}, nil)
// 	imgurMock.On("UploadImage", []byte{1, 2, 3}).Return("http://test.com/image", nil)

// 	URL, err := svc.Screenshot("test-channel")

// 	assert.Equal(t, "http://test.com/image", URL)
// 	assert.Nil(t, err)

// 	twitchMock.AssertCalled(t, "GetSource", "test-channel")
// 	ffmpegMock.AssertCalled(t, "Screenshot", "http://test.com/source")
// 	imgurMock.AssertCalled(t, "UploadImage", []byte{1, 2, 3})
// }
