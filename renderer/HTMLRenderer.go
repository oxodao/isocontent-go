package renderer

import (
	"fmt"
	"github.com/oxodao/isocontent-go/AST"
	"github.com/oxodao/isocontent-go/iscgoerrors"
	"github.com/oxodao/isocontent-go/specs"
	"strings"
)

type HTMLRenderer struct {
	tags []tag
}

type tag struct {
	Specification specs.Specification
	HTMLElement string
}

func NewHTMLRenderer() HTMLRenderer {
	return HTMLRenderer{
		tags: []tag{
			{ specs.Type("paragraph"), "p" },
			{ specs.Type("inline_text"), "span" },
			{ specs.Type("emphasis"), "em" },
			{ specs.Type("strong"), "strong" },
			{ specs.Type("generic"), "span" },
			{ specs.Type("list").And(specs.Argument("ordered", false)), "ul" },
			{ specs.Type("list").And(specs.Argument("ordered", true)), "ol" },
			{ specs.Type("list_item"), "li" },
			{ specs.Type("title").And(specs.Argument("level", 1)), "h1"},
			{ specs.Type("title").And(specs.Argument("level", 2)), "h2"},
			{ specs.Type("title").And(specs.Argument("level", 3)), "h3"},
			{ specs.Type("title").And(specs.Argument("level", 4)), "h4"},
			{ specs.Type("title").And(specs.Argument("level", 5)), "h5"},
			{ specs.Type("title").And(specs.Argument("level", 6)), "h6"},
			{ specs.Type("quote"), "blockquote"},
			{ specs.Type("new_line"), "br"},
			{ specs.Type("link"), "a"},
			{ specs.Type("striped"), "del"},
			{ specs.Type("separator"), "hr"},
			{ specs.Type("subscript"), "sub"},
			{ specs.Type("superscript"), "sup"},
			{ specs.Type("code"), "code"},
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
	for _, t := range h.tags {
		if t.Specification.IsSatisfiedBy(node) {
			tagName = t.HTMLElement
		}
	}

	arguments := ""
	if args, ok := node.Arguments["arguments"]; ok {
		argsArr, err := args.(map[string]string)
		if !err {
			return "", iscgoerrors.BadArgument
		}

		var renderedArgs []string
		for k, v := range argsArr {
			renderedArgs = append(renderedArgs, fmt.Sprintf(`%v="%v"`, k, v))
		}

		arguments = " " + strings.Join(renderedArgs, " ")
	}

	if node.Children == nil {
		return fmt.Sprintf("<%v%v />", tagName, arguments), nil
	}

	childrenRender, err := h.Render(*node.Children)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("<%v%v>%v</%v>", tagName, arguments, childrenRender, tagName), nil
}

