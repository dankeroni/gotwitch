package gotwitch

import (
	"net/http"
	"net/url"

	"github.com/pajlada/jsonapi"
)

type AppAccessTokenResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scopes       []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

// GetAppAccessToken will also authenticate this TwitchAPI object, so any further requests that require the "app access token" has them
func (twitchAPI *TwitchAPI) GetAppAccessToken(onSuccess func(AppAccessTokenResponse), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {

	var response AppAccessTokenResponse
	onSuccessfulRequest := func() {
		twitchAPI.Credentials.AppAccessToken = response.AccessToken
		onSuccess(response)
	}

	parameters := make(url.Values)
	parameters.Add("client_id", twitchAPI.Credentials.ClientID)
	parameters.Add("client_secret", twitchAPI.Credentials.ClientSecret)
	parameters.Add("grant_type", "client_credentials")

	twitchAPI.idAPI.Post("/oauth2/token", parameters, nil, &response, onSuccessfulRequest, onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) GetAppAccessTokenSimple() (response *AppAccessTokenResponse, err error) {
	var errorChannel = make(chan error)
	onSuccessfulRequest := func(r AppAccessTokenResponse) {
		response = &r
		errorChannel <- nil
	}

	go twitchAPI.GetAppAccessToken(onSuccessfulRequest, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))

	err = <-errorChannel

	return
}

type ValidateResponse struct {
	ClientID string   `json:"client_id"`
	Login    string   `json:"login"`
	Scopes   []string `json:"scopes"`
	UserID   string   `json:"user_id"`
}

// ValidateOAuthToken will validate an oauth token with your client id
func (twitchAPI *TwitchAPI) ValidateOAuthToken(oauthToken string, onSuccess func(ValidateResponse), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	var data ValidateResponse
	onSuccessfulRequest := func() {
		onSuccess(data)
	}

	return twitchAPI.IDAuthenticatedGet("/oauth2/validate", nil, oauthToken, &data, onSuccessfulRequest, onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) ValidateOAuthTokenSimple(oauthToken string) (data *ValidateResponse, response *http.Response, err error) {
	var errorChannel = make(chan error)

	onSuccessfulRequest := func(r ValidateResponse) {
		data = &r
		errorChannel <- nil
	}

	go func() {
		response, err = twitchAPI.ValidateOAuthToken(oauthToken, onSuccessfulRequest, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))
	}()

	err = <-errorChannel

	return
}
