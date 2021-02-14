package parser

import "github.com/oxodao/isocontent-go/AST"

type Parser interface {
	Parse(interface{}) ([]AST.Node, error)
	SupportsFormat(string) bool
}
