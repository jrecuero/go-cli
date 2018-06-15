package syntax_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/syntax"
)

func TestSyntax_ContentNode(t *testing.T) {
	content := syntax.NewContent("main", "main help", nil)
	cn := syntax.NewContentNode(content.GetLabel(), content)
	g := graph.NewGraph(nil)
	g.AddNode(syntax.ContentNodeToNode(cn))
	node := syntax.NodeToContentNode(g.Root.Children[0])
	if !reflect.DeepEqual(node, cn) {
		t.Errorf("content node mismatch:\n\texp: %#v\n\tgot: %#v\n", cn, node)
	}
}
