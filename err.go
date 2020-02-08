package gotwitch

import "errors"

var (
	// ErrNotAuthenticated is returned by functions that requires user-authentication
	ErrNotAuthenticated = errors.New("not authenticated")

	// ErrMissingParameters is returned if parameters are not passed to a function that requires parameters
	ErrMissingParameters = errors.New("missing required parameters")

	// ErrCannotPassEmptyStringAsLookupValue is returned if
	ErrCannotPassEmptyStringAsLookupValue = errors.New("cannot pass empty string as lookup value")

	// ErrMissingID is returned if the ID class was not initialized properly for the helix api object
	ErrMissingID = errors.New("missing id class in helix class")

	// ErrMissingClientSecret is returned if a function requires a client secret to be set but no client secret has been set
	ErrMissingClientSecret = errors.New("missing client secret")
)
