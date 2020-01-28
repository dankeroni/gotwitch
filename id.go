package gotwitch

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type IDError struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type IDResponse struct {
	Data interface{} `json:"data"`

	// TODO: Handle this
	Pagination interface{} `json:"pagination"`
}

func NewTwitchAPIID(clientID string) *TwitchAPIID {
	a := &TwitchAPIID{
		c: resty.New().
			SetHostURL("https://id.twitch.tv/oauth2").
			SetHeader("Accept", "application/json").
			SetHeader("Content-Type", "application/json").
			SetHeader("Client-ID", clientID).
			SetError(&IDError{}),
	}

	return a
}

type TwitchAPIID struct {
	c *resty.Client

	authenticated bool
}

func (a *TwitchAPIID) Authenticate(oauthToken string) error {
	if oauthToken == "" {
		return errors.New("oauthToken may not be empty")
	}

	a.c.SetHeader("Authorization", "OAuth "+oauthToken)

	a.authenticated = true

	return nil
}

func (a *TwitchAPIID) Authenticated() bool {
	return a.authenticated
}
