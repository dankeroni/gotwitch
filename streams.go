package gotwitch

import "time"

// Stream json to struct
type Stream struct {
	ID          int64     `json:"_id"`
	Game        string    `json:"game"`
	Viewers     int       `json:"viewers"`
	CreatedAt   time.Time `json:"created_at"`
	VideoHeight int       `json:"video_height"`
	AverageFps  int       `json:"average_fps"`
	Delay       int       `json:"delay"`
	IsPlaylist  bool      `json:"is_playlist"`
	Preview     struct {
		Small    string `json:"small"`
		Medium   string `json:"medium"`
		Large    string `json:"large"`
		Template string `json:"template"`
	} `json:"preview"`
	Channel struct {
		Mature                       bool        `json:"mature"`
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
		ProfileBannerBackgroundColor interface{} `json:"profile_banner_background_color"`
		Partner                      bool        `json:"partner"`
		URL                          string      `json:"url"`
		Views                        int         `json:"views"`
		Followers                    int         `json:"followers"`
	} `json:"channel"`
}

type streamsStream struct {
	Stream Stream `json:"stream"`
}

// StreamSuccessCallback runs on a successfull request and parse using the Stream method
type StreamSuccessCallback func(Stream)

// GetStream request for https://api.twitch.tv/kraken/streams/:channel
func (twitchAPI *TwitchAPI) GetStream(streamName string, onSuccess StreamSuccessCallback, onHTTPError HTTPErrorCallback, onInternalError InternalErrorCallback) {
	var streamsStream streamsStream
	onSuccessfulRequest := func() {
		onSuccess(streamsStream.Stream)
	}
	twitchAPI.Get("/streams/"+streamName, nil, "", &streamsStream, onSuccessfulRequest, onHTTPError, onInternalError)
}
