package parser

import (
	"errors"
	"github.com/oxodao/isocontent-go/builder"
	"github.com/oxodao/isocontent-go/iscgoerrors"
	"golang.org/x/net/html"
	"strings"
)

type DOMParser struct{}

func (dp DOMParser) Parse(builder *builder.Builder, input interface{}) error {
	inputStr, ok := input.(string)
	if !ok {
		return iscgoerrors.InvalidParserInput
	}

	doc, err := html.Parse(strings.NewReader(inputStr))
	if err != nil {
		return err
	}

	body, err := dp.findBody(doc)
	if err != nil {
		return err
	}

	for node := body.FirstChild; node != nil; node = node.NextSibling {
		dp.parseNode(builder, node)
	}

	return nil
}

func (dp DOMParser) SupportsFormat(format string) bool {
	return format == "html"
}

/** Thanks to https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags **/
func (dp DOMParser) findBody(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)

	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(doc)
	if body != nil {
		return body, nil
	}

	return nil, errors.New("could not find the body, maybe something with the golang.org/x/net/html library")
}

func (dp DOMParser) parseNode(currBuilder *builder.Builder, node *html.Node) {
	var childBuilder *builder.Builder = nil

	switch node.Type {
	case html.TextNode:
		currBuilder.AddTextNode(node.Data)
		return
	case html.ElementNode:
		a := dp.parseBlockType(node)
		childBuilder = currBuilder.AddBlockNode(a)
		break
	default:
		return
	}

	if node.FirstChild == nil {
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dp.parseNode(childBuilder, child)
	}
}

func (dp DOMParser) parseBlockType(node *html.Node) *builder.IsoBlockType {
	switch node.Data {
	case "h1":
		return &builder.IsoBlockType{IsoName: "title", Arguments: map[string]interface{}{"level": 1}}
	case "h2":
		return &builder.IsoBlockType{IsoName: "title", Arguments: map[string]interface{}{"level": 2}}
	case "h3":
		return &builder.IsoBlockType{IsoName: "title", Arguments: map[string]interface{}{"level": 3}}
	case "h4":
		return &builder.IsoBlockType{IsoName: "title", Arguments: map[string]interface{}{"level": 4}}
	case "h5":
		return &builder.IsoBlockType{IsoName: "title", Arguments: map[string]interface{}{"level": 5}}
	case "h6":
		return &builder.IsoBlockType{IsoName: "title", Arguments: map[string]interface{}{"level": 6}}
	case "p":
		return &builder.IsoBlockType{IsoName: "paragraph"}
	case "em":
		return &builder.IsoBlockType{IsoName: "emphasis"}
	case "strong":
		return &builder.IsoBlockType{IsoName: "strong"}
	case "span":
		return &builder.IsoBlockType{IsoName: "inline_text"}
	case "ul":
		return &builder.IsoBlockType{IsoName: "list", Arguments: map[string]interface{}{"ordered": false}}
	case "ol":
		return &builder.IsoBlockType{IsoName: "list", Arguments: map[string]interface{}{"ordered": true}}
	case "li":
		return &builder.IsoBlockType{IsoName: "list_item"}
	case "blockquote":
		return &builder.IsoBlockType{IsoName: "quote"}
	case "br":
		return &builder.IsoBlockType{IsoName: "new_line"}
	case "a":
		args := map[string]string{}
		appendAttribute(node, args, "href")
		return &builder.IsoBlockType{IsoName: "link", Arguments: map[string]interface{}{"arguments": args}}
	case "del":
		return &builder.IsoBlockType{IsoName: "stripped"}
	case "hr":
		return &builder.IsoBlockType{IsoName: "separator"}
	case "sub":
		return &builder.IsoBlockType{IsoName: "subscript"}
	case "sup":
		return &builder.IsoBlockType{IsoName: "supscript"}
	case "code":
		return &builder.IsoBlockType{IsoName: "code"}
	}

	return &builder.IsoBlockType{IsoName: "generic"}
}

func appendAttribute(node *html.Node, args map[string]string, key string) {
	for _, a := range node.Attr {
		if a.Key == key {
			args[key] = a.Val
			return
		}
	}
}