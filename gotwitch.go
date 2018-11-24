package gotwitch

import (
	"net/url"

	"github.com/pajlada/jsonapi"
)

type errorResponse struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Credentials struct {
	ClientID       string
	ClientSecret   string
	AppAccessToken string
}

// TwitchAPI struct
type TwitchAPI struct {
	JSONAPI   jsonapi.JSONAPI
	IDJSONAPI jsonapi.JSONAPI

	Credentials *Credentials
}

// New instantiates a new TwitchAPI object
func New(clientID string) *TwitchAPI {
	return &TwitchAPI{
		JSONAPI: jsonapi.JSONAPI{
			BaseURL: "https://api.twitch.tv/helix",
			Headers: map[string]string{
				"Client-ID":    clientID,
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
		},
		IDJSONAPI: jsonapi.JSONAPI{
			BaseURL: "https://id.twitch.tv",
			Headers: map[string]string{
				"Client-ID":    clientID,
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
		},

		Credentials: &Credentials{
			ClientID: clientID,
		},
	}
}

// NewV3 instantiates a new TwitchAPI object
func NewV3(clientID string) *TwitchAPI {
	return &TwitchAPI{
		JSONAPI: jsonapi.JSONAPI{
			BaseURL: "https://api.twitch.tv/kraken",
			Headers: map[string]string{
				"Client-ID": clientID,
				"Accept":    "application/vnd.twitchtv.v3+json",
			},
		},
		IDJSONAPI: jsonapi.JSONAPI{
			BaseURL: "https://id.twitch.tv/",
			Headers: map[string]string{
				"Client-ID":    clientID,
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
		},
	}
}

// Get request
func (twitchAPI *TwitchAPI) Get(url string, parameters url.Values, responseBody interface{},
	onSuccess jsonapi.SuccessCallback, onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	twitchAPI.JSONAPI.Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError, nil)
}

// Post request
func (twitchAPI *TwitchAPI) Post(url string, parameters url.Values, requestBody interface{}, responseBody interface{},
	onSuccess jsonapi.SuccessCallback, onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	twitchAPI.JSONAPI.Post(url, parameters,
		requestBody, responseBody,
		onSuccess, onHTTPError, onInternalError, nil)
}
