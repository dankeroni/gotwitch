package gotwitch

import (
	"errors"
	"fmt"
)

var (
	// ErrMissingGameOrTitle is returned from V5's UpdateChannel in case neither a game or title has been passed to the parameters
	ErrMissingGameOrTitle = errors.New("missing required game or title parameters")
)

// V5Channel is the struct returned from the V5 channel get API
type V5Channel struct {
	ID                           string      `json:"_id"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	BroadcasterSoftware          string      `json:"broadcaster_software"`
	BroadcasterType              string      `json:"broadcaster_type"`
	CreatedAt                    string      `json:"created_at"`
	Description                  string      `json:"description"`
	DisplayName                  string      `json:"display_name"`
	Followers                    int         `json:"followers"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	Logo                         string      `json:"logo"`
	Mature                       bool        `json:"mature"`
	Name                         string      `json:"name"`
	Partner                      bool        `json:"partner"`
	PrivacyOptionsEnabled        bool        `json:"privacy_options_enabled"`
	PrivateVideo                 bool        `json:"private_video"`
	ProfileBanner                interface{} `json:"profile_banner"`
	ProfileBannerBackgroundColor interface{} `json:"profile_banner_background_color"`
	Status                       string      `json:"status"`
	UpdatedAt                    string      `json:"updated_at"`
	URL                          string      `json:"url"`
	VideoBanner                  interface{} `json:"video_banner"`
	Views                        int         `json:"views"`
}

// UpdateChannelParameters are parameters sent through to the V5 Update Channel API
type UpdateChannelParameters struct {
	Game  *string `json:"game,omitempty"`
	Title *string `json:"status,omitempty"`
}

// NewUpdateChannelParameters is a helper for creating an UpdateChannelParameters struct
func NewUpdateChannelParameters() *UpdateChannelParameters {
	return &UpdateChannelParameters{}
}

// ResetGame resets the game
func (p *UpdateChannelParameters) ResetGame() *UpdateChannelParameters {
	p.Game = nil

	return p
}

// SetGame sets the game
func (p *UpdateChannelParameters) SetGame(v string) *UpdateChannelParameters {
	p.Game = &v

	return p
}

// ResetTitle resets the title
func (p *UpdateChannelParameters) ResetTitle() *UpdateChannelParameters {
	p.Title = nil

	return p
}

// SetTitle sets the title
func (p *UpdateChannelParameters) SetTitle(v string) *UpdateChannelParameters {
	p.Title = &v

	return p
}

// Validate validates the given parameters
// Rule #1: Either Game or Title MUST be set
func (p *UpdateChannelParameters) Validate() error {
	if p.Game != nil {
		return nil
	}

	if p.Title != nil {
		return nil
	}

	return ErrMissingGameOrTitle
}

// UpdateChannel request for GET https://api.twitch.tv/kraken/channel
func (a *TwitchAPIV5) UpdateChannel(channelID string, parameters *UpdateChannelParameters) (rResult *V5Channel, err error) {
	if !a.Authenticated() {
		err = ErrNotAuthenticated
		return
	}

	result := V5Channel{}

	if parameters == nil {
		err = ErrMissingParameters
		return
	}
	if err = parameters.Validate(); err != nil {
		return
	}

	type tBody struct {
		Channel UpdateChannelParameters `json:"channel"`
	}

	body := tBody{Channel: *parameters}

	resp, err := a.c.R().
		SetBody(body).
		SetResult(&result).
		Put("channels/" + channelID)

	if resp.IsError() {
		e := resp.Error().(*V5Error)
		err = fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
		return
	}

	rResult = &result

	return
}
