package gotwitch

import (
	"net/url"
	"time"

	"github.com/pajlada/jsonapi"
)

// Block json to struct
type Block struct {
	ID        int       `json:"_id"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

type usersUserBlocks struct {
	Blocks []Block `json:"blocks"`
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
