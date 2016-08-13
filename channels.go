package gotwitch

import "time"

// Channel json to struct
type Channel struct {
	Mature                       interface{} `json:"mature"`
	Status                       string      `json:"status"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	DisplayName                  string      `json:"display_name"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	ID                           int         `json:"_id"`
	Name                         string      `json:"name"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Delay                        interface{} `json:"delay"`
	Logo                         string      `json:"logo"`
	Banner                       interface{} `json:"banner"`
	VideoBanner                  string      `json:"video_banner"`
	Background                   interface{} `json:"background"`
	ProfileBanner                string      `json:"profile_banner"`
	ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
	Partner                      bool        `json:"partner"`
	URL                          string      `json:"url"`
	Views                        int         `json:"views"`
	Followers                    int         `json:"followers"`
}

// AuthenticatedChannel json to struct
type AuthenticatedChannel struct {
	Mature                       bool        `json:"mature"`
	Status                       string      `json:"status"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	DisplayName                  string      `json:"display_name"`
	Game                         string      `json:"game"`
	Delay                        interface{} `json:"delay"`
	Language                     string      `json:"language"`
	ID                           int         `json:"_id"`
	Name                         string      `json:"name"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Logo                         string      `json:"logo"`
	Banner                       string      `json:"banner"`
	VideoBanner                  string      `json:"video_banner"`
	Background                   interface{} `json:"background"`
	ProfileBanner                string      `json:"profile_banner"`
	ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
	Partner                      bool        `json:"partner"`
	URL                          string      `json:"url"`
	Views                        int         `json:"views"`
	Followers                    int         `json:"followers"`
	Email                        string      `json:"email"`
	StreamKey                    string      `json:"stream_key"`
}

// ChannelSuccessCallback runs on a successfull request and parse using the Channel method
type ChannelSuccessCallback func(Channel)

// AuthenticatedChannelSuccessCallback runs on a successfull request and parse
// using the AuthenticatedChannel method
type AuthenticatedChannelSuccessCallback func(AuthenticatedChannel)

// Channel request for https://api.twitch.tv/kraken/channels/:channel
func (twitchAPI *TwitchAPI) Channel(channelName string, onSuccess ChannelSuccessCallback,
	onHTTPError HTTPErrorCallback, onInternalError InternalErrorCallback) {
	var channel Channel
	onSuccessfulRequest := func() {
		onSuccess(channel)
	}
	twitchAPI.Get("/channels/"+channelName, nil, "", &channel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// AuthenticatedChannel request for https://api.twitch.tv/kraken/channel
func (twitchAPI *TwitchAPI) AuthenticatedChannel(oauthToken string,
	onSuccess AuthenticatedChannelSuccessCallback, onHTTPError HTTPErrorCallback,
	onInternalError InternalErrorCallback) {
	var authenticatedChannel AuthenticatedChannel
	onSuccessfulRequest := func() {
		onSuccess(authenticatedChannel)
	}
	twitchAPI.Get("/channel", nil, oauthToken, &authenticatedChannel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
