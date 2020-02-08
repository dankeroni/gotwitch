package gotwitch

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type TMIError json.RawMessage

type TwitchAPITMI struct {
	c *resty.Client
}

func NewTwitchAPITMI(clientID string) *TwitchAPITMI {
	a := &TwitchAPITMI{
		c: resty.New().
			SetHostURL("https://tmi.twitch.tv/helix").
			SetHeader("Accept", "application/json").
			SetHeader("Content-Type", "application/json").
			SetHeader("Client-ID", clientID).
			SetError(&TMIError{}),
	}

	return a
}
