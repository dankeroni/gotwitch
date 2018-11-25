package gotwitch

import (
	"time"

	"github.com/pajlada/jsonapi"
)

// Channel json to struct
type Channel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Language                     string    `json:"language"`
	ID                           int       `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Delay                        int       `json:"delay"`
	Logo                         string    `json:"logo"`
	Banner                       string    `json:"banner"`
	VideoBanner                  string    `json:"video_banner"`
	Background                   string    `json:"background"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	URL                          string    `json:"url"`
	Views                        int       `json:"views"`
	Followers                    int       `json:"followers"`
	Email                        string    `json:"email"`
	StreamKey                    string    `json:"stream_key"`
}

// GetName returns the users display name if it's set, otherwise return the login name
func (c *Channel) GetName() string {
	if c.DisplayName == "" {
		return c.Name
	}

	return c.DisplayName
}

// GetChannel request for GET https://api.twitch.tv/kraken/channels/:channel
func (twitchAPI *TwitchAPI) GetChannel(channelName string, onSuccess func(Channel),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var channel Channel
	onSuccessfulRequest := func() {
		onSuccess(channel)
	}
	twitchAPI.get("/channels/"+channelName, nil, &channel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// AuthenticatedGetChannel request for GET https://api.twitch.tv/kraken/channel
func (twitchAPI *TwitchAPI) AuthenticatedGetChannel(oauthToken string,
	onSuccess func(Channel), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var channel Channel
	onSuccessfulRequest := func() {
		onSuccess(channel)
	}
	twitchAPI.AuthenticatedGet("/channel", nil, oauthToken, &channel,
		onSuccessfulRequest, onHTTPError, onInternalError)
}
