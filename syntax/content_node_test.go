package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/syntax"
)

func TestContentNode_ContentNode(t *testing.T) {
	content := syntax.NewContent("main", "main help", nil)
	cn := &syntax.ContentNode{
		graph.NewNode("main", content),
	}
	if cn.GetContent() != content {
		t.Errorf("new content mismatch:\n\texp: %#v\n\tgot: %#v\n", content, cn.GetContent())
	}
}

func TestContentNode_Match(t *testing.T) {
	content := syntax.NewContent("main", "main help", nil)
	cn := &syntax.ContentNode{
		graph.NewNode("main", content),
	}
	if index, ok := cn.Match(nil, []string{"main", "zero"}, 0); index != 1 || ok != true {
		t.Errorf("match mismatch:\n\texp: %d %v\n\tgot: %d %v\n", 1, true, index, ok)
	}
	if index, ok := cn.Match(nil, []string{"main", "zero"}, 1); index != 1 || ok != false {
		t.Errorf("match mismatch:\n\texp: %d %v\n\tgot: %d %v\n", 1, false, index, ok)
	}
}
