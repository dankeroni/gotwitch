package gotwitch

import (
	"net/url"

	"github.com/dankeroni/jsonapi"
)

type errorResponse struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// TwitchAPI struct
type TwitchAPI struct {
	JSONAPI      jsonapi.JSONAPI
	ClientID     string
	ClientSecret string
}

// New instantiates a new TwitchAPI object
func New(clientID string) *TwitchAPI {
	return &TwitchAPI{
		JSONAPI: jsonapi.JSONAPI{
			BaseURL: "https://api.twitch.tv/helix",
			Headers: map[string]string{
				"Client-ID": clientID,
				"Accept":    "application/json",
			},
		},
	}
}

// AuthenticatedGet request
func (twitchAPI *TwitchAPI) AuthenticatedGet(url string, parameters url.Values, oauthToken string,
	responseBody interface{}, onSuccess jsonapi.SuccessCallback, onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	parameters = authenticationParameters(oauthToken, parameters)
	twitchAPI.Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError)
}

// AuthenticatedPut request
func (twitchAPI *TwitchAPI) AuthenticatedPut(url string, parameters url.Values, oauthToken string,
	requestBody interface{}, responseBody interface{}, onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	parameters = authenticationParameters(oauthToken, parameters)
	twitchAPI.JSONAPI.Put(url, parameters, requestBody, responseBody,
		onSuccess, onHTTPError, onInternalError)
}

// AuthenticatedPost request
func (twitchAPI *TwitchAPI) AuthenticatedPost(url string, parameters url.Values, oauthToken string,
	requestBody interface{}, responseBody interface{}, onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	parameters = authenticationParameters(oauthToken, parameters)
	twitchAPI.JSONAPI.Post(url, parameters, requestBody, responseBody,
		onSuccess, onHTTPError, onInternalError)
}

// AuthenticatedDelete request
func (twitchAPI *TwitchAPI) AuthenticatedDelete(url string, parameters url.Values, oauthToken string,
	responseBody interface{}, onSuccess jsonapi.SuccessCallback, onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	parameters = authenticationParameters(oauthToken, parameters)
	twitchAPI.JSONAPI.Delete(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError)
}

// Get request
func (twitchAPI *TwitchAPI) Get(url string, parameters url.Values, responseBody interface{},
	onSuccess jsonapi.SuccessCallback, onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	twitchAPI.JSONAPI.Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError)
}

func authenticationParameters(oauthToken string, parameters url.Values) url.Values {
	if parameters == nil {
		parameters = url.Values{}
	}
	parameters.Add("oauth_token", oauthToken)
	return parameters
}
