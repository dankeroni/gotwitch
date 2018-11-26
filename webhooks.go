package gotwitch

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/pajlada/jsonapi"
)

type Pagination struct {
	Cursor string `json:"cursor"`
}

type WebhookSubscription struct {
	Topic     string    `json:"topic"`
	Callback  string    `json:"callback"`
	ExpiresAt time.Time `json:"expires_at"`
}

type WebhookSubscriptionsResponse struct {
	Total      int                   `json:"total"`
	Data       []WebhookSubscription `json:"data"`
	Pagination Pagination            `json:"pagination"`
}

type WebhookSubscribeResponse json.RawMessage

// https://dev.twitch.tv/docs/api/reference/#get-webhook-subscriptions
func (twitchAPI *TwitchAPI) GetWebhookSubscriptions(after, first string,
	onSuccess func(WebhookSubscriptionsResponse),
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	var data WebhookSubscriptionsResponse
	onSuccessfulRequest := func() {
		onSuccess(data)
	}

	parameters := make(url.Values)
	if after != "" {
		parameters.Add("after", after)
	}

	firstValue := "20"
	if first != "" {
		firstValue = first
	}
	parameters.Add("first", firstValue)

	return twitchAPI.authenticatedAPI.Get("/webhooks/subscriptions", parameters, &data,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// GetWebhookSubscriptionsSimple simplifies GetWebhookSubscriptions
// https://dev.twitch.tv/docs/api/reference/#get-webhook-subscriptions
func (twitchAPI *TwitchAPI) GetWebhookSubscriptionsSimple(after, first string) (data *WebhookSubscriptionsResponse, response *http.Response, err error) {
	var errorChannel = make(chan error)
	onSuccessfulRequest := func(d WebhookSubscriptionsResponse) {
		data = &d
		errorChannel <- nil
	}
	go func() {
		response, err = twitchAPI.GetWebhookSubscriptions(after, first, onSuccessfulRequest, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))
	}()

	err = <-errorChannel

	return
}

type webhookHubRequest struct {
	Mode     string `json:"hub.mode"`
	Topic    string `json:"hub.topic"`
	Callback string `json:"hub.callback"`
	Lease    int    `json:"hub.lease_seconds"`
	Secret   string `json:"hub.secret"`
}

// https://dev.twitch.tv/docs/api/webhooks-reference/#subscribe-tounsubscribe-from-events
func (twitchAPI *TwitchAPI) WebhookSubscribe(callbackURL string, topic WebhookTopic, twitchUserID string, lease time.Duration, secret string,
	onSuccess func(WebhookSubscribeResponse), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	request := webhookHubRequest{
		Mode:     "subscribe",
		Topic:    topic.URL(twitchUserID),
		Callback: callbackURL,
		Lease:    int(lease.Seconds()),
		Secret:   secret,
	}

	var d WebhookSubscribeResponse
	onSuccessfulRequest := func() {
		onSuccess(d)
	}

	return twitchAPI.authenticatedAPI.Post("/webhooks/hub", nil, &request, &d,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) WebhookSubscribeSimple(callbackURL string, topic WebhookTopic, twitchUserID string, lease time.Duration, secret string) (data *WebhookSubscribeResponse, response *http.Response, err error) {
	var errorChannel = make(chan error)

	onSuccessfulRequest := func(d WebhookSubscribeResponse) {
		data = &d
		errorChannel <- nil
	}

	go func() {
		response, err = twitchAPI.WebhookSubscribe(callbackURL, topic, twitchUserID, lease, secret, onSuccessfulRequest, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))
	}()
	err = <-errorChannel
	return
}
