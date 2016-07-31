package gotwitch

import "time"

// Channel json to struct
type Channel struct {
	Mature                       interface{} `json:"mature"`
	Status                       string      `json:"status"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	DisplayName                  string      `json:"display_name"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	ID                           int         `json:"_id"`
	Name                         string      `json:"name"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Delay                        interface{} `json:"delay"`
	Logo                         string      `json:"logo"`
	Banner                       interface{} `json:"banner"`
	VideoBanner                  string      `json:"video_banner"`
	Background                   interface{} `json:"background"`
	ProfileBanner                string      `json:"profile_banner"`
	ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
	Partner                      bool        `json:"partner"`
	URL                          string      `json:"url"`
	Views                        int         `json:"views"`
	Followers                    int         `json:"followers"`
	Links                        struct {
		Self          string `json:"self"`
		Follows       string `json:"follows"`
		Commercial    string `json:"commercial"`
		StreamKey     string `json:"stream_key"`
		Chat          string `json:"chat"`
		Features      string `json:"features"`
		Subscriptions string `json:"subscriptions"`
		Editors       string `json:"editors"`
		Teams         string `json:"teams"`
		Videos        string `json:"videos"`
	} `json:"_links"`
}
