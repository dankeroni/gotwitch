package gotwitch

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

// HelixStream json to struct
type HelixStream struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	GameID       string    `json:"game_id"`
	CommunityIds []string  `json:"community_ids"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
}

type GetStreamsParameters struct {
	GameID     string
	Language   string
	UserIDs    []string
	UserLogins []string
}

func NewGetStreamsParameters() *GetStreamsParameters {
	return &GetStreamsParameters{}
}

func (p *GetStreamsParameters) SetUserIDs(v []string) *GetStreamsParameters {
	p.UserIDs = v

	return p
}

func (p *GetStreamsParameters) Validate() error {
	for _, v := range p.UserIDs {
		if v == "" {
			return ErrCannotPassEmptyStringAsLookupValue
		}
	}

	return nil
}

func (p *GetStreamsParameters) Apply(req *resty.Request) {
	values := url.Values{}

	for _, userID := range p.UserIDs {
		values.Add("user_id", userID)
	}

	req.SetQueryParamsFromValues(values)
}

func (a *TwitchAPIHelix) GetStreams(parameters *GetStreamsParameters) ([]HelixStream, error) {
	type Result struct {
		Data []HelixStream `json:"data"`

		// TODO: Handle this
		Pagination interface{} `json:"pagination"`
	}

	if err := parameters.Validate(); err != nil {
		return nil, err
	}

	result := Result{}

	request := a.c.R().
		SetResult(&result)

	if parameters != nil {
		parameters.Apply(request)
	}

	resp, err := request.
		Get("streams")

	if resp.IsError() {
		e := resp.Error().(*HelixError)
		return nil, fmt.Errorf("API error code %d: %s - %s", e.Status, e.Error, e.Message)
	}

	return result.Data, err
}
