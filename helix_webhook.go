package gotwitch

import "time"

type WebhookSubscription struct {
	Topic     string    `json:"topic"`
	Callback  string    `json:"callback"`
	ExpiresAt time.Time `json:"expires_at"`
}

type webhookHubRequest struct {
	Mode     string `json:"hub.mode"`
	Topic    string `json:"hub.topic"`
	Callback string `json:"hub.callback"`
	Lease    int    `json:"hub.lease_seconds"`
	Secret   string `json:"hub.secret"`
}
