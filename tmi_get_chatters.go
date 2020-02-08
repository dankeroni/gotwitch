package gotwitch

import "fmt"

// Chatters json to struct
type Chatters struct {
	Moderators []string `json:"moderators"`
	Staff      []string `json:"staff"`
	Admins     []string `json:"admins"`
	GlobalMods []string `json:"global_mods"`
	Viewers    []string `json:"viewers"`
}

func (a *TwitchAPITMI) GetChatters(login string) (Chatters, error) {
	type Result struct {
		Links        struct{} `json:"_links"`
		ChatterCount int      `json:"chatter_count"`
		Chatters     Chatters `json:"chatters"`
	}

	result := Result{}

	request := a.c.R().
		SetResult(&result)

	resp, err := request.
		Get("streams")

	if resp.IsError() {
		e := resp.Error().(*TMIError)
		return result.Chatters, fmt.Errorf("TMI API error code %s", string(*e))
	}

	return result.Chatters, err
}
