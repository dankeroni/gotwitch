package gotwitch

import (
	"net/url"
	"time"

	"github.com/dankeroni/jsonapi"
)

// Stream json to struct
type Stream struct {
	ID          int64     `json:"_id"`
	Game        string    `json:"game"`
	Viewers     int       `json:"viewers"`
	CreatedAt   time.Time `json:"created_at"`
	VideoHeight int       `json:"video_height"`
	AverageFps  float64   `json:"average_fps"`
	Delay       int       `json:"delay"`
	IsPlaylist  bool      `json:"is_playlist"`
	Preview     struct {
		Small    string `json:"small"`
		Medium   string `json:"medium"`
		Large    string `json:"large"`
		Template string `json:"template"`
	} `json:"preview"`
	Channel Channel `json:"channel"`
}

type streamsChannel struct {
	Stream Stream `json:"stream"`
}

type streamsListChannel struct {
	Streams []Stream `json:"streams"`
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
func (twitchAPI *TwitchAPI) GetStreams(channelList string, onSuccess func([]Stream),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var streamsListChannel streamsListChannel
	onSuccessfulRequest := func() {
		onSuccess(streamsListChannel.Streams)
	}
	parameters := make(url.Values)
	parameters.Add("channel", channelList)
	twitchAPI.Get("/streams", parameters, &streamsListChannel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
