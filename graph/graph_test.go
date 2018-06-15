package graph_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/graph"
)

// TestGraph_NewGraph ensures the graph structure works properly
func TestGraph_NewGraph(t *testing.T) {
	var tests = []struct {
		g   *graph.Graph
		exp *graph.Graph
	}{
		{
			g: graph.NewGraph(nil),
			exp: &graph.Graph{
				Root: &graph.Node{
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
				Sink: &graph.Node{
					Label:    "SINK",
					Children: nil,
					IsRoot:   false,
					IsSink:   true,
					IsStart:  false,
					IsEnd:    false,
					IsLoop:   false,
					IsJoint:  true,
					InPath:   false,
					BlockID:  -1,
				},
				Hook: &graph.Node{
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
				Blocks:      nil,
				ActiveBlock: nil,
				Terminated:  false,
				Setup:       &graph.SetupGraph{},
			},
		},
	}

	for i, tt := range tests {
		if !reflect.DeepEqual(tt.g, tt.exp) {
			t.Errorf("%d. graph mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.exp, tt.g)
		}
	}
}

// TestGraph_AddNode ensures node is added properly to the graph
func TestGraph_AddNode(t *testing.T) {
	g := graph.NewGraph(nil)
	c1 := graph.NewNode("child-1", "child-1")
	if g.AddNode(c1) == false {
		t.Errorf("add node operation failed")
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("length mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", 1, len(g.Root.Children))
	}
	if !reflect.DeepEqual(g.Root.Children[0], c1) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", g.Root.Children[0], c1)
	}
	if !reflect.DeepEqual(g.Hook, c1) {
		t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", g.Hook, c1)
	}
}

// TestGraph_BlockNoLoopAndSkip ensures the graph structure works properly
func TestGraph_BlockNoLoopAndSkip(t *testing.T) {
	g := graph.NewGraph(nil)
	g.NewBlockNoLoopAndSkip()
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
}

// TestGraph_BlockNoLoopAndNoSkip ensures the graph structure works properly
func TestGraph_BlockNoLoopAndNoSkip(t *testing.T) {
	g := graph.NewGraph(nil)
	g.NewBlockNoLoopAndNoSkip()
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
}

// TestGraph_BlockLoopAndSkip ensures the graph structure works properly
func TestGraph_BlockLoopAndSkip(t *testing.T) {
	g := graph.NewGraph(nil)
	g.NewBlockLoopAndSkip()
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
}

// TestGraph_BlockLoopAndSkip ensures the graph structure works properly
func TestGraph_BlockLoopAndNoSkip(t *testing.T) {
	g := graph.NewGraph(nil)
	g.NewBlockLoopAndNoSkip()
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
}

// TestGraph_EndLoop ensures the graph structure works properly
func TestGraph_EndLoop(t *testing.T) {
	g := graph.NewGraph(nil)
	g.NewBlockNoLoopAndSkip()
	if g.EndLoop() == false {
		t.Errorf("end loop operation failed")
	}
}

// TestGraph_AddNodeToLoop ensures the graph structure works properly
func TestGraph_AddNodeToLoop(t *testing.T) {
	g := graph.NewGraph(nil)
	g.NewBlockNoLoopAndSkip()
	c1 := graph.NewNode("child-1", "child-1")
	if g.AddNodeToBlock(c1) == false {
		t.Errorf("add node to loop operation failed")
	}
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Node was not added to loop properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	if len(g.Hook.Children) != 2 {
		t.Errorf("Node was not added to loop properly Start.Children.Len=%d is not 2", len(g.Hook.Children))
	}
	if g.Hook.Children[0] != g.ActiveBlock.End {
		t.Errorf("Node was not added to loop properly Hook.Children[0]=%p not equal to End=%p\n\n", g.Hook.Children[0], g.ActiveBlock.End)
	}
	if g.Hook.Children[1] != c1 {
		t.Errorf("Node was not added to  properly Hook.Children[1]=%p not equal to new-node=%p\n\n", g.Hook.Children[1], c1)
	}
	if len(c1.Children) != 1 {
		t.Errorf("Node was not added  properly new-node.Len=%d is not 2", len(c1.Children))
	}
	if c1.Children[0] != g.ActiveBlock.Loop {
		t.Errorf("Block was not created properly new-node.Children=%p not equal to Loop=%p\n\n", c1.Children[0], g.ActiveBlock.Loop)
	}
}

// TestGraph_Terminate ensures the graph structure works properly
func TestGraph_Terminate(t *testing.T) {
	g := graph.NewGraph(nil)
	g.Terminate()
	if g.Hook != nil {
		t.Errorf("graph was not terminated properly Hook=%p", g.Hook)
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("graph was not terminated properly len(g.Root.Children)=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.Sink {
		t.Errorf("graph was not terminated properly Root.Children=%p not equal to Sink=%p\n\n", g.Root.Children[0], g.Sink)
	}

}

// TestGraph_ToString ensures the graph is serialize properly
func TestGraph_ToString(t *testing.T) {
	g := graph.NewGraph(nil)
	g.Terminate()
	var tests = []struct {
		got string
		exp string
	}{
		{
			got: g.ToString(),
			exp: "ROOT 1\nSINK 0\n",
		},
	}
	for i, tt := range tests {
		if !reflect.DeepEqual(tt.got, tt.exp) {
			t.Errorf("%d. node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.exp, tt.got)
		}
	}
}
