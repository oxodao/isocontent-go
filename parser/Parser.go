package parser

import (
	"github.com/oxodao/isocontent-go/builder"
)

type Parser interface {
	Parse(*builder.Builder, interface{}) error
	SupportsFormat(string) bool
}
