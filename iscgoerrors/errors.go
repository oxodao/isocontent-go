package iscgoerrors

import "errors"

var (
	NoRenderer = errors.New("no renderers found for given format")
	BadArgument = errors.New("something went wrong rendering the arguments")

	NoParser = errors.New("no parsers found for given format")
	InvalidParserInput = errors.New("invalid parser input")
)
