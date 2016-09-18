package gotwitch

import (
	"github.com/dankeroni/jsonapi"
	"net/url"
	"time"
)

// User json to struct
type User struct {
	DisplayName string    `json:"display_name"`
	ID          int       `json:"_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
}

// FollowedStreams json to struct
type FollowedStreams struct {
	Streams []Stream `json:"streams"`
	Total   int      `json:"_total"`
}

// GetFollowedStreams request for GET https://api.twitch.tv/kraken/streams/followed
func (twitchAPI *TwitchAPI) GetFollowedStreams(oauthToken string, parameters url.Values,
	onSuccess func(FollowedStreams), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var followedStreams FollowedStreams
	onSuccessfulRequest := func() {
		onSuccess(followedStreams)
	}
	twitchAPI.AuthenticatedGet("/streams/followed", parameters, oauthToken, &followedStreams,
		onSuccessfulRequest, onHTTPError, onInternalError)
}
