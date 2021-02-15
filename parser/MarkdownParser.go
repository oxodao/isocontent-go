package parser

import (
	"fmt"
	"github.com/oxodao/isocontent-go/builder"
	"github.com/oxodao/isocontent-go/iscgoerrors"
	"github.com/yuin/goldmark"
	mdast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
)

type MarkdownParser struct {
	goldmark goldmark.Markdown
}

func NewMarkdownParser() MarkdownParser {
	mk := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
		),
	)
	return MarkdownParser{
		goldmark: mk,
	}
}

func (md MarkdownParser) Parse(builder *builder.Builder, input interface{}) error {
	inputStr, ok := input.(string)
	if !ok {
		return iscgoerrors.InvalidParserInput
	}

	doc := md.goldmark.Parser().Parse(text.NewReader([]byte(inputStr)))
	for child := doc.FirstChild(); child != nil; child = child.NextSibling() {
		md.parseNode(builder, &child)
	}

	return nil
}

func (md MarkdownParser) SupportsFormat(format string) bool {
	return format == "markdown"
}

func (md MarkdownParser) parseNode(currBuilder *builder.Builder, node *mdast.Node) {

	var childBuilder *builder.Builder = nil


	switch (*node).Type() {
	case mdast.TypeInline:
		//panic("INLINE: " + (*node).Kind().String())
	case mdast.TypeBlock:
		fmt.Println("Block child (", (*node).Kind(), ")")
		fmt.Println((*node).Lines())
		break
	default:
		return
	}

	if (*node).FirstChild() == nil {
		return
	}

	for child := (*node).FirstChild(); child != nil; child = child.NextSibling() {
		md.parseNode(childBuilder, &child)
	}
}
