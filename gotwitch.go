package gotwitch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

// RequestParameters for the request
type RequestParameters map[string]string

// SuccessCallback runs on a successfull request and parse
type SuccessCallback func(data interface{})

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

// Get can be also used for requests which aren't covered by the library yet
func (twitchAPI *TwitchAPI) Get(url string, parameters RequestParameters, dataType interface{}, onSuccess SuccessCallback,
	onHTTPError HTTPErrorCallback, onInternalError InternalErrorCallback) {
	url = url + "?"
	for _, parameter := range parameters {
		url = url + parameter + "=" + parameters[parameter]
	}
	response, err := http.Get(url)
	if err != nil {
		handleHTTPError(response, onHTTPError, onInternalError)
		return
	}

	handleSuccess(response, dataType, onSuccess, onInternalError)
}

func handleSuccess(response *http.Response, dataType interface{}, onSuccess SuccessCallback, onInternalError InternalErrorCallback) {
	body, err := body(response)
	if err != nil {
		onInternalError(err)
		return
	}

	err = json.Unmarshal(body, &dataType)
	if err != nil {
		onInternalError(err)
		return
	}

	onSuccess(dataType)
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
