package gotwitch

import (
	"errors"
	"fmt"
)

type ValidateResponse struct {
	ClientID string   `json:"client_id"`
	Login    string   `json:"login"`
	UserID   string   `json:"user_id"`
	Scopes   []string `json:"scopes"`
}

// Validate request for GET https://id.twitch.tv/oauth2/validate
func (a *TwitchAPIID) Validate() (ValidateResponse, error) {
	result := ValidateResponse{}

	if !a.Authenticated() {
		return result, errors.New("not authenticated")
	}

	request := a.c.R().
		SetResult(&result)

	resp, err := request.
		Get("validate")

	if resp.IsError() {
		e := resp.Error().(*IDError)
		return result, fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
	}

	return result, err
}
