package isocontent_go

import (
	"github.com/oxodao/isocontent-go/AST"
	"github.com/oxodao/isocontent-go/iscgoerrors"
	"github.com/oxodao/isocontent-go/renderer"
)

type Isocontent struct {
	Renderers []renderer.Renderer
}

func NewIsocontent() Isocontent {
	return Isocontent{
		Renderers: []renderer.Renderer{
			renderer.NewHTMLRenderer(),
			renderer.JsonRenderer{},
		},
	}
}

func (i *Isocontent) RegisterRenderer(renderer renderer.Renderer) {
	i.Renderers = append(i.Renderers, renderer)
}

func (i *Isocontent) Render(ast []AST.Node, format string) (interface{}, error) {
	for _, r := range i.Renderers {
		if r.SupportsFormat(format) {
			return r.Render(ast)
		}
	}

	return nil, iscgoerrors.NoRenderer
}

func ExampleAST() AST.NodeList {
	return AST.NodeList{
		Nodes: []AST.Node{
			{
				NodeType:  AST.NodeTypeBlock,
				BlockType: "paragraph",
				Children: &[]AST.Node{
					{
						NodeType: AST.NodeTypeText,
						Value:    "Text",
					},
				},
				Arguments: &map[string]interface{}{},
			},

			{
				NodeType:  AST.NodeTypeBlock,
				BlockType: "list",
				Children: &[]AST.Node{
					{
						NodeType:  AST.NodeTypeBlock,
						BlockType: "list_item",
						Children: &[]AST.Node{
							{
								NodeType: AST.NodeTypeText,
								Value:    "Un",
							},
						},
						Arguments: &map[string]interface{}{},
					},
					{
						NodeType:  AST.NodeTypeBlock,
						BlockType: "list_item",
						Children: &[]AST.Node{
							{
								NodeType: AST.NodeTypeText,
								Value:    "Deux",
							},
						},
						Arguments: &map[string]interface{}{},
					},
					{
						NodeType:  AST.NodeTypeBlock,
						BlockType: "list_item",
						Children: &[]AST.Node{
							{
								NodeType: AST.NodeTypeText,
								Value:    "Trois",
							},
						},
						Arguments: &map[string]interface{}{},
					},
				},
				Arguments: &map[string]interface{}{
					"ordered": false,
				},
			},

			{
				NodeType:  AST.NodeTypeBlock,
				BlockType: "list",
				Children: &[]AST.Node{
					{
						NodeType:  AST.NodeTypeBlock,
						BlockType: "list_item",
						Children: &[]AST.Node{
							{
								NodeType: AST.NodeTypeText,
								Value:    "Un",
							},
						},
						Arguments: &map[string]interface{}{},
					},
					{
						NodeType:  AST.NodeTypeBlock,
						BlockType: "list_item",
						Children: &[]AST.Node{
							{
								NodeType: AST.NodeTypeText,
								Value:    "Deux",
							},
						},
						Arguments: &map[string]interface{}{},
					},
					{
						NodeType:  AST.NodeTypeBlock,
						BlockType: "list_item",
						Children: &[]AST.Node{
							{
								NodeType: AST.NodeTypeText,
								Value:    "Trois",
							},
						},
						Arguments: &map[string]interface{}{},
					},
				},
				Arguments: &map[string]interface{}{
					"ordered": true,
				},
			},

			{
				NodeType:  AST.NodeTypeBlock,
				BlockType: "link",
				Children:  &[]AST.Node{
					{
						NodeType: AST.NodeTypeText,
						Value: "Google :)",
					},
				},
				Arguments: &map[string]interface{}{
					"arguments": map[string]string{
						"href": "https://google.fr/",
					},
				},
			},
		},
	}
}
