package youtube

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var NoEmbedURL = "https://noembed.com/embed?url=%s" // see: https://noembed.com/embed?url=https://youtu.be/J4t4pMZBXZg

type VideoInfo struct {
	AuthorURL       string `json:"author_url"`
	ProviderName    string `json:"provider_name"`
	Type            string `json:"type"`
	HTML            string `json:"html"`
	AuthorName      string `json:"author_name"`
	ThumbnailURL    string `json:"thumbnail_url"`
	Error           string `json:"error"`
	URL             string `json:"url"`
	ProviderURL     string `json:"provider_url"`
	Version         string `json:"version"`
	Title           string `json:"title"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	Width           int    `json:"width"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	Height          int    `json:"height"`
}

//go:generate mockery --name Youtuber
type Youtuber interface {
	GetVideoTitle(URL string) (string, error)
	GetVideoInfo(URL string) (videoInfo VideoInfo, err error)
}

type Client struct {
	httpClient *http.Client
}

func New() Youtuber {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) GetVideoInfo(URL string) (videoInfo VideoInfo, err error) {
	embedURL := fmt.Sprintf(NoEmbedURL, URL)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, embedURL, http.NoBody)
	if err != nil {
		return videoInfo, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return videoInfo, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return videoInfo, errors.New("error getting video info")
	}

	if err := json.NewDecoder(res.Body).Decode(&videoInfo); err != nil {
		return videoInfo, err
	}

	if videoInfo.Error != "" {
		return videoInfo, errors.New(videoInfo.Error)
	}

	return videoInfo, nil
}

func (c *Client) GetVideoTitle(URL string) (string, error) {
	videoInfo, err := c.GetVideoInfo(URL)
	if err != nil {
		return "", err
	}

	return videoInfo.Title, err
}
