package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractLink(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		shouldExist bool
	}{
		{
			name:        "extracts http link",
			input:       "Hello world https://google.com",
			expected:    "https://google.com",
			shouldExist: true,
		},
		{
			name:        "extracts https link",
			input:       "Hello world https://github.com",
			expected:    "https://github.com",
			shouldExist: true,
		},
		{
			name:        "does not extract non-link text",
			input:       "Hello world",
			expected:    "",
			shouldExist: false,
		},
		{
			name:        "does not extract invalid link",
			input:       "Hello world https://",
			expected:    "",
			shouldExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link, found := ExtractLink(tt.input)

			assert.Equal(t, tt.expected, link)
			assert.Equal(t, tt.shouldExist, found)
		})
	}
}
