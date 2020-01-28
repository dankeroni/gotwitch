package gotwitch

import (
	"errors"
	"fmt"
)

var validCommercialLengths = map[int]struct{}{
	30:  {},
	60:  {},
	90:  {},
	120: {},
	150: {},
	180: {},
}

type StartCommercialParameters struct {
	Length int `json:"length"`
}

func NewStartCommercialParameters() *StartCommercialParameters {
	return &StartCommercialParameters{
		Length: 30,
	}
}

func (p *StartCommercialParameters) ResetLength() *StartCommercialParameters {
	p.Length = 30

	return p
}

func (p *StartCommercialParameters) SetLength(v int) *StartCommercialParameters {
	p.Length = v

	return p
}

func (p *StartCommercialParameters) Valid() bool {
	_, ok := validCommercialLengths[p.Length]
	return ok
}

type V5StartCommercialResponse struct {
	Length     int
	Message    string
	RetryAfter int
}

// StartCommercial request for GET https://dev.twitch.tv/docs/v5/reference/channels#start-channel-commercial
func (a *TwitchAPIV5) StartCommercial(channelID string, parameters *StartCommercialParameters) (*V5StartCommercialResponse, error) {
	if !a.Authenticated() {
		return nil, errors.New("not authenticated")
	}

	channel := &V5StartCommercialResponse{}

	if parameters == nil || !parameters.Valid() {
		return nil, errors.New("Invalid commercial parameters")
	}

	resp, err := a.c.R().
		SetBody(parameters).
		SetResult(&channel).
		Post("channels/" + channelID + "/commercial")

	if resp.IsError() {
		e := resp.Error().(*V5Error)
		return nil, fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
	}

	return channel, err
}
