package gotwitch

// Provides the following functions:
// V5 endpoint: GetClip(clipSlug)

import (
	"net/http"
	"time"

	"github.com/pajlada/jsonapi"
)

type V5GetClipResponse struct {
	Slug        string `json:"slug"`
	TrackingID  string `json:"tracking_id"`
	URL         string `json:"url"`
	EmbedURL    string `json:"embed_url"`
	EmbedHTML   string `json:"embed_html"`
	Broadcaster struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		ChannelURL  string `json:"channel_url"`
		Logo        string `json:"logo"`
	} `json:"broadcaster"`
	Curator struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		ChannelURL  string `json:"channel_url"`
		Logo        string `json:"logo"`
	} `json:"curator"`
	Vod struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"vod"`
	Game       string    `json:"game"`
	Language   string    `json:"language"`
	Title      string    `json:"title"`
	Views      int       `json:"views"`
	Duration   float64   `json:"duration"`
	CreatedAt  time.Time `json:"created_at"`
	Thumbnails struct {
		Medium string `json:"medium"`
		Small  string `json:"small"`
		Tiny   string `json:"tiny"`
	} `json:"thumbnails"`
}

func (twitchAPI *TwitchAPI) GetClipVerbose(clipSlug string,
	onSuccess func(V5GetClipResponse),
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {
	var outerData V5GetClipResponse
	onSuccessfulRequest := func() {
		onSuccess(outerData)
	}
	return twitchAPI.get("/clips/"+clipSlug, nil, &outerData,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) GetClip(clipSlug string) (data V5GetClipResponse, response *http.Response, err error) {
	var errorChannel = make(chan error)
	onSuccessfulRequest := func(d V5GetClipResponse) {
		data = d
		errorChannel <- nil
	}
	go func() {
		response, err = twitchAPI.GetClipVerbose(clipSlug, onSuccessfulRequest, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))
	}()

	err = <-errorChannel

	return
}
