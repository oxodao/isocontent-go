package renderer

import (
	"encoding/json"
	"github.com/oxodao/isocontent-go/AST"
)

type JsonRenderer struct {}

func (j JsonRenderer) Render(ast []AST.Node) (interface{}, error) {

	jsonAst, err := json.Marshal(ast)
	if err != nil {
		return nil, err
	}

	return string(jsonAst), nil
}

func (j JsonRenderer) SupportsFormat(format string) bool {
	return format == "json"
}
