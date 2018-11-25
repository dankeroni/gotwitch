package gotwitch

import (
	"net/http"
	"net/url"

	"github.com/pajlada/jsonapi"
)

// IDAuthenticatedGet request on the ID endpoint
func (twitchAPI *TwitchAPI) IDAuthenticatedGet(url string,
	parameters url.Values,
	oauthToken string,
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	return twitchAPI.idAPI.R().SetHeader("Authorization", "OAuth "+oauthToken).Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError)
}

// AuthenticatedGet request
func (twitchAPI *TwitchAPI) AuthenticatedGet(url string,
	parameters url.Values,
	oauthToken string,
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	return twitchAPI.authUser(oauthToken).Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError)
}

// AuthenticatedPut request
func (twitchAPI *TwitchAPI) AuthenticatedPut(url string,
	parameters url.Values,
	oauthToken string,
	requestBody interface{},
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	return twitchAPI.authUser(oauthToken).Put(url, parameters, requestBody, responseBody, onSuccess, onHTTPError, onInternalError)
}

// AuthenticatedPost request
func (twitchAPI *TwitchAPI) AuthenticatedPost(url string,
	parameters url.Values,
	oauthToken string,
	requestBody interface{},
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	return twitchAPI.authUser(oauthToken).Post(url, parameters, requestBody, responseBody, onSuccess, onHTTPError, onInternalError)
}

// AuthenticatedDelete request
func (twitchAPI *TwitchAPI) AuthenticatedDelete(url string,
	parameters url.Values,
	oauthToken string,
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {

	return twitchAPI.authUser(oauthToken).Delete(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError)
}
