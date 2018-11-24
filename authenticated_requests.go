package gotwitch

import (
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
	onInternalError jsonapi.InternalErrorCallback) {

	extraHeaders := authenticationHeaders(oauthToken)
	twitchAPI.IDJSONAPI.Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError, extraHeaders)
}

// AuthenticatedGet request
func (twitchAPI *TwitchAPI) AuthenticatedGet(url string,
	parameters url.Values,
	oauthToken string,
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {

	extraHeaders := authenticationHeaders(oauthToken)
	twitchAPI.JSONAPI.Get(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError, extraHeaders)
}

// AuthenticatedPut request
func (twitchAPI *TwitchAPI) AuthenticatedPut(url string,
	parameters url.Values,
	oauthToken string,
	requestBody interface{},
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {

	extraHeaders := authenticationHeaders(oauthToken)
	twitchAPI.JSONAPI.Put(url, parameters, requestBody, responseBody,
		onSuccess, onHTTPError, onInternalError, extraHeaders)
}

// AuthenticatedPost request
func (twitchAPI *TwitchAPI) AuthenticatedPost(url string,
	parameters url.Values,
	oauthToken string,
	requestBody interface{},
	responseBody interface{},
	onSuccess jsonapi.SuccessCallback,
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {

	extraHeaders := authenticationHeaders(oauthToken)
	twitchAPI.JSONAPI.Post(url, parameters, requestBody, responseBody,
		onSuccess, onHTTPError, onInternalError, extraHeaders)
}

// AuthenticatedDelete request
func (twitchAPI *TwitchAPI) AuthenticatedDelete(url string, parameters url.Values, oauthToken string,
	responseBody interface{}, onSuccess jsonapi.SuccessCallback, onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {

	extraHeaders := authenticationHeaders(oauthToken)
	twitchAPI.JSONAPI.Delete(url, parameters, responseBody, onSuccess, onHTTPError, onInternalError, extraHeaders)
}
