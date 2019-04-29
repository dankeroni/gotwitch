package gotwitch

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pajlada/jsonapi"
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
	twitchAPI.get("/users/"+user+"/follows/channels", parameters, &follows,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// GetFollow request for GET https://api.twitch.tv/kraken/users/:user/follows/channels/:target
func (twitchAPI *TwitchAPI) GetFollow(user, target string, onSuccess func(Follow),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var follow Follow
	onSuccessfulRequest := func() {
		onSuccess(follow)
	}
	twitchAPI.get("/users/"+user+"/follows/channels/"+target, nil, &follow, onSuccessfulRequest,
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

type Webhook struct {
	// URL where notifications will be delivered
	CallbackURL string `json:"hub.callback"`

	// Type of request. Valid values: subscribe, unsubscribe
	Mode string `json:"hub.mode"`

	// URL for the topic to subscribe to or unsubscribe from
	Topic string `json:"hub.topic"`

	LeaseSeconds int `json:"hub.lease_seconds,omitempty"`

	Secret string `json:"hub.secret,omitempty"`
}

// SubscribeFollows subscribes to the follow webhook
func (twitchAPI *TwitchAPI) SubscribeFollows(userID, callbackURL string, onSuccess func(), onError func()) {
	var follows Follows
	onSuccessfulRequest := func() {
		onSuccess()
	}
	requestBody := Webhook{
		CallbackURL:  callbackURL,
		Mode:         "subscribe",
		Topic:        "https://api.twitch.tv/helix/users/follows?first=1&to_id=" + userID,
		LeaseSeconds: 600,
	}

	onHTTPError := func(statusCode int, statusMessage, errorMessage string) {
		fmt.Println("error", errorMessage, "ok", statusMessage)
		onError()
	}

	onInternalError := func(err error) {
		fmt.Println(err)
		onError()
	}

	parameters := url.Values{}
	twitchAPI.post("/webhooks/hub", parameters, requestBody, &follows,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// SubscribeStreams xd
func (twitchAPI *TwitchAPI) SubscribeStreams(userID, callbackURL string, onSuccess func(), onError func()) {
	var follows Follows
	onSuccessfulRequest := func() {
		onSuccess()
	}
	requestBody := Webhook{
		CallbackURL:  callbackURL,
		Mode:         "subscribe",
		Topic:        "https://api.twitch.tv/helix/streams?user_id=" + userID,
		LeaseSeconds: 3600,
	}

	onHTTPError := func(statusCode int, statusMessage, errorMessage string) {
		fmt.Println("subscriebstreamseirror", errorMessage, "ok", statusMessage)
		onError()
	}

	onInternalError := func(err error) {
		fmt.Println(err)
		onError()
	}

	parameters := url.Values{}
	twitchAPI.post("/webhooks/hub", parameters, requestBody, &follows,
		onSuccessfulRequest, onHTTPError, onInternalError)
}
