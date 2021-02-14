package renderer

import (
	"fmt"
	"github.com/oxodao/isocontent-go/AST"
)

type HTMLRenderer struct {
	tags map[string]string
}

func NewHTMLRenderer() HTMLRenderer {
	return HTMLRenderer{
		tags: map[string]string{
			"paragraph": "p",
			"list": "ul",
			"list_item": "li",
		},
	}
}

func (h HTMLRenderer) Render(ast []AST.Node) (interface{}, error) {

	memo := ""

	for _, n := range ast {
		if n.NodeType == AST.NodeTypeText {
			memo += n.Value
		} else if n.NodeType == AST.NodeTypeBlock {
			renderedNode, err := h.renderNode(n)
			if err != nil {
				return nil, err
			}

			memo += renderedNode
		}
	}

	return memo, nil
}

func (h HTMLRenderer) SupportsFormat(format string) bool {
	return format == "html"
}

func (h HTMLRenderer) renderNode(node AST.Node) (string, error) {
	tagName := "span"
	for k, v := range h.tags {
		if node.BlockType == k {
			tagName = v
		}
	}

	if node.Children == nil {
		return fmt.Sprintf("<%v />", tagName), nil
	}

	childrenRender, err := h.Render(*node.Children)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<%v>%v</%v>", tagName, childrenRender, tagName), nil
}

