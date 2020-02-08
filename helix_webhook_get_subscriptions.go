package gotwitch

import "fmt"

type WebhookSubscriptionsResponse struct {
	Total int                   `json:"total"`
	Data  []WebhookSubscription `json:"data"`

	// TODO: Handle this
	Pagination interface{} `json:"pagination"`
}

func (a *TwitchAPIHelix) GetWebhookSubscriptions(after, first string) (result WebhookSubscriptionsResponse, err error) {
	request := a.c.R().
		SetResult(&result)

	if after != "" {
		request.SetQueryParam("after", after)
	}

	if first != "" {
		request.SetQueryParam("first", first)
	} else {
		request.SetQueryParam("first", "20")
	}

	resp, err := request.Get("webhooks/subscriptions")

	if resp.IsError() {
		e := resp.Error().(*HelixError)
		err = fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
		return
	}

	return result, err
}
