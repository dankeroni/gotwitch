package gotwitch

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidCommercialLength is returned from StartCommercial if the commercial length parameter is INVALID
	ErrInvalidCommercialLength = errors.New("invalid commercial length")

	validCommercialLengths = map[int]struct{}{
		30:  {},
		60:  {},
		90:  {},
		120: {},
		150: {},
		180: {},
	}
)

// StartCommercialParameters are parameters sent through to the V5 StartCommercial function
type StartCommercialParameters struct {
	Length int `json:"length"`
}

// NewStartCommercialParameters creates a start commercial parameters struct with proper default values
func NewStartCommercialParameters() *StartCommercialParameters {
	return &StartCommercialParameters{
		Length: 30,
	}
}

// ResetLength resets the length of the parameters to its default value (30)
func (p *StartCommercialParameters) ResetLength() *StartCommercialParameters {
	p.Length = 30

	return p
}

// SetLength sets the commercial length to the given value. Note that parameter validity is now checked here, but in the StartCommercial function
func (p *StartCommercialParameters) SetLength(v int) *StartCommercialParameters {
	p.Length = v

	return p
}

// Validate validates the parameters
func (p *StartCommercialParameters) Validate() error {
	if _, ok := validCommercialLengths[p.Length]; !ok {
		return ErrInvalidCommercialLength
	}

	return nil
}

// V5StartCommercialResponse is returned from V5's StartCommercial function
// The fields indicate how long the commercial will run and in how long another commercial can be run
type V5StartCommercialResponse struct {
	Length     int
	Message    string
	RetryAfter int
}

// StartCommercial request for GET https://dev.twitch.tv/docs/v5/reference/channels#start-channel-commercial
func (a *TwitchAPIV5) StartCommercial(channelID string, parameters *StartCommercialParameters) (rResult *V5StartCommercialResponse, err error) {
	if !a.Authenticated() {
		err = ErrNotAuthenticated
		return
	}

	result := V5StartCommercialResponse{}

	if parameters == nil {
		err = ErrMissingParameters
		return
	}
	if err = parameters.Validate(); err != nil {
		return
	}

	resp, err := a.c.R().
		SetBody(parameters).
		SetResult(&result).
		Post("channels/" + channelID + "/commercial")

	if resp.IsError() {
		e := resp.Error().(*V5Error)
		return nil, fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
	}

	rResult = &result

	return
}
