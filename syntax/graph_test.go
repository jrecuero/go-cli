package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestGraph_Graph ensures the graph structure works properly
func TestGraph_Graph(t *testing.T) {
	//var tests = []struct {
	//    g   *syntax.Graph
	//    exp *syntax.Graph
	//}{
	//    {
	//        g: syntax.NewGraph(),
	//        exp: &syntax.Graph{
	//            Root: &syntax.Node{
	//                ID:       1,
	//                Label:    "ROOT",
	//                Children: nil,
	//                IsRoot:   true,
	//                IsSink:   false,
	//                IsStart:  false,
	//                IsEnd:    false,
	//                IsLoop:   false,
	//                IsJoint:  true,
	//                InPath:   false,
	//                BlockID:  -1,
	//            },
	//            Sink: &syntax.Node{
	//                ID:       2,
	//                Label:    "SINK",
	//                Children: nil,
	//                IsRoot:   false,
	//                IsSink:   true,
	//                IsStart:  false,
	//                IsEnd:    false,
	//                IsLoop:   false,
	//                IsJoint:  true,
	//                InPath:   false,
	//                BlockID:  -1,
	//            },
	//            Hook: &syntax.Node{
	//                ID:       1,
	//                Label:    "ROOT",
	//                Children: nil,
	//                IsRoot:   true,
	//                IsSink:   false,
	//                IsStart:  false,
	//                IsEnd:    false,
	//                IsLoop:   false,
	//                IsJoint:  true,
	//                InPath:   false,
	//                BlockID:  -1,
	//            },
	//            Blocks:      nil,
	//            ActiveBlock: nil,
	//            Terminated:  false,
	//        },
	//    },
	//}

	//for i, tt := range tests {
	//    if !reflect.DeepEqual(tt.g, tt.exp) {
	//        t.Errorf("%d. graph mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.exp, tt.g)
	//    }
	//}
}

// TestGraph_AddNode ensures node is added properly to the graph
func TestGraph_AddNode(t *testing.T) {
	//g := syntax.NewGraph()
	//c1 := syntax.NewNode("child-1", "child-1")
	//if g.AddNode(c1) == false {
	//    t.Errorf("add node operation failed")
	//}
	//if len(g.Root.Children) != 1 {
	//    t.Errorf("length mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", 1, len(g.Root.Children))
	//}
	//if !reflect.DeepEqual(g.Root.Children[0], c1) {
	//    t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", g.Root.Children[0], c1)
	//}
	//if !reflect.DeepEqual(g.Hook, c1) {
	//    t.Errorf("node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", g.Hook, c1)
	//}
}

// TestGraph_BlockNoLoopAndSkip ensures the graph structure works properly
func TestGraph_BlockNoLoopAndSkip(t *testing.T) {
	g := syntax.NewGraph()
	g.NewBlockNoLoopAndSkip()
	//if g.Start == nil || g.End == nil || g.Loop == nil {
	//    t.Errorf("Block was not created properly Start=%p End=%p Loop=%p\n\n", g.Start, g.End, g.Loop)
	//}
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	//if len(g.Start.Children) != 1 {
	//    t.Errorf("Block was not created properly Start.Children.Len=%d is not 1", len(g.Start.Children))
	//}
	//if g.Start.Children[0] != g.End {
	//    t.Errorf("Block was not created properly Start.Children=%p not equal to End=%p\n\n", g.Start.Children[0], g.End)
	//}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
	//if len(g.Loop.Children) != 1 {
	//    t.Errorf("Block was not created properly Loop.Children.Len=%d is not 1", len(g.Loop.Children))
	//}
	//if g.Loop.Children[0] != g.End {
	//    t.Errorf("Block was not created properly Loop.Children=%p not equal to End=%p\n\n", g.Loop.Children[0], g.End)
	//}
}

// TestGraph_BlockNoLoopAndNoSkip ensures the graph structure works properly
func TestGraph_BlockNoLoopAndNoSkip(t *testing.T) {
	g := syntax.NewGraph()
	g.NewBlockNoLoopAndNoSkip()
	//if g.Start == nil || g.End == nil || g.Loop == nil {
	//    t.Errorf("Block was not created properly Start=%p End=%p Loop=%p\n\n", g.Start, g.End, g.Loop)
	//}
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
	//if len(g.Loop.Children) != 1 {
	//    t.Errorf("Block was not created properly Loop.Children.Len=%d is not 1", len(g.Loop.Children))
	//}
	//if g.Loop.Children[0] != g.End {
	//    t.Errorf("Block was not created properly Loop.Children=%p not equal to End=%p\n\n", g.Loop.Children[0], g.End)
	//}
}

// TestGraph_BlockLoopAndSkip ensures the graph structure works properly
func TestGraph_BlockLoopAndSkip(t *testing.T) {
	g := syntax.NewGraph()
	g.NewBlockLoopAndSkip()
	//if g.Start == nil || g.End == nil || g.Loop == nil {
	//    t.Errorf("Block was not created properly Start=%p End=%p Loop=%p\n\n", g.Start, g.End, g.Loop)
	//}
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	//if len(g.Start.Children) != 1 {
	//    t.Errorf("Block was not created properly Start.Children.Len=%d is not 1", len(g.Start.Children))
	//}
	//if g.Start.Children[0] != g.End {
	//    t.Errorf("Block was not created properly Start.Children=%p not equal to End=%p\n\n", g.Start.Children[0], g.End)
	//}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
	//if len(g.Loop.Children) != 2 {
	//    t.Errorf("Block was not created properly Loop.Children.Len=%d is not 2", len(g.Loop.Children))
	//}
	//if g.Loop.Children[0] != g.Start {
	//    t.Errorf("Block was not created properly Loop.Children[0]=%p not equal to Start=%p\n\n", g.Loop.Children[0], g.Start)
	//}
	//if g.Loop.Children[1] != g.End {
	//    t.Errorf("Block was not created properly Loop.Children[1]=%p not equal to End=%p\n\n", g.Loop.Children[1], g.End)
	//}
}

// TestGraph_BlockLoopAndSkip ensures the graph structure works properly
func TestGraph_BlockLoopAndNoSkip(t *testing.T) {
	g := syntax.NewGraph()
	g.NewBlockLoopAndNoSkip()
	//if g.Start == nil || g.End == nil || g.Loop == nil {
	//    t.Errorf("Block was not created properly Start=%p End=%p Loop=%p\n\n", g.Start, g.End, g.Loop)
	//}
	if g.Hook != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	}
	if len(g.Root.Children) != 1 {
		t.Errorf("Block was not created properly Root.Children.Len=%d is not 1", len(g.Root.Children))
	}
	if g.Root.Children[0] != g.ActiveBlock.Start {
		t.Errorf("Block was not created properly Root.Children=%p not equal to Start=%p\n\n", g.Root.Children[0], g.ActiveBlock.Start)
	}
	//if len(g.Loop.Children) != 2 {
	//    t.Errorf("Block was not created properly Loop.Children.Len=%d is not 2", len(g.Loop.Children))
	//}
	//if g.Loop.Children[0] != g.Start {
	//    t.Errorf("Block was not created properly Loop.Children[0]=%p not equal to Start=%p\n\n", g.Loop.Children[0], g.Start)
	//}
	//if g.Loop.Children[1] != g.End {
	//    t.Errorf("Block was not created properly Loop.Children[1]=%p not equal to End=%p\n\n", g.Loop.Children[1], g.End)
	//}
}

// TestGraph_EndLoop ensures the graph structure works properly
func TestGraph_EndLoop(t *testing.T) {
	g := syntax.NewGraph()
	g.NewBlockNoLoopAndSkip()
	if g.EndLoop() == false {
		t.Errorf("end loop operation failed")
	}
	//if g.Start != nil || g.End != nil || g.Loop != nil || g.Hook == nil {
	//    t.Errorf("Block was not ended properly Start=%p End=%p Loop=%p Hook=%p\n\n", g.Start, g.End, g.Loop, g.Hook)
	//}
}

// TestGraph_AddNodeToLoop ensures the graph structure works properly
func TestGraph_AddNodeToLoop(t *testing.T) {
	//g := syntax.NewGraph()
	//g.NewBlockNoLoopAndSkip()
	//c1 := syntax.NewNode("child-1", "child-1")
	//if g.AddNodeToBlock(c1) == false {
	//    t.Errorf("add node to loop operation failed")
	//}
	//if g.Hook != g.ActiveBlock.Start {
	//    t.Errorf("Node was not added to loop properly Start=%p not equal to Hook=%p\n\n", g.ActiveBlock.Start, g.Hook)
	//}
	//if len(g.Hook.Children) != 2 {
	//    t.Errorf("Node was not added to loop properly Start.Children.Len=%d is not 2", len(g.Hook.Children))
	//}
	//if g.Hook.Children[0] != g.ActiveBlock.End {
	//    t.Errorf("Node was not added to loop properly Hook.Children[0]=%p not equal to End=%p\n\n", g.Hook.Children[0], g.ActiveBlock.End)
	//}
	//if g.Hook.Children[1] != c1 {
	//    t.Errorf("Node was not added to  properly Hook.Children[1]=%p not equal to new-node=%p\n\n", g.Hook.Children[1], c1)
	//}
	//if len(c1.Children) != 1 {
	//    t.Errorf("Node was not added  properly new-node.Len=%d is not 2", len(c1.Children))
	//}
	//if c1.Children[0] != g.ActiveBlock.Loop {
	//    t.Errorf("Block was not created properly new-node.Children=%p not equal to Loop=%p\n\n", c1.Children[0], g.ActiveBlock.Loop)
	//}
}

// TestGraph_Terminate ensures the graph structure works properly
func TestGraph_Terminate(t *testing.T) {
	g := syntax.NewGraph()
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
	//g := syntax.NewGraph()
	//g.Terminate()
	//var tests = []struct {
	//    got string
	//    exp string
	//}{
	//    {
	//        got: g.ToString(),
	//        exp: "39 ROOT ROOT 1\n40 SINK SINK 0\n",
	//    },
	//}
	//for i, tt := range tests {
	//    if !reflect.DeepEqual(tt.got, tt.exp) {
	//        t.Errorf("%d. node mistmatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.exp, tt.got)
	//    }
	//}
}
