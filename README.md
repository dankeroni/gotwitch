# gotwitch [![Build Status](https://travis-ci.org/dankeroni/gotwitch.svg?branch=master)](https://travis-ci.org/dankeroni/gotwitch)

## Example for getting a Stream object
```go
package main

import (
    "fmt"

    "github.com/dankeroni/gotwitch"
)

var api = gotwitch.New("<ClientID>")

func main() {
    api.GetStream("pajlada", onSuccess, onHTTPError, onInternalError)
}

func onSuccess(stream gotwitch.Stream) {
    fmt.Printf("%+v\n", stream)
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

