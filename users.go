package gotwitch

import (
	"github.com/dankeroni/jsonapi"
	"net/url"
	"time"
)

// User json to struct
type User struct {
	DisplayName string    `json:"display_name"`
	ID          int       `json:"_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
}

// Block json to struct
type Block struct {
	ID        int       `json:"_id"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

type usersUserBlocks struct {
	Blocks []Block `json:"blocks"`
}

// Follow json to struct
type Follow struct {
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
	Channel       Channel   `json:"channel"`
}

// Follows json to struct
type Follows struct {
	Follows []Follow `json:"follows"`
	Total   int      `json:"_total"`
}

// GetBlocks request for GET https://api.twitch.tv/kraken/users/:user/blocks
func (twitchAPI *TwitchAPI) GetBlocks(oauthToken, user string, parameters url.Values,
	onSuccess func([]Block), onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var usersUserBlocks usersUserBlocks
	onSuccessfulRequest := func() {
		onSuccess(usersUserBlocks.Blocks)
	}
	twitchAPI.AuthenticatedGet("/users/"+user+"/blocks", parameters, oauthToken, &usersUserBlocks,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// GetBlock request for GET https://api.twitch.tv/kraken/users/:user/blocks/:target
func (twitchAPI *TwitchAPI) GetBlock(oauthToken, user, target string, onSuccess func(Block),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var block Block
	onSuccessfulRequest := func() {
		onSuccess(block)
	}
	twitchAPI.AuthenticatedGet("/users/"+user+"/blocks/"+target, nil, oauthToken, &block,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// PutBlock request for PUT https://api.twitch.tv/kraken/users/:user/blocks/:target
func (twitchAPI *TwitchAPI) PutBlock(oauthToken, user, target string, onSuccess func(Block),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var block Block
	onSuccessfulRequest := func() {
		onSuccess(block)
	}
	twitchAPI.AuthenticatedPut("/users/"+user+"/blocks/"+target, nil, oauthToken, nil, &block,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// DeleteBlock request for DELETE https://api.twitch.tv/kraken/users/:user/blocks/:target
func (twitchAPI *TwitchAPI) DeleteBlock(oauthToken, user, target string, onSuccess func(),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	twitchAPI.AuthenticatedDelete("/users/"+user+"/blocks/"+target, nil, oauthToken, nil, onSuccess,
		onHTTPError, onInternalError)
}

// GetFollows request for GET https://api.twitch.tv/kraken/users/:user/follows/channels
func (twitchAPI *TwitchAPI) GetFollows(user string, parameters url.Values, onSuccess func(Follows),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var follows Follows
	onSuccessfulRequest := func() {
		onSuccess(follows)
	}
	twitchAPI.Get("/users/"+user+"/follows/channels", parameters, &follows,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// GetFollow request for GET https://api.twitch.tv/kraken/users/:user/follows/channels/:target
func (twitchAPI *TwitchAPI) GetFollow(user, target string, onSuccess func(Follow),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var follow Follow
	onSuccessfulRequest := func() {
		onSuccess(follow)
	}
	twitchAPI.Get("/users/"+user+"/follows/channels/"+target, nil, &follow, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// PutFollow request for PUT https://api.twitch.tv/kraken/users/:user/follows/channels/:target
func (twitchAPI *TwitchAPI) PutFollow(oauthToken, user, target string, onSuccess func(Follow),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var follow Follow
	onSuccessfulRequest := func() {
		onSuccess(follow)
	}
	twitchAPI.AuthenticatedPut("/users/"+user+"/follows/channels/"+target, nil, oauthToken, nil, &follow,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// DeleteFollow request for DELETE https://api.twitch.tv/kraken/users/:user/follows/channels/:target
func (twitchAPI *TwitchAPI) DeleteFollow(oauthToken, user, target string, onSuccess func(),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	twitchAPI.AuthenticatedDelete("/users/"+user+"/follows/channels/"+target, nil, oauthToken, nil, onSuccess,
		onHTTPError, onInternalError)
}
