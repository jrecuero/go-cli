package syntax_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestGraph_Node ensures the node structure works properly
func TestGraph_Node(t *testing.T) {
	var tests = []struct {
		n   *syntax.Node
		exp *syntax.Node
	}{
		{
			n: syntax.NewNode("test", "test"),
			exp: &syntax.Node{
				ID:       1,
				Name:     "test",
				Label:    "test",
				Children: nil,
				IsRoot:   false,
				IsSink:   false,
				IsStart:  false,
				IsEnd:    false,
				IsLoop:   false,
				IsJoint:  false,
				InPath:   false,
				BlockID:  -1,
			},
		},
		{
			n: syntax.NewJoint("joint", "joint"),
			exp: &syntax.Node{
				ID:       2,
				Name:     "joint",
				Label:    "joint",
				Children: nil,
				IsRoot:   false,
				IsSink:   false,
				IsStart:  false,
				IsEnd:    false,
				IsLoop:   false,
				IsJoint:  true,
				InPath:   false,
				BlockID:  -1,
			},
		},
		{
			n: syntax.NewRoot(),
			exp: &syntax.Node{
				ID:       3,
				Name:     "ROOT",
				Label:    "ROOT",
				Children: nil,
				IsRoot:   true,
				IsSink:   false,
				IsStart:  false,
				IsEnd:    false,
				IsLoop:   false,
				IsJoint:  true,
				InPath:   false,
				BlockID:  -1,
			},
		},
		{
			n: syntax.NewStart(4),
			exp: &syntax.Node{
				ID:       4,
				Name:     "START",
				Label:    "START",
				Children: nil,
				IsRoot:   false,
				IsSink:   false,
				IsStart:  true,
				IsEnd:    false,
				IsLoop:   false,
				IsJoint:  true,
				InPath:   false,
				BlockID:  4,
			},
		},
		{
			n: syntax.NewEnd(5),
			exp: &syntax.Node{
				ID:       5,
				Name:     "END",
				Label:    "END",
				Children: nil,
				IsRoot:   false,
				IsSink:   false,
				IsStart:  false,
				IsEnd:    true,
				IsLoop:   false,
				IsJoint:  true,
				InPath:   false,
				BlockID:  5,
			},
		},
		{
			n: syntax.NewLoop(6),
			exp: &syntax.Node{
				ID:       6,
				Name:     "LOOP",
				Label:    "LOOP",
				Children: nil,
				IsRoot:   false,
				IsSink:   false,
				IsStart:  false,
				IsEnd:    false,
				IsLoop:   true,
				IsJoint:  true,
				InPath:   false,
				BlockID:  6,
			},
		},
	}

	for i, tt := range tests {
		if !reflect.DeepEqual(tt.n, tt.exp) {
			t.Errorf("%d. node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.exp, tt.n)
		}
	}

	n := syntax.NewNode("main", "main")
	c1 := syntax.NewNode("child-1", "child-1")
	if n.AddChild(c1) == false {
		t.Errorf("add child operation failed")
	}
	if len(n.Children) != 1 {
		t.Errorf("length mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", 1, len(n.Children))

	} else if !reflect.DeepEqual(n.Children[0], c1) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", n.Children[0], c1)
	}

	c2 := syntax.NewNode("child-2", "child-2")
	if n.AddChild(c2) == false {
		t.Errorf("add child operation failed")
	}
	if len(n.Children) != 2 {
		t.Errorf("length mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", 2, len(n.Children))

	} else if !reflect.DeepEqual(n.Children[1], c2) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", n.Children[1], c2)
	}
}
