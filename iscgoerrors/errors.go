package iscgoerrors

import "errors"

var (
	NoRenderer = errors.New("no renderers found for given format")
)
