package gotwitch

import (
	"fmt"
)

func simpleOnHTTPError(channel chan error) func(code int, statusMessage, errorMessage string) {
	return func(code int, statusMessage, errorMessage string) {
		fmt.Println("ERROR")
		channel <- fmt.Errorf("HTTP Error %d: %s - %s", code, statusMessage, errorMessage)
		fmt.Println("lol")
	}
}

func simpleOnInternalError(channel chan error) func(err error) {
	return func(err error) {
		fmt.Println("INTERNAL ERROR")
		channel <- err
		fmt.Println("lol")
	}
}

func authenticationHeaders(oauthToken string) map[string]string {
	r := make(map[string]string)
	r["Authorization"] = "OAuth " + oauthToken
	return r
}
