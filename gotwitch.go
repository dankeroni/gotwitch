package gotwitch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	DisplayName string    `json:"display_name"`
	ID          int       `json:"_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
	Links       struct {
		Self string `json:"self"`
	} `json:"_links"`
}

func getBody(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return body, err
}

func DecodeJsonFromUrl(url string, v interface{}) error {
	body, err := getBody(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)
	return err
}

func GetUser(name string) (User, error) {
	var user User
	err := DecodeJsonFromUrl("https://api.twitch.tv/kraken/users/"+name, &user)
	return user, err
}
