package gotwitch

import (
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

// BlockSuccessCallback runs on a successfull request and parse using the Block/PutBlock method
type BlockSuccessCallback func(Block)

// BlocksSuccessCallback runs on a successfull request and parse using the Blocks method
type BlocksSuccessCallback func([]Block)

// Blocks request for https://api.twitch.tv/kraken/users/:user/blocks
func (twitchAPI *TwitchAPI) Blocks(oauthToken, user string, parameters url.Values,
	onSuccess BlocksSuccessCallback, onHTTPError HTTPErrorCallback,
	onInternalError InternalErrorCallback) {
	var usersUserBlocks usersUserBlocks
	onSuccessfulRequest := func() {
		onSuccess(usersUserBlocks.Blocks)
	}
	twitchAPI.Get("/users/"+user+"/blocks", parameters, oauthToken, &usersUserBlocks,
		onSuccessfulRequest, onHTTPError, onInternalError)
}

// Block request for https://api.twitch.tv/kraken/users/:user/blocks/:target
func (twitchAPI *TwitchAPI) Block(oauthToken, user, target string, onSuccess BlockSuccessCallback,
	onHTTPError HTTPErrorCallback, onInternalError InternalErrorCallback) {
	var block Block
	onSuccessfulRequest := func() {
		onSuccess(block)
	}
	twitchAPI.Get("/users/"+user+"/blocks/"+target, nil, oauthToken, &block, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// PutBlock request for https://api.twitch.tv/kraken/users/:user/blocks/:target
func (twitchAPI *TwitchAPI) PutBlock(oauthToken, user, target string,
	onSuccess BlockSuccessCallback, onHTTPError HTTPErrorCallback,
	onInternalError InternalErrorCallback) {
	var block Block
	onSuccessfulRequest := func() {
		onSuccess(block)
	}
	baseURL := "/users/" + user + "/blocks/" + target
	twitchAPI.Put(baseURL, nil, oauthToken, nil, &block, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// DeleteBlock request for https://api.twitch.tv/kraken/users/:user/blocks/:target
func (twitchAPI *TwitchAPI) DeleteBlock(oauthToken, user, target string,
	onSuccess func(), onHTTPError HTTPErrorCallback,
	onInternalError InternalErrorCallback) {
	twitchAPI.Delete("/users/"+user+"/blocks/"+target, nil, oauthToken, nil, onSuccess,
		onHTTPError, onInternalError)
}
