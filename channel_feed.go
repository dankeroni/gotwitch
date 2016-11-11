package gotwitch

import (
	"github.com/dankeroni/jsonapi"
	"net/url"
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
		UserIDs []int  `json:"user_ids"`
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

// SharedPost json to struct
type SharedPost struct {
	Tweet string `json:"tweet"`
	Post  Post   `json:"post"`
}

// Posts json to struct
type Posts struct {
	Total  int    `json:"_total"`
	Cursor string `json:"_cursor"`
	Topic  string `json:"_topic"`
	Posts  []Post `json:"posts"`
}

// Reaction json to struct
type Reaction struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	EmoteID   string    `json:"emote_id"`
	User      User      `json:"user"`
}

// GetPost request for GET https://api.twitch.tv/kraken/feed/:channel/posts/:id
func (twitchAPI *TwitchAPI) GetPost(postID, channelName string, onSuccess func(Post),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var post Post
	onSuccessfulRequest := func() {
		onSuccess(post)
	}
	twitchAPI.Get("/feed/"+channelName+"/posts/"+postID, nil, &post, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// AuthenticatedGetPost request for GET https://api.twitch.tv/kraken/feed/:channel/posts/:id
func (twitchAPI *TwitchAPI) AuthenticatedGetPost(oauthToken, postID, channelName string,
	onSuccess func(Post), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var post Post
	onSuccessfulRequest := func() {
		onSuccess(post)
	}
	twitchAPI.AuthenticatedGet("/feed/"+channelName+"/posts/"+postID, nil, oauthToken, &post,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// PostPost request for POST https://api.twitch.tv/kraken/feed/:channel/posts
func (twitchAPI *TwitchAPI) PostPost(oauthToken, content, channelName string, parameters url.Values,
	onSuccess func(SharedPost), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var sharedPost SharedPost
	onSuccessfulRequest := func() {
		onSuccess(sharedPost)
	}
	if parameters == nil {
		parameters = url.Values{}
	}
	parameters.Add("content", content)
	twitchAPI.AuthenticatedPost("/feed/"+channelName+"/posts", parameters, oauthToken, nil,
		&sharedPost, onSuccessfulRequest, onHTTPError, onInternalError)
}

// DeletePost request for DELETE https://api.twitch.tv/kraken/feed/:channel/posts/:id
func (twitchAPI *TwitchAPI) DeletePost(oauthToken, postID, channelName string, onSuccess func(),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	twitchAPI.AuthenticatedDelete("/feed/"+channelName+"/posts/"+postID, nil, oauthToken, nil,
		onSuccess, onHTTPError, onInternalError)
}

// GetPosts request for GET https://api.twitch.tv/kraken/feed/:channel/posts
func (twitchAPI *TwitchAPI) GetPosts(channelName string, parameters url.Values,
	onSuccess func(Posts), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var posts Posts
	onSuccessfulRequest := func() {
		onSuccess(posts)
	}
	twitchAPI.Get("/feed/"+channelName+"/posts", parameters, &posts, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// AuthenticatedGetPosts request for GET https://api.twitch.tv/kraken/feed/:channel/posts
func (twitchAPI *TwitchAPI) AuthenticatedGetPosts(oauthToken, channelName string,
	parameters url.Values, onSuccess func(Posts), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var posts Posts
	onSuccessfulRequest := func() {
		onSuccess(posts)
	}
	twitchAPI.AuthenticatedGet("/feed/"+channelName+"/posts", parameters, oauthToken, &posts,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// PostReaction request for POST https://api.twitch.tv/kraken/feed/:channel/posts/:id/reactions
func (twitchAPI *TwitchAPI) PostReaction(oauthToken, emoteID, postID, channelName string,
	onSuccess func(Reaction), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var reaction Reaction
	onSuccessfulRequest := func() {
		onSuccess(reaction)
	}
	parameters := url.Values{}
	parameters.Add("emote_id", emoteID)
	twitchAPI.AuthenticatedPost("/feed/"+channelName+"/posts/"+postID+"/reactions", parameters,
		oauthToken, nil, &reaction, onSuccessfulRequest, onHTTPError, onInternalError)
}

// DeleteReaction request for DELETE https://api.twitch.tv/kraken/feed/:channel/posts/:id/reactions
func (twitchAPI *TwitchAPI) DeleteReaction(oauthToken, emoteID, postID, channelName string,
	onSuccess func(), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	parameters := url.Values{}
	parameters.Add("emote_id", emoteID)
	twitchAPI.AuthenticatedDelete("/feed/"+channelName+"/posts/"+postID+"/reactions", parameters,
		oauthToken, nil, onSuccess, onHTTPError, onInternalError)
}
