package graph_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/graph"
)

// TestNode_NewNode ensures the node structure works properly
func TestNode_NewNode(t *testing.T) {
	var tests = []struct {
		n   *graph.Node
		exp *graph.Node
	}{
		{
			n: graph.NewNode("test", nil),
			exp: &graph.Node{
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
			n: graph.NewNodeJoint("joint"),
			exp: &graph.Node{
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
			n: graph.NewNodeRoot(),
			exp: &graph.Node{
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
			n: graph.NewNodeStart(4),
			exp: &graph.Node{
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
			n: graph.NewNodeEnd(5),
			exp: &graph.Node{
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
			n: graph.NewNodeLoop(6),
			exp: &graph.Node{
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
}

// TestNode_AddChildren ensures the node structure works properly
func TestNode_AddChildren(t *testing.T) {
	n := graph.NewNode("main", nil)
	c1 := graph.NewNode("child-1", nil)
	if n.AddChild(c1) == false {
		t.Errorf("add child operation failed")
	}
	if len(n.Children) != 1 {
		t.Errorf("length mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", 1, len(n.Children))

	} else if !reflect.DeepEqual(n.Children[0], c1) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", n.Children[0], c1)
	}

	c2 := graph.NewNode("child-2", nil)
	if n.AddChild(c2) == false {
		t.Errorf("add child operation failed")
	}
	if len(n.Children) != 2 {
		t.Errorf("length mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", 2, len(n.Children))

	} else if !reflect.DeepEqual(n.Children[1], c2) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", n.Children[1], c2)
	}
}

// TestNode_PrependChildren ensures the node structure works properly
func TestNode_PrependChildren(t *testing.T) {
	n := graph.NewNode("main", nil)
	c1 := graph.NewNode("child-1", nil)
	if n.AddChild(c1) == false {
		t.Errorf("add child operation failed")
	}
	c2 := graph.NewNode("child-2", nil)
	if n.PrependChild(c2) == false {
		t.Errorf("prepend child operation failed")
	}
	if len(n.Children) != 2 {
		t.Errorf("length mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", 2, len(n.Children))

	} else if !reflect.DeepEqual(n.Children[0], c2) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", n.Children[0], c2)
	} else if !reflect.DeepEqual(n.Children[1], c1) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", n.Children[1], c1)
	}
}

// TestNode_IsIn ensures the node structure works properly
func TestNode_IsIn(t *testing.T) {
	n := graph.NewNode("main", nil)
	c1 := graph.NewNode("child-1", nil)
	c2 := graph.NewNode("child-2", nil)
	if n.IsIn([]*graph.Node{c1, c2}) == true {
		t.Errorf("IsIn operation failed")
	} else if n.IsIn([]*graph.Node{n, c2}) == false {
		t.Errorf("IsIn operation failed")
	}
}

// TestNode_ToMermaid ensures mermaid conversion works properly.
func TestNode_ToMermaid(t *testing.T) {
	n := graph.NewNode("main", nil)
	c1 := graph.NewNode("child-1", nil)
	if n.AddChild(c1) == false {
		t.Errorf("add child operation failed")
	}
	c2 := graph.NewNode("child-2", nil)
	if n.AddChild(c2) == false {
		t.Errorf("add child operation failed")
	}
	output := n.ToMermaidChildren()
	exp := "main --> child-1\nmain --> child-2\n"
	if !reflect.DeepEqual(output, exp) {
		t.Errorf("node mermaid mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", exp, output)
	}
}

// TestNode_ToContent ensures content conversion works properly.
func TestNode_ToContent(t *testing.T) {
	n := graph.NewNode("main", "MAIN")
	output := n.ToContent()
	exp := "[string              ]	\"MAIN\"\n"
	if !reflect.DeepEqual(output, exp) {
		t.Errorf("node content mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", exp, output)
	}

	c1 := graph.NewNode("child-1", "CHILD ONE")
	c2 := graph.NewNode("child-2", "CHILD TWO")
	if n.AddChild(c1) == false {
		t.Errorf("add child operation failed")
	}
	if n.AddChild(c2) == false {
		t.Errorf("add child operation failed")
	}
	output = n.ToContentChildren()
	exp = "[string              ]	\"CHILD ONE\"\n[string              ]	\"CHILD TWO\"\n"
	if !reflect.DeepEqual(output, exp) {
		t.Errorf("node content mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", exp, output)
	}
}
