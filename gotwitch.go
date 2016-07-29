package gotwitch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type errorResponse struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// TwitchAPI struct
type TwitchAPI struct {
	ClientID string
}

// SuccessCallback runs on a successfull request and parse
type SuccessCallback func()

// HTTPErrorCallback runs on a errored HTTP request
type HTTPErrorCallback func(statusCode int, statusMessage, errorMessage string)

// InternalErrorCallback runs on an internal error
type InternalErrorCallback func(error)

// New instantiates a new TwitchAPI object
func New(clientID string) *TwitchAPI {
	return &TwitchAPI{
		ClientID: clientID,
	}
}

var client = &http.Client{}

// Get can be also used for requests which aren't covered by the library yet
func (twitchAPI *TwitchAPI) Get(url string, parameters url.Values, data interface{}, onSuccess SuccessCallback,
	onHTTPError HTTPErrorCallback, onInternalError InternalErrorCallback) {
	url = "https://api.twitch.tv/kraken" + url + "?" + parameters.Encode()
	request, err := http.NewRequest("GET", url, nil)
	twitchAPI.setHeaders(request)
	response, err := client.Do(request)
	if err != nil {
		onInternalError(err)
		return
	}

	if response.StatusCode != 200 {
		handleHTTPError(response, onHTTPError, onInternalError)
		return
	}

	handleSuccess(response, data, onSuccess, onInternalError)
}

func (twitchAPI *TwitchAPI) setHeaders(request *http.Request) {
	request.Header.Add("Client-ID", twitchAPI.ClientID)
	request.Header.Add("Accept", "application/vnd.twitchtv.v3+json")
}

func handleSuccess(response *http.Response, data interface{}, onSuccess SuccessCallback, onInternalError InternalErrorCallback) {
	body, err := body(response)
	if err != nil {
		onInternalError(err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		onInternalError(err)
		return
	}

	onSuccess()
}

func handleHTTPError(response *http.Response, onHTTPError HTTPErrorCallback, onInternalError InternalErrorCallback) {
	body, err := body(response)
	if err != nil {
		onInternalError(err)
		return
	}

	var errorResponse errorResponse
	err = json.Unmarshal(body, &errorResponse)
	if err != nil {
		onInternalError(err)
		return

	}

	onHTTPError(errorResponse.Status, errorResponse.Message, errorResponse.Error)
}

func body(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return body, err
}
