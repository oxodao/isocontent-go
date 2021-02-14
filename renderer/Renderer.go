package renderer

import "github.com/oxodao/isocontent-go/AST"

type Renderer interface {
	Render([]AST.Node) (interface{}, error)
	SupportsFormat(string) bool
}
