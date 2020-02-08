package gotwitch

import (
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
		clientID: clientID,

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
	clientID string

	c *resty.Client

	authenticated bool

	clientSecret string
}

func (a *TwitchAPIID) SetClientSecret(clientSecret string) {
	a.clientSecret = clientSecret
}

func (a *TwitchAPIID) Authenticate(oauthToken string) *TwitchAPIID {
	id := NewTwitchAPIID(a.clientID)
	id.SetClientSecret(a.clientSecret)
	id.c.SetHeader("Authorization", "OAuth "+oauthToken)
	id.authenticated = true
	return id
}

func (a *TwitchAPIID) Authenticated() bool {
	return a.authenticated
}
