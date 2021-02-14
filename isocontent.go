package isocontent_go

import (
	"github.com/oxodao/isocontent-go/AST"
	"github.com/oxodao/isocontent-go/builder"
	"github.com/oxodao/isocontent-go/iscgoerrors"
	"github.com/oxodao/isocontent-go/parser"
	"github.com/oxodao/isocontent-go/renderer"
)

type Isocontent struct {
	Parsers   []parser.Parser
	Renderers []renderer.Renderer
}

func New() Isocontent {
	return Isocontent{
		Parsers: []parser.Parser{
			parser.DOMParser{},
		},
		Renderers: []renderer.Renderer{
			renderer.NewHTMLRenderer(),
			renderer.JsonRenderer{},
		},
	}
}

func (i *Isocontent) RegisterParser(parser parser.Parser) {
	i.Parsers = append(i.Parsers, parser)
}

func (i *Isocontent) RegisterRenderer(renderer renderer.Renderer) {
	i.Renderers = append(i.Renderers, renderer)
}

func (i *Isocontent) Parse(input interface{}, format string) ([]AST.Node, error) {
	currBuilder := builder.New("", nil)

	for _, p := range i.Parsers {
		if p.SupportsFormat(format) {
			err := p.Parse(currBuilder, input)
			if err != nil {
				return nil, err
			}

			return *currBuilder.GetAST().Children, nil
		}
	}

	return nil, iscgoerrors.NoParser
}

func (i *Isocontent) Render(ast []AST.Node, format string) (interface{}, error) {
	for _, r := range i.Renderers {
		if r.SupportsFormat(format) {
			return r.Render(ast)
		}
	}

	return nil, iscgoerrors.NoRenderer
}

func ExampleAST() []AST.Node {
	return []AST.Node{
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
			Children: &[]AST.Node{
				{
					NodeType: AST.NodeTypeText,
					Value:    "Google :)",
				},
			},
			Arguments: &map[string]interface{}{
				"arguments": map[string]string{
					"href": "https://google.fr/",
				},
			},
		},
	}
}
