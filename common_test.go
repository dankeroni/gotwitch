package gotwitch

import "os"

var Twitch *TwitchAPI
var TwitchV5 *TwitchAPI
var oauthToken string

func onInternalError(err error) {
	panic(err)
}

func onHTTPError(code int, statusMessage, errorMessage string) {
	panic(statusMessage)
}

func init() {
	Twitch = New(os.Getenv("GOTWITCH_CLIENT_ID"))
	Twitch.Credentials.ClientSecret = os.Getenv("GOTWITCH_CLIENT_SECRET")
	oauthToken = os.Getenv("GOTWITCH_OAUTH_TOKEN")

	TwitchV5 = NewV5(os.Getenv("GOTWITCH_CLIENT_ID"))
	TwitchV5.Credentials.ClientSecret = os.Getenv("GOTWITCH_CLIENT_SECRET")
	oauthToken = os.Getenv("GOTWITCH_OAUTH_TOKEN")
}
