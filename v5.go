package gotwitch

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type V5Error struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewTwitchAPIV5(clientID string) *TwitchAPIV5 {
	a := &TwitchAPIV5{
		c: resty.New().
			SetHostURL("https://api.twitch.tv/kraken").
			SetHeader("Accept", "application/vnd.twitchtv.v5+json").
			SetHeader("Client-ID", clientID).
			SetError(&V5Error{}),
	}

	return a
}

type TwitchAPIV5 struct {
	c *resty.Client

	authenticated bool
}

func (a *TwitchAPIV5) Authenticate(oauthToken string) error {
	if oauthToken == "" {
		return errors.New("oauthToken may not be empty")
	}

	a.c.SetHeader("Authorization", "OAuth "+oauthToken)

	a.authenticated = true

	return nil
}

func (a *TwitchAPIV5) Authenticated() bool {
	return a.authenticated
}
