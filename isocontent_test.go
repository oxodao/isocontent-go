package isocontent_go

import (
	"github.com/oxodao/isocontent-go/AST"
	"testing"
)

func TestErrorUsingInexistentParser(t *testing.T) {
	iscgo := New()
	_, err := iscgo.Parse("<p>simple paragraph</p>", "inexistent_parser")
	if err == nil {
		t.Errorf("isocontent should not be able to parse with an unexisting parser")
	}
}

func TestErrorUsingInexistentRenderer(t *testing.T) {
	iscgo := New()
	_, err := iscgo.Render([]AST.Node{}, "inexistent_renderer")
	if err == nil {
		t.Errorf("isocontent should not be able to render with an unexisting renderer")
	}
}
