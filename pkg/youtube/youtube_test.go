package youtube

import (
	_ "embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed test/TestJsonNoEmbed.json
var TestJsonNoEmbed []byte

func TestGetVideoInfo(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/embed" {
			w.Header().Set("Content-Type", "application/json")
			w.Write(TestJsonNoEmbed)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	// Replace NoEmbedURL with the test server URL
	NoEmbedURL = ts.URL + "/embed?url=%s"

	c := New()

	// Test with a valid URL
	videoInfo, err := c.GetVideoInfo("https://youtu.be/J4t4pMZBXZg")
	require.NoError(t, err)
	assert.Equal(t, "Skeler - N i g h t D r i v e スケラー PART II", videoInfo.Title)
}
