package builder

import (
	"github.com/oxodao/isocontent-go/AST"
)

type Builder struct {
	Nodes []*Builder
	Type  AST.NodeType
	Data  map[string]interface{}
}

type IsoBlockType struct {
	IsoName   string
	Arguments map[string]interface{}
}

func New(nodeType AST.NodeType, data map[string]interface{}) *Builder {
	return &Builder{
		Nodes: []*Builder{},
		Type:  nodeType,
		Data:  data,
	}
}

func (b *Builder) AddTextNode(text string) *Builder {
	b.Nodes = append(b.Nodes, New(AST.NodeTypeText, map[string]interface{}{
		"text": text,
	}))

	return b
}

func (b *Builder) AddBlockNode(blockType *IsoBlockType) *Builder {
	subBuilder := New(AST.NodeTypeBlock, map[string]interface{}{
		"block_type": blockType.IsoName,
		"arguments":  blockType.Arguments,
	})

	b.Nodes = append(b.Nodes, subBuilder)
	return subBuilder
}

func (b *Builder) GetAST() *AST.Node {

	if b.Type == AST.NodeTypeText {
		text := b.Data["text"].(string)
		return AST.FromText(text)
	}

	if b.Type == AST.NodeTypeBlock {
		blockType := b.Data["block_type"].(string)
		blockArguments := b.Data["arguments"].(map[string]interface{})

		if b.Nodes == nil || len(b.Nodes) == 0 {
			return AST.FromBlockType(blockType, &blockArguments, nil)
		}

		var children []AST.Node
		for _, n := range b.Nodes {
			ast := n.GetAST()
			children = append(children, *ast)
		}

		return AST.FromBlockType(blockType, &blockArguments, &children)
	}

	var children []AST.Node
	for _, n := range b.Nodes {
		node := n.GetAST()
		children = append(children, *node)
	}

	return &AST.Node{
		NodeType: AST.NodeTypeList,
		Children: &children,
	}
}
