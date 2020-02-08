package gotwitch

import (
	"encoding/json"
	"fmt"
	"time"
)

type WebhookSubscribeResponse json.RawMessage

func (a *TwitchAPIHelix) WebhookSubscribe(callbackURL string, topic WebhookTopic, twitchUserID string, lease time.Duration, secret string) (result WebhookSubscribeResponse, err error) {
	requestData := webhookHubRequest{
		Mode:     "subscribe",
		Topic:    topic.URL(twitchUserID),
		Callback: callbackURL,
		Lease:    int(lease.Seconds()),
		Secret:   secret,
	}

	request := a.c.R().
		SetResult(&result).
		SetBody(requestData)

	resp, err := request.Post("webhooks/hub")

	if resp.IsError() {
		e := resp.Error().(*HelixError)
		return nil, fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
	}

	return result, err
}
