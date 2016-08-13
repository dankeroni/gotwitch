package gotwitch

import (
	"net/url"
	"time"
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

type streamsFollowed struct {
	Streams []Stream `json:"streams"`
}

// StreamSuccessCallback runs on a successfull request and parse using the Stream method
type StreamSuccessCallback func(Stream)

// FollowedStreamsSuccessCallback runs on a successfull request and parse using the Stream method
type FollowedStreamsSuccessCallback func([]Stream)

// Stream request for https://api.twitch.tv/kraken/streams/:channel
func (twitchAPI *TwitchAPI) Stream(channelName string, onSuccess StreamSuccessCallback,
	onHTTPError HTTPErrorCallback, onInternalError InternalErrorCallback) {
	var streamsChannel streamsChannel
	onSuccessfulRequest := func() {
		onSuccess(streamsChannel.Stream)
	}
	twitchAPI.Get("/streams/"+channelName, nil, "", &streamsChannel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// FollowedStreams request for https://api.twitch.tv/kraken/streams/followed
func (twitchAPI *TwitchAPI) FollowedStreams(oauthToken string, parameters url.Values,
	onSuccess FollowedStreamsSuccessCallback, onHTTPError HTTPErrorCallback,
	onInternalError InternalErrorCallback) {
	var streamsFollowed streamsFollowed
	onSuccessfulRequest := func() {
		onSuccess(streamsFollowed.Streams)
	}
	twitchAPI.Get("/streams/followed", parameters, oauthToken, &streamsFollowed,
		onSuccessfulRequest, onHTTPError, onInternalError)
}
