package gotwitch

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type HelixError struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type HelixResponse struct {
	Data interface{} `json:"data"`

	// TODO: Handle this
	Pagination interface{} `json:"pagination"`
}

func NewTwitchAPIHelix(clientID string) *TwitchAPIHelix {
	a := &TwitchAPIHelix{
		c: resty.New().
			SetHostURL("https://api.twitch.tv/helix").
			SetHeader("Accept", "application/json").
			SetHeader("Content-Type", "application/json").
			SetHeader("Client-ID", clientID).
			SetError(&HelixError{}),
	}

	return a
}

type TwitchAPIHelix struct {
	c *resty.Client

	authenticated bool
}

func (a *TwitchAPIHelix) Authenticate(oauthToken string) error {
	if oauthToken == "" {
		return errors.New("oauthToken may not be empty")
	}

	a.c.SetHeader("Authorization", "OAuth "+oauthToken)

	a.authenticated = true

	return nil
}

func (a *TwitchAPIHelix) Authenticated() bool {
	return a.authenticated
}
