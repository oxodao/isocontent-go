package iscgoerrors

import "errors"

var (
	NoRenderer = errors.New("no renderers found for given format")
	BadArgument = errors.New("something went wrong rendering the arguments")
)
