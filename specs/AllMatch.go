package specs

import "github.com/oxodao/isocontent-go/AST"

type AllMatch struct {
	Specifications []Specification
}

func NewAllMatch(specs ...Specification) AllMatch {
	return AllMatch{Specifications: specs}
}

func (am AllMatch) IsSatisfiedBy(node AST.Node) bool {
	for _, s := range am.Specifications {
		if !s.IsSatisfiedBy(node) {
			return false
		}
	}

	return true
}

func (am AllMatch) And(specification Specification) Specification {
	return nil
}