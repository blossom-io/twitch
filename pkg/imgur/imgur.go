package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	apiEndpoint = "https://api.imgur.com/3/"
	clientID    = "Client-ID 678db8338a6d17b"
	// apiEndpointRapidAPI            = "https://imgur-apiv3.p.rapidapi.com/3/"
	// apiEndpointGenerateAccessToken = "https://api.imgur.com/oauth2/token"
)

type imageInfoDataWrapper struct {
	Ii      *ImageInfo `json:"data"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}

// ImageInfo contains all image information provided by imgur
type ImageInfo struct {
	Limit       *RateLimit // Current rate limit
	Name        string     `json:"name,omitempty"`       // OPTIONAL, the original filename, if you're logged in as the image owner
	Title       string     `json:"title"`                // The title of the image.
	Description string     `json:"description"`          // Description of the image.
	MimeType    string     `json:"type"`                 // Image MIME type.
	Vote        string     `json:"vote"`                 // The current user's vote on the album. null if not signed in, if the user hasn't voted on it, or if not submitted to the gallery.
	ID          string     `json:"id"`                   // The ID for the image
	Mp4         string     `json:"mp4,omitempty"`        // OPTIONAL, The direct link to the .mp4. Only available if the image is animated and type is 'image/gif'.
	Gifv        string     `json:"gifv,omitempty"`       // OPTIONAL, The .gifv link. Only available if the image is animated and type is 'image/gif'.
	Link        string     `json:"link"`                 // The direct link to the the image. (Note: if fetching an animated GIF that was over 20MB in original size, a .gif thumbnail will be returned)
	Section     string     `json:"section"`              // If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Deletehash  string     `json:"deletehash,omitempty"` // OPTIONAL, the deletehash, if you're logged in as the image owner
	Width       int        `json:"width"`                // The width of the image in pixels
	Bandwidth   int        `json:"bandwidth"`            // Bandwidth consumed by the image in bytes
	Views       int        `json:"views"`                // The number of image views
	Size        int        `json:"size"`                 // The size of the image in bytes
	Height      int        `json:"height"`               // The height of the image in pixels
	Mp4Size     int        `json:"mp4_size,omitempty"`   // OPTIONAL, The Content-Length of the .mp4. Only available if the image is animated and type is 'image/gif'. Note that a zero value (0) is possible if the video has not yet been generated
	Datetime    int        `json:"datetime"`             // Time uploaded, epoch time
	Looping     bool       `json:"looping,omitempty"`    // OPTIONAL, Whether the image has a looping animation. Only available if the image is animated and type is 'image/gif'.
	Favorite    bool       `json:"favorite"`             // Indicates if the current user favorited the image. Defaults to false if not signed in.
	Nsfw        bool       `json:"nsfw"`                 // Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Animated    bool       `json:"animated"`             // is the image animated
	InGallery   bool       `json:"in_gallery"`           // True if the image has been submitted to the gallery, false if otherwise.
}

type RateLimit struct {
	// Timestamp for when the credits will be reset.
	UserReset time.Time
	// Total credits that can be allocated.
	UserLimit int64
	// Total credits available.
	UserRemaining int64
	// Total credits that can be allocated for the application in a day.
	ClientLimit int64
	// Total credits remaining for the application in a day.
	ClientRemaining int64
}

//go:generate mockery --name Imgurer
type Imgurer interface {
	UploadImage(img []byte) (imgURL string, err error)
}

type Client struct {
	httpClient *http.Client
}

func New() Imgurer {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) UploadImage(img []byte) (imgURL string, err error) {
	URL := fmt.Sprint(apiEndpoint + "image")

	form := createUploadForm(img)

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBufferString(form.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", clientID)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var imgInfo imageInfoDataWrapper
	err = json.Unmarshal(body, &imgInfo)
	if err != nil {
		return "", fmt.Errorf("imgur - UploadImage - json.Unmarshal - JSON: '%v' err: %w", string(body), err)
	}

	if !imgInfo.Success {
		return "", fmt.Errorf("imgur - UploadImage: %d", imgInfo.Status)
	}

	return imgInfo.Ii.Link, nil
}

// createUploadForm creates a form for uploading an image to imgur.
func createUploadForm(img []byte) url.Values {
	form := url.Values{}
	form.Add("image", string(img[:]))

	return form
}
