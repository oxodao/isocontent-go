package specs

import "github.com/oxodao/isocontent-go/AST"

type Specification interface {
	IsSatisfiedBy(AST.Node) bool
	And(Specification) Specification
}
