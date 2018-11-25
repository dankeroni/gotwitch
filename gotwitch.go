package gotwitch

import (
	"net/http"
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

func (c *Credentials) middleware(r *http.Request) error {
	if c.AppAccessToken != "" {
		r.Header.Set("Authorization", "Bearer "+c.AppAccessToken)
	}
	return nil
}

// TwitchAPI struct
type TwitchAPI struct {
	unauthenticatedAPI jsonapi.JSONAPI
	authenticatedAPI   jsonapi.JSONAPI
	idAPI              jsonapi.JSONAPI

	Credentials *Credentials
}

// New instantiates a new TwitchAPI object
func New(clientID string) *TwitchAPI {
	api := &TwitchAPI{
		unauthenticatedAPI: jsonapi.JSONAPI{
			BaseURL: "https://api.twitch.tv/helix",
			Headers: map[string]string{
				"Client-ID":    clientID,
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
		},
		// authenticated api is authenticated with the server app token
		// for requests that need to be authenticated by a user, user unauthenticatedAPI and set the header manually
		authenticatedAPI: jsonapi.JSONAPI{
			BaseURL: "https://api.twitch.tv/helix",
			Headers: map[string]string{
				"Client-ID":    clientID,
				"Accept":       "application/json",
				"Content-Type": "application/json",
			},
		},
		idAPI: jsonapi.JSONAPI{
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

	api.authenticatedAPI.Use(api.Credentials.middleware)

	return api
}

// Get request
func (twitchAPI *TwitchAPI) get(url string,
	parameters url.Values,
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (r *http.Response, err error) {
	return twitchAPI.authenticatedAPI.Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError)
}

// Post request
func (twitchAPI *TwitchAPI) post(url string,
	parameters url.Values,
	requestBody interface{},
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (r *http.Response, err error) {
	return twitchAPI.authenticatedAPI.Post(url, parameters, requestBody, responseBody, onSuccess, onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) authUser(oauthToken string) *jsonapi.Request {
	return twitchAPI.unauthenticatedAPI.R().SetHeader("Authorization", "Bearer "+oauthToken)
}
