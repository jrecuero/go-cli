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
				Label:         "test",
				Children:      nil,
				IsRoot:        false,
				IsSink:        false,
				IsNext:        false,
				IsStart:       false,
				IsEnd:         false,
				IsLoop:        false,
				IsJoint:       false,
				InPath:        false,
				AllowChildren: true,
				BlockID:       -1,
			},
		},
		{
			n: graph.NewNodeJoint("joint", nil),
			exp: &graph.Node{
				Label:         "joint",
				Children:      nil,
				IsRoot:        false,
				IsSink:        false,
				IsNext:        false,
				IsStart:       false,
				IsEnd:         false,
				IsLoop:        false,
				IsJoint:       true,
				InPath:        false,
				AllowChildren: true,
				BlockID:       -1,
			},
		},
		{
			n: graph.NewNodeRoot(nil),
			exp: &graph.Node{
				Label:         "ROOT",
				Children:      nil,
				IsRoot:        true,
				IsSink:        false,
				IsNext:        false,
				IsStart:       false,
				IsEnd:         false,
				IsLoop:        false,
				IsJoint:       true,
				InPath:        false,
				AllowChildren: true,
				BlockID:       -1,
			},
		},
		{
			n: graph.NewNodeStart(4, nil),
			exp: &graph.Node{
				Label:         "START",
				Children:      nil,
				IsRoot:        false,
				IsSink:        false,
				IsNext:        false,
				IsStart:       true,
				IsEnd:         false,
				IsLoop:        false,
				IsJoint:       true,
				InPath:        false,
				AllowChildren: true,
				BlockID:       4,
			},
		},
		{
			n: graph.NewNodeEnd(5, nil),
			exp: &graph.Node{
				Label:         "END",
				Children:      nil,
				IsRoot:        false,
				IsSink:        false,
				IsNext:        false,
				IsStart:       false,
				IsEnd:         true,
				IsLoop:        false,
				IsJoint:       true,
				InPath:        false,
				AllowChildren: true,
				BlockID:       5,
			},
		},
		{
			n: graph.NewNodeLoop(6, nil),
			exp: &graph.Node{
				Label:         "LOOP",
				Children:      nil,
				IsRoot:        false,
				IsSink:        false,
				IsNext:        false,
				IsStart:       false,
				IsEnd:         false,
				IsLoop:        true,
				IsJoint:       true,
				InPath:        false,
				AllowChildren: true,
				BlockID:       6,
			},
		},
	}

	for i, tt := range tests {
		if tt.n.Label != tt.exp.Label {
			t.Errorf("%d. node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.exp.Label, tt.n.Label)
		}
		//if !reflect.DeepEqual(tt.n, tt.exp) {
		//    t.Errorf("%d. node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.exp, tt.n)
		//}
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
	exp := "main-69[main] --> child-1-70[child-1]\nmain-69[main] --> child-2-71[child-2]\n"
	if !reflect.DeepEqual(output, exp) {
		t.Errorf("node mermaid mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", exp, output)
	}
}

// TestNode_ToContent ensures content conversion works properly.
func TestNode_ToContent(t *testing.T) {
	n := graph.NewNode("main", &contenido{"MAIN"})
	output := n.ToContent()
	exp := "[*graph_test.contenido]\t\"MAIN\"\n"
	if !reflect.DeepEqual(output, exp) {
		t.Errorf("node content mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", exp, output)
	}

	c1 := graph.NewNode("child-1", &contenido{"CHILD ONE"})
	c2 := graph.NewNode("child-2", &contenido{"CHILD TWO"})
	if n.AddChild(c1) == false {
		t.Errorf("add child operation failed")
	}
	if n.AddChild(c2) == false {
		t.Errorf("add child operation failed")
	}
	output = n.ToContentChildren()
	//exp = "[string              ]	\"CHILD ONE\"\n[string              ]	\"CHILD TWO\"\n"
	exp = "[*graph_test.contenido]\t\"CHILD ONE\"\n[*graph_test.contenido]\t\"CHILD TWO\"\n"
	if !reflect.DeepEqual(output, exp) {
		t.Errorf("node content mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", exp, output)
	}
}

// TestNode_Match tests Node.Match behavior..
func TestNode_Match(t *testing.T) {
	n := graph.NewNode("main", &contenido{"MAIN"})
	if index, ok := n.Match(nil, nil, 0); index != 0 || ok != true {
		t.Errorf("match operation failed: index: %d ok: %v", index, ok)
	}
}

// TestNode_Help tests Node.Help behavior.
func TestNode_Help(t *testing.T) {
	n := graph.NewNode("main", &contenido{"MAIN"})
	if itf, ok := n.Help(nil, nil, 0); itf.(*contenido).label != "MAIN" || ok != true {
		t.Errorf("help operation failed")
	}
}

// TestNode_Query tests Node.Query behavior.
func TestNode_Query(t *testing.T) {
	n := graph.NewNode("main", &contenido{"MAIN"})
	if itf, ok := n.Query(nil, nil, 0); itf != nil || ok != true {
		t.Errorf("query operation failed")
	}
}

// TestNode_Complete tests Node.Complete behavior.
func TestNode_Complete(t *testing.T) {
	n := graph.NewNode("main", &contenido{"MAIN"})
	if itf, ok := n.Complete(nil, nil, 0); itf.(*contenido).label != "MAIN" || ok != true {
		t.Errorf("complete operation failed")
	}
}

// TestNode_Validate tests Node.Validate behavior.
func TestNode_Validate(t *testing.T) {
	n := graph.NewNode("main", &contenido{"MAIN"})
	if ok := n.Validate(nil, nil, 0); ok != true {
		t.Errorf("validate operation failed")
	}
}
