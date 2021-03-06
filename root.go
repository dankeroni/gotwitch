package gotwitch

import (
	"net/http"
	"time"

	"github.com/pajlada/jsonapi"
)

// Self json to struct
type Self struct {
	Identified bool `json:"identified"`
	Links      struct {
		User     string `json:"user"`
		Channel  string `json:"channel"`
		Search   string `json:"search"`
		Streams  string `json:"streams"`
		Ingests  string `json:"ingests"`
		Teams    string `json:"teams"`
		Users    string `json:"users"`
		Channels string `json:"channels"`
		Chat     string `json:"chat"`
	} `json:"_links"`
	Token struct {
		Valid         bool `json:"valid"`
		Authorization struct {
			Scopes    []string  `json:"scopes"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"authorization"`
		UserName string `json:"user_name"`
		ClientID string `json:"client_id"`
	} `json:"token"`
}

// GetSelf request for GET https://api.twitch.tv/kraken/
func (twitchAPI *TwitchAPI) GetSelf(oauthToken string, onSuccess func(Self),
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {
	var self Self
	onSuccessfulRequest := func() {
		onSuccess(self)
	}
	return twitchAPI.AuthenticatedGet("/", nil, oauthToken, &self,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) GetSelfSimple(oauthToken string) (data *Self, response *http.Response, err error) {
	var errorChannel = make(chan error)
	onSuccessfulRequest := func(d Self) {
		data = &d
		errorChannel <- nil
	}
	go func() {
		response, err = twitchAPI.GetSelf(oauthToken, onSuccessfulRequest, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))
	}()

	err = <-errorChannel

	return
}
