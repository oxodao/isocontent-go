package AST

import (
	"fmt"
	"testing"

	"github.com/oxodao/isocontent-go/test"
)

func TestCreateTextNode(t *testing.T) {
	node := FromText("Test")
	test.AssertEqual(t, "text", node.NodeType, fmt.Sprintf("Generated Text node should have a 'text' blocktype (got %v)", node.NodeType))
	test.AssertEqual(t, "Test", node.Value, fmt.Sprintf("Generated Text node content do not match (Expected 'Test' got '%v')", node.Value))
}
