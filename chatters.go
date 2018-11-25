package gotwitch

import "github.com/pajlada/jsonapi"

// Chatters json to struct
type Chatters struct {
	Moderators []string `json:"moderators"`
	Staff      []string `json:"staff"`
	Admins     []string `json:"admins"`
	GlobalMods []string `json:"global_mods"`
	Viewers    []string `json:"viewers"`
}

// ChattersResponse json to struct
type ChattersResponse struct {
	Links        struct{} `json:"_links"`
	ChatterCount int      `json:"chatter_count"`
	Chatters     Chatters `json:"chatters"`
}

func GetChatters(channelName string, onSuccess func(Chatters), onHTTPError jsonapi.HTTPErrorCallback,
	onInternalError jsonapi.InternalErrorCallback) {
	var chattersResponse ChattersResponse
	onSuccessfulRequest := func() {
		onSuccess(chattersResponse.Chatters)
	}

	tmiAPI := jsonapi.JSONAPI{
		BaseURL: "https://tmi.twitch.tv",
	}

	tmiAPI.Get("/group/user/"+channelName+"/chatters", nil, &chattersResponse, onSuccessfulRequest, onHTTPError, onInternalError)
}
