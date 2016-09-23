package gotwitch

import (
	"github.com/dankeroni/jsonapi"
	"time"
)

// Post json to struct
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Deleted   bool      `json:"deleted"`
	Body      string    `json:"body"`
	Emotes    []struct {
		ID    int `json:"id"`
		Start int `json:"start"`
		End   int `json:"end"`
		Set   int `json:"set"`
	} `json:"emotes"`
	Embeds []struct {
		Type         string    `json:"type"`
		TwitchType   string    `json:"twitch_type"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		AuthorName   string    `json:"author_name"`
		ThumbnailURL string    `json:"thumbnail_url"`
		Game         string    `json:"game"`
		PlayerHTML   string    `json:"player_html"`
		CreatedAt    time.Time `json:"created_at"`
		RequestURL   string    `json:"request_url"`
		VideoLength  int       `json:"video_length"`
		ProviderName string    `json:"provider_name"`
	} `json:"embeds"`
	Reactions map[string]struct {
		Emote   string `json:"emote"`
		Count   int    `json:"count"`
		UserIds []int  `json:"user_ids"`
	} `json:"reactions"`
	User     User `json:"user"`
	Comments struct {
		Total    int           `json:"_total"`
		Cursor   string        `json:"_cursor"`
		Comments []interface{} `json:"comments"`
	} `json:"comments"`
	Permissions struct {
		CanReply    bool `json:"can_reply"`
		CanModerate bool `json:"can_moderate"`
		CanDelete   bool `json:"can_delete"`
	} `json:"permissions"`
}

// Posts json to struct
type Posts struct {
	Total  int    `json:"_total"`
	Cursor string `json:"_cursor"`
	Topic  string `json:"_topic"`
	Posts  []Post `json:"posts"`
}

// GetPost request for GET https://api.twitch.tv/kraken/feed/:channel/posts/:id
func (twitchAPI *TwitchAPI) GetPost(channelName string, id string, onSuccess func(Post),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var post Post
	onSuccessfulRequest := func() {
		onSuccess(post)
	}
	twitchAPI.Get("/feed/"+channelName+"/posts/"+id, nil, &post, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
