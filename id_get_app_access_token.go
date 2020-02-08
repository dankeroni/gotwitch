package gotwitch

import "fmt"

func (a *TwitchAPIID) GetAppAccessToken() (Token, error) {
	return a.Token("client_credentials")
}

func (a *TwitchAPIID) Token(grantType string) (result Token, err error) {
	if a.clientSecret == "" {
		err = ErrMissingClientSecret
		return
	}

	request := a.c.R().
		SetResult(&result)

	request.SetQueryParam("client_id", a.clientID)
	request.SetQueryParam("client_secret", a.clientSecret)
	request.SetQueryParam("grant_type", grantType)

	resp, err := request.
		Post("token")

	if resp.IsError() {
		e := resp.Error().(*IDError)
		err = fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
		return
	}

	return
}
