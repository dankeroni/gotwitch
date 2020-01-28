package gotwitch

import (
	"errors"
	"fmt"
)

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

type UpdateChannelParameters struct {
	Game  *string `json:"game,omitempty"`
	Title *string `json:"status,omitempty"`
}

func NewUpdateChannelParameters() *UpdateChannelParameters {
	return &UpdateChannelParameters{}
}

func (p *UpdateChannelParameters) ResetGame() *UpdateChannelParameters {
	p.Game = nil

	return p
}

func (p *UpdateChannelParameters) SetGame(v string) *UpdateChannelParameters {
	p.Game = &v

	return p
}

func (p *UpdateChannelParameters) ResetTitle() *UpdateChannelParameters {
	p.Title = nil

	return p
}

func (p *UpdateChannelParameters) SetTitle(v string) *UpdateChannelParameters {
	p.Title = &v

	return p
}

func (p *UpdateChannelParameters) Valid() bool {
	if p.Game != nil {
		return true
	}

	if p.Title != nil {
		return true
	}

	return false
}

// UpdateChannel request for GET https://api.twitch.tv/kraken/channel
func (a *TwitchAPIV5) UpdateChannel(channelID string, parameters *UpdateChannelParameters) (*V5Channel, error) {
	if !a.Authenticated() {
		return nil, errors.New("not authenticated")
	}

	channel := &V5Channel{}

	if parameters == nil || !parameters.Valid() {
		return nil, errors.New("Invalid update parameters")
	}

	type tBody struct {
		Channel UpdateChannelParameters `json:"channel"`
	}

	body := tBody{Channel: *parameters}

	resp, err := a.c.R().
		SetBody(body).
		SetResult(&channel).
		Put("channels/" + channelID)

	if resp.IsError() {
		e := resp.Error().(*V5Error)
		return nil, fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
	}

	return channel, err
}
