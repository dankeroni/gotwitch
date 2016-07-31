# gotwitch [![Build Status](https://travis-ci.org/dankeroni/gotwitch.svg?branch=master)](https://travis-ci.org/dankeroni/gotwitch)

## Example for unimplemented requests
```go
package main

import (
	"fmt"

	"github.com/dankeroni/gotwitch"
	"net/url"
)

// GamesTop https://api.twitch.tv/kraken/games/top (https://mholt.github.io/json-to-go/)
// The oauthToken is optional and can be replaced by ""
type GamesTop struct {
	Total int `json:"_total"`
	Links struct {
		Self string `json:"self"`
		Next string `json:"next"`
	} `json:"_links"`
	Top []struct {
		Game struct {
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			ID          int    `json:"_id"`
			GiantbombID int    `json:"giantbomb_id"`
			Box         struct {
				Large    string `json:"large"`
				Medium   string `json:"medium"`
				Small    string `json:"small"`
				Template string `json:"template"`
			} `json:"box"`
			Logo struct {
				Large    string `json:"large"`
				Medium   string `json:"medium"`
				Small    string `json:"small"`
				Template string `json:"template"`
			} `json:"logo"`
			Links struct {
			} `json:"_links"`
		} `json:"game"`
		Viewers  int `json:"viewers"`
		Channels int `json:"channels"`
	} `json:"top"`
}

var api = gotwitch.New("<ClientID>")

func main() {
	var gamesTop GamesTop
	requestParameters := url.Values{}
	requestParameters.Add("limit", "1")
	requestParameters.Add("offset", "4")
	onSuccess := func() {
		fmt.Printf("%+v", gamesTop)
	}
	api.Get("/games/top", requestParameters, "<oauthToken>", &gamesTop, onSuccess, onHTTPError, onInternalError)
}

func onHTTPError(statusCode int, statusMessage, errorMessage string) {
	fmt.Println("statusCode:", statusCode)
	fmt.Println("statusMessage:", statusMessage)
	fmt.Println("errorMessage:", errorMessage)
}

func onInternalError(err error) {
	fmt.Println("internalError:", err)
}
```
