package specs

import (
	"github.com/oxodao/isocontent-go/AST"
)

type BlockArgumentMatch struct {
	Key string
	Value interface{}
}

func (ba BlockArgumentMatch) IsSatisfiedBy(node AST.Node) bool {
	if node.NodeType != AST.NodeTypeBlock {
		return false
	}

	if node.Arguments == nil || len(*node.Arguments) == 0 {
		return false
	}

	if arg, ok := (*node.Arguments)[ba.Key]; ok {
		return arg == ba.Value
	}

	return false
}

func (ba BlockArgumentMatch) And(specification Specification) Specification {
	return NewAllMatch(ba, specification)
}

func Argument(key string, val interface{}) BlockArgumentMatch {
	return BlockArgumentMatch{
		Key:   key,
		Value: val,
	}
}