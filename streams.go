package gotwitch

import (
	"net/http"
	"net/url"
	"time"

	"github.com/pajlada/jsonapi"
)

type streamsListChannel struct {
	Data       []Stream `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// Stream json to struct
type Stream struct {
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

type streamsChannel struct {
	Stream Stream `json:"stream"`
}

// GetStream request for GET https://api.twitch.tv/kraken/streams/:channel
func (twitchAPI *TwitchAPI) GetStream(channelName string, onSuccess func(Stream),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var streamsChannel streamsChannel
	onSuccessfulRequest := func() {
		onSuccess(streamsChannel.Stream)
	}
	twitchAPI.get("/streams/"+channelName, nil, &streamsChannel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// GetStreams request for GET https://api.twitch.tv/helix/streams
// https://dev.twitch.tv/docs/api/reference/#get-streams
func (twitchAPI *TwitchAPI) GetStreams(userIDs []string,
	userLogins []string,
	onSuccess func([]Stream),
	onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) (response *http.Response, err error) {
	var streamsListChannel streamsListChannel
	onSuccessfulRequest := func() {
		onSuccess(streamsListChannel.Data)
	}
	parameters := make(url.Values)
	n := 100
	for _, userID := range userIDs {
		if n == 0 {
			break
		}
		parameters.Add("user_id", userID)
		n++
	}

	for _, userLogin := range userLogins {
		if n == 0 {
			break
		}
		parameters.Add("user_login", userLogin)
		n++
	}

	return twitchAPI.get("/streams", parameters, &streamsListChannel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

func (twitchAPI *TwitchAPI) GetStreamsSimple(userIDs []string, userLogins []string) (data []Stream, response *http.Response, err error) {
	var errorChannel = make(chan error)

	onSuccess := func(r []Stream) {
		data = r
		errorChannel <- nil
	}

	go func() {
		response, err = twitchAPI.GetStreams(userIDs, userLogins, onSuccess, simpleOnHTTPError(errorChannel), simpleOnInternalError(errorChannel))
	}()
	err = <-errorChannel
	return
}
