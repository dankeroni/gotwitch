package gotwitch

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

var (
	ErrGetUsersNoUsers      = errors.New("at least one user must be specified")
	ErrGetUsersTooManyUsers = errors.New("too many users specified")
)

// User json to struct (HELIX)
type User struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

type GetUsersParameters struct {
	UserIDs    []string
	UserLogins []string
}

func NewGetUsersParameters() *GetUsersParameters {
	return &GetUsersParameters{}
}

func (p *GetUsersParameters) SetUserIDs(v []string) *GetUsersParameters {
	p.UserIDs = v

	return p
}

func (p *GetUsersParameters) SetUserLogins(v []string) *GetUsersParameters {
	p.UserLogins = v

	return p
}

func (p *GetUsersParameters) Validate() error {
	total := len(p.UserIDs) + len(p.UserLogins)

	if total == 0 {
		return ErrGetUsersNoUsers
	}

	if total > 100 {
		return ErrGetUsersTooManyUsers
	}

	for _, v := range p.UserIDs {
		if v == "" {
			return ErrCannotPassEmptyStringAsLookupValue
		}
	}

	for _, v := range p.UserLogins {
		if v == "" {
			return ErrCannotPassEmptyStringAsLookupValue
		}
	}

	return nil
}

func (p *GetUsersParameters) Apply(req *resty.Request) {
	values := url.Values{}

	for _, userID := range p.UserIDs {
		values.Add("id", userID)
	}

	for _, userID := range p.UserLogins {
		values.Add("login", userID)
	}

	req.SetQueryParamsFromValues(values)
}

// GetUsers request for GET https://api.twitch.tv/helix/users
func (a *TwitchAPIHelix) GetUsers(parameters *GetUsersParameters) ([]User, error) {
	type Result struct {
		Data []User `json:"data"`
	}

	result := Result{}

	if parameters == nil {
		return nil, ErrMissingParameters
	}

	if err := parameters.Validate(); err != nil {
		return nil, err
	}

	request := a.c.R().
		SetResult(&result)

	if parameters != nil {
		parameters.Apply(request)
	}

	resp, err := request.
		Get("users")

	if resp.IsError() {
		e := resp.Error().(*HelixError)
		return nil, fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
	}

	return result.Data, err
}
