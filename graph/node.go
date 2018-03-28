package graph

import (
	"bytes"
	"fmt"
)

// Block represents block type can be present in a graph.
type Block int

const (
	// ILLEGAL block.
	ILLEGAL Block = iota

	// NOBLOCK block
	NOBLOCK
	// LOOPandSKIP block.
	LOOPandSKIP
	// LOOPandNOSKIP block.
	LOOPandNOSKIP
	// NOLOOPandSKIP block.
	NOLOOPandSKIP
	// NOLOOPandNOSKIP block
	NOLOOPandNOSKIP
)

var nodeID int

// Node represents a node in the graph.
type Node struct {
	ID       int
	Name     string
	Label    string
	Children []interface{}
	IsRoot   bool
	IsSink   bool
	IsStart  bool
	IsEnd    bool
	IsLoop   bool
	IsJoint  bool
}

// NewNode creates a new graph node.
func NewNode(name string, label string) *Node {
	nodeID++
	node := &Node{
		ID:    nodeID,
		Name:  name,
		Label: label,
	}
	return node
}

// NewJoint create a new graph joint node.
// Joint node is any node that does not contain information but is used
// to build the graph.
func NewJoint(name string, label string) *Node {
	node := NewNode(name, label)
	node.IsJoint = true
	return node
}

// NewRoot creates a new graph root node.
// Root node is at the top of the graph, it starts the graph and only
// can be one Root node in the graph.
func NewRoot() *Node {
	node := NewJoint("ROOT", "ROOT")
	node.IsRoot = true
	return node
}

// NewSink creates a new graph sink node.
// Sink node is at the bottom of the graph, it terminates the graph only
// can be one sink node.
func NewSink() *Node {
	node := NewJoint("SINK", "SINK")
	node.IsSink = true
	return node
}

// NewStart creates a new graph start node.
// Start node is used for building loop graphs, and it identifies the
// start point of the loop.
func NewStart() *Node {
	node := NewJoint("START", "START")
	node.IsStart = true
	return node
}

// NewEnd creates a new graph end node.
// End node is used for building loop graphs, and it identifies the
// end point or exit of the loop.
func NewEnd() *Node {
	node := NewJoint("END", "END")
	node.IsEnd = true
	return node
}

// NewLoop creast a new graph loop node.
// Loop node is used for building loop graphs, and it identfies the
// loop part which will point to the start of the loop.
func NewLoop() *Node {
	node := NewJoint("LOOP", "LOOP")
	node.IsLoop = true
	return node
}

// AddChild adds a new child node.
func (n *Node) AddChild(child *Node) bool {
	n.Children = append(n.Children, child)
	return true
}

// Graph represents a full graph.
type Graph struct {
	Root  *Node
	Sink  *Node
	Hook  *Node
	Start *Node
	End   *Node
	Loop  *Node
}

// NewGraph creates a new graph.
func NewGraph() *Graph {
	g := &Graph{
		Root: NewRoot(),
		Sink: NewSink(),
	}
	g.Hook = g.Root
	return g
}

// AddNode adds a new node to the graph.
func (g *Graph) AddNode(n *Node) bool {
	g.Hook.AddChild(n)
	g.Hook = n
	return true
}

// StartBlockNoLoopAndSkip creates a graph block without a loop
// but it can be skipped.
func (g *Graph) StartBlockNoLoopAndSkip() bool {
	g.Start = NewStart()
	g.End = NewEnd()
	g.Loop = NewLoop()

	// This is required for loops that can be skipped.
	g.Start.AddChild(g.End)

	g.Hook.AddChild(g.Start)
	g.Loop.AddChild(g.End)
	g.Hook = g.Start
	return true
}

// StartBlockNoLoopAndNoSkip creates a graph block without a loop
// and it can not be skipped.
func (g *Graph) StartBlockNoLoopAndNoSkip() bool {
	g.Start = NewStart()
	g.End = NewEnd()
	g.Loop = NewLoop()

	g.Hook.AddChild(g.Start)
	g.Loop.AddChild(g.End)
	g.Hook = g.Start
	return true
}

// StartBlockLoopAndSkip creates a graph block with a loop
// and it can be skipped.
func (g *Graph) StartBlockLoopAndSkip() bool {
	g.Start = NewStart()
	g.End = NewEnd()
	g.Loop = NewLoop()

	// This is required for repeated loops.
	g.Loop.AddChild(g.Start)

	// This is required for loops that can be skipped.
	g.Start.AddChild(g.End)

	g.Hook.AddChild(g.Start)
	g.Loop.AddChild(g.End)
	g.Hook = g.Start
	return true
}

// StartBlockLoopAndNoSkip creates a graph block with a loop
// and it can not be skipped.
func (g *Graph) StartBlockLoopAndNoSkip() bool {
	g.Start = NewStart()
	g.End = NewEnd()
	g.Loop = NewLoop()

	// This is required for repeated loops.
	g.Loop.AddChild(g.Start)

	g.Hook.AddChild(g.Start)
	g.Loop.AddChild(g.End)
	g.Hook = g.Start
	return true
}

// EndLoop ends a graph loop.
func (g *Graph) EndLoop() bool {
	g.Hook = g.End
	g.Start = nil
	g.End = nil
	g.Loop = nil
	return true
}

// AddNodeToLoop adds a node to a graph loop.
func (g *Graph) AddNodeToLoop(n *Node) bool {
	g.Hook.AddChild(n)
	n.AddChild(g.Loop)
	return true
}

// Terminate terminates a graph.
// Graph is terminated when no more nodes can be added and the sink
// node has been linked to the last node in the graph.
func (g *Graph) Terminate() {
	g.Hook.AddChild(g.Sink)
	g.Hook = nil
}

// ToString returns the graph in a string format.
func (g *Graph) ToString() string {
	var buffer bytes.Buffer
	traverse := g.Root
	buffer.WriteString(fmt.Sprintf("%d %s %s %d\n",
		traverse.ID, traverse.Name, traverse.Label, len(traverse.Children)))
	for traverse != nil {
		for _, c := range traverse.Children {
			child := c.(*Node)
			buffer.WriteString(fmt.Sprintf("%d %s %s %d\n",
				child.ID, child.Name, child.Label, len(child.Children)))
		}
		if len(traverse.Children) > 0 {
			traverse = traverse.Children[0].(*Node)
		} else {
			traverse = nil
		}
	}
	return buffer.String()
}

// MapBlockToGraphFunc maps block type with method to be used.
var MapBlockToGraphFunc = map[Block]func(g *Graph) bool{
	NOLOOPandSKIP: func(g *Graph) bool {
		return g.StartBlockNoLoopAndSkip()
	},
	LOOPandSKIP: func(g *Graph) bool {
		return g.StartBlockLoopAndSkip()
	},
	LOOPandNOSKIP: func(g *Graph) bool {
		return g.StartBlockLoopAndNoSkip()
	},
	NOLOOPandNOSKIP: func(g *Graph) bool {
		return g.StartBlockNoLoopAndNoSkip()
	},
}
