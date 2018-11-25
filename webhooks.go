package gotwitch

import (
	"net/url"
	"time"

	"github.com/pajlada/jsonapi"
)

type WebhookSubscriptionsResponse struct {
	Total int `json:"total"`
	Data  []struct {
		Topic     string    `json:"topic"`
		Callback  string    `json:"callback"`
		ExpiresAt time.Time `json:"expires_at"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// https://dev.twitch.tv/docs/api/reference/#get-webhook-subscriptions
func (twitchAPI *TwitchAPI) GetWebhookSubscriptions(after, first string,
	onSuccess func(WebhookSubscriptionsResponse), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {

	var response WebhookSubscriptionsResponse
	onSuccessfulRequest := func() {
		onSuccess(response)
	}

	parameters := make(url.Values)
	if after != "" {
		parameters.Add("after", after)
	}

	if first != "" {
		parameters.Add("first", first)
	}

	twitchAPI.authenticatedAPI.Get("/webhooks/subscriptions", parameters, &response,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

type WebhookTopic int

const (
	WebhookTopicFollowers = iota
	WebhookTopicStreams
	WebhookTopicUserChanged
)

func (t WebhookTopic) URL(twitchUserID string) string {
	switch t {
	case WebhookTopicFollowers:
		return "https://api.twitch.tv/helix/users/follows?first=1&to_id=" + twitchUserID
	case WebhookTopicStreams:
		return "https://api.twitch.tv/helix/streams?user_id=" + twitchUserID
	case WebhookTopicUserChanged:
		return "https://api.twitch.tv/helix/users?id=" + twitchUserID
	}

	return ""
}

type webhookHubRequest struct {
	Mode     string `json:"hub.mode"`
	Topic    string `json:"hub.topic"`
	Callback string `json:"hub.callback"`
	Lease    int    `json:"hub.lease_seconds"`
	Secret   string `json:"hub.secret"`
}

// https://dev.twitch.tv/docs/api/webhooks-reference/#subscribe-tounsubscribe-from-events
func (twitchAPI *TwitchAPI) WebhookSubscribe(callbackURL string, topic WebhookTopic, twitchUserID string, lease int, secret string,
	onSuccess func(WebhookSubscriptionsResponse), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {

	request := webhookHubRequest{
		Mode:     "subscribe",
		Topic:    topic.URL(twitchUserID),
		Callback: callbackURL,
		Lease:    lease,
		Secret:   secret,
	}

	var response WebhookSubscriptionsResponse
	onSuccessfulRequest := func() {
		onSuccess(response)
	}

	twitchAPI.post("/webhooks/hub", nil, &request, &response,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) WebhookSubscribeSimple(callbackURL string, topic WebhookTopic, twitchUserID string, lease int, secret string) (response *WebhookSubscriptionsResponse, err error) {
	var errorChannel = make(chan error)

	onSuccessfulRequest := func(r WebhookSubscriptionsResponse) {
		response = &r
		errorChannel <- nil
	}

	go twitchAPI.WebhookSubscribe(callbackURL, topic, twitchUserID, lease, secret, onSuccessfulRequest, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))
	err = <-errorChannel
	return
}
