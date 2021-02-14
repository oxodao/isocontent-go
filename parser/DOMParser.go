package parser

import (
	"fmt"
	"github.com/oxodao/isocontent-go/AST"
	"github.com/oxodao/isocontent-go/iscgoerrors"
	"golang.org/x/net/html"
	"strings"
)

type DOMParser struct { }

func (dp DOMParser) Parse(input interface{}) ([]AST.Node, error) {
	inputStr, ok := input.(string)
	if !ok {
		return nil, iscgoerrors.InvalidParserInput
	}

	inputStr = `<body>` + inputStr + `</body>`

	doc, err := html.Parse(strings.NewReader(inputStr))
	if err != nil {
		return nil, err
	}

	for node := doc.NextSibling; node != nil; {
		fmt.Println(node.Type)
	}

	return nil, nil
}

func (dp DOMParser) SupportsFormat(format string) bool {
	return format == "html"
}
