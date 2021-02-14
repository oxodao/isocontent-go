package specs

import (
	"github.com/oxodao/isocontent-go/AST"
)

type BlockTypeMatch struct {
	BlockType string
}

func (bt BlockTypeMatch) IsSatisfiedBy(node AST.Node) bool {
	return node.NodeType == AST.NodeTypeBlock && node.BlockType == bt.BlockType
}

func (bt BlockTypeMatch) And(specification Specification) Specification {
	return NewAllMatch(bt, specification)
}

func Type(blockType string) BlockTypeMatch {
	return BlockTypeMatch{BlockType: blockType}
}