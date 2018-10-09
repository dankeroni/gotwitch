package gotwitch

import (
	"net/url"
	"time"

	"github.com/dankeroni/jsonapi"
)

type streamsListChannel struct {
	Data       []Stream `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// Stream json to struct
type Stream struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	GameID       string    `json:"game_id"`
	CommunityIds []string  `json:"community_ids"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
}

type streamsChannel struct {
	Stream Stream `json:"stream"`
}

// GetStream request for GET https://api.twitch.tv/kraken/streams/:channel
func (twitchAPI *TwitchAPI) GetStream(channelName string, onSuccess func(Stream),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var streamsChannel streamsChannel
	onSuccessfulRequest := func() {
		onSuccess(streamsChannel.Stream)
	}
	twitchAPI.Get("/streams/"+channelName, nil, &streamsChannel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// GetStreams request for GET https://api.twitch.tv/kraken/streams?channel=:channelList
// channelList should be a comma-separated list of streams
func (twitchAPI *TwitchAPI) GetStreams(onSuccess func([]Stream),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var streamsListChannel streamsListChannel
	onSuccessfulRequest := func() {
		onSuccess(streamsListChannel.Data)
	}
	parameters := make(url.Values)
	twitchAPI.Get("/streams", parameters, &streamsListChannel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
