package gotwitch

import (
	"fmt"
	"net/url"
)

func simpleOnHTTPError(channel chan error) func(code int, statusMessage, errorMessage string) {
	return func(code int, statusMessage, errorMessage string) {
		channel <- fmt.Errorf("HTTP Error %d: %s - %s", code, statusMessage, errorMessage)
	}
}

func simpleOnInternalError(channel chan error) func(err error) {
	return func(err error) {
		channel <- err
	}
}

func authenticationParameters(oauthToken string, parameters url.Values) url.Values {
	if parameters == nil {
		parameters = url.Values{}
	}
	parameters.Add("oauth_token", oauthToken)
	return parameters
}

func authenticationHeaders(oauthToken string) map[string]string {
	r := make(map[string]string)
	r["Authorization"] = "OAuth " + oauthToken
	return r
}
