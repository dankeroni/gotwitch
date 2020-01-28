package gotwitch

import "errors"

var (
	ErrMissingParameters                  = errors.New("missing required parameters")
	ErrCannotPassEmptyStringAsLookupValue = errors.New("cannot pass empty string as lookup value")
)
