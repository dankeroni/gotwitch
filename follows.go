package gotwitch

import (
	"github.com/dankeroni/jsonapi"
	"net/url"
	"time"
)

// Follow json to struct
type Follow struct {
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
	Channel       Channel   `json:"channel"`
}

// Follows json to struct
type Follows struct {
	Follows []Follow `json:"follows"`
	Total   int      `json:"_total"`
}

// GetFollows request for GET https://api.twitch.tv/kraken/users/:user/follows/channels
func (twitchAPI *TwitchAPI) GetFollows(user string, parameters url.Values, onSuccess func(Follows),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var follows Follows
	onSuccessfulRequest := func() {
		onSuccess(follows)
	}
	twitchAPI.Get("/users/"+user+"/follows/channels", parameters, &follows,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// GetFollow request for GET https://api.twitch.tv/kraken/users/:user/follows/channels/:target
func (twitchAPI *TwitchAPI) GetFollow(user, target string, onSuccess func(Follow),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var follow Follow
	onSuccessfulRequest := func() {
		onSuccess(follow)
	}
	twitchAPI.Get("/users/"+user+"/follows/channels/"+target, nil, &follow, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// PutFollow request for PUT https://api.twitch.tv/kraken/users/:user/follows/channels/:target
func (twitchAPI *TwitchAPI) PutFollow(oauthToken, user, target string, onSuccess func(Follow),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var follow Follow
	onSuccessfulRequest := func() {
		onSuccess(follow)
	}
	twitchAPI.AuthenticatedPut("/users/"+user+"/follows/channels/"+target, nil, oauthToken, nil, &follow,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// DeleteFollow request for DELETE https://api.twitch.tv/kraken/users/:user/follows/channels/:target
func (twitchAPI *TwitchAPI) DeleteFollow(oauthToken, user, target string, onSuccess func(),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	twitchAPI.AuthenticatedDelete("/users/"+user+"/follows/channels/"+target, nil, oauthToken, nil, onSuccess,
		onHTTPError, onInternalError)
}
