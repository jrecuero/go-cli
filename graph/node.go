package graph

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	Children []*Node
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

// PrependChild adds a new child node first in the array.
func (n *Node) PrependChild(child *Node) bool {
	n.Children = append([]*Node{child}, n.Children...)
	return true
}

// IsIn checks if the node is in the given array.
func (n *Node) IsIn(array []*Node) bool {
	for _, node := range array {
		if node == n {
			return true
		}
	}
	return false
}

// Graph represents a full graph.
type Graph struct {
	Root   *Node
	Sink   *Node
	Hook   *Node
	Start  *Node
	End    *Node
	Loop   *Node
	IsSkip bool
	IsLoop bool
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
	g.IsLoop = false
	g.IsSkip = true

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
	g.IsLoop = false
	g.IsSkip = false

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
	g.IsLoop = true
	g.IsSkip = true

	// This is required for loops that can be skipped.
	g.Start.AddChild(g.End)

	// This is required for repeated loops.
	g.Loop.AddChild(g.Start)

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
	g.IsLoop = true
	g.IsSkip = false

	// This is required for repeated loops.
	g.Loop.AddChild(g.Start)

	g.Hook.AddChild(g.Start)
	g.Loop.AddChild(g.End)
	g.Hook = g.Start
	return true
}

// EndLoop ends a graph loop.
func (g *Graph) EndLoop() bool {
	// Blocks with skip option are adding a child to END first, but it is
	// better to place that child at the end of the array.
	if g.IsSkip == true {
		childrenLen := len(g.Start.Children)
		g.Start.Children = append(g.Start.Children[1:childrenLen], g.Start.Children[0])
	}
	g.Hook = g.End
	g.Start = nil
	g.End = nil
	g.Loop = nil
	g.IsLoop = false
	g.IsSkip = false
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

// Explore implements a mechanism to interactibily explore the graph.
func (g *Graph) Explore() {
	reader := bufio.NewReader(os.Stdin)
	traverse := g.Root
	parents := []*Node{}
	var index int
	for {
		fmt.Printf(fmt.Sprintf("\n\nID: %d Name: %s Label: %s\n", traverse.ID, traverse.Name, traverse.Label))
		fmt.Printf(fmt.Sprintf("Nbr of children: %d\n", len(traverse.Children)))
		if len(traverse.Children) > 0 {
			for i, child := range traverse.Children {
				fmt.Printf("\t> %d %s\n", i, child.Name)
			}
			fmt.Printf("\n[0-%d] Select children", len(traverse.Children)-1)
		}
		fmt.Printf("\n[-] Select Parent\n[x] Exit\nSelect: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "x" {
			break
		} else if text == "-" {
			traverse = parents[len(parents)-1]
			parents = parents[:len(parents)-1]
			fmt.Printf("Parent selected %s\n", traverse.Name)
		} else {
			index, _ = strconv.Atoi(text)
			parents = append(parents, traverse)
			traverse = traverse.Children[index]
			fmt.Printf("Children selected %d - %s\n", index, traverse.Name)
		}
	}
}

// childrenToString returns all children information for a given node.
// It traverse all children in a recursive way.
func (g *Graph) childrenToString(root *Node, visited []*Node) string {
	var buffer bytes.Buffer
	for _, child := range root.Children {
		if child.IsIn(visited) == true {
			continue
		}
		//fmt.Printf("visiting node %s %+v\n", child.Name, visited)
		visited = append(visited, child)
		buffer.WriteString(fmt.Sprintf("%d %s %s %d\n", child.ID, child.Name, child.Label, len(child.Children)))
		//fmt.Printf("visited %+v\n", visited)
		buffer.WriteString(g.childrenToString(child, visited))
	}
	return buffer.String()
}

// ToString returns the graph in a string format.
func (g *Graph) ToString() string {
	var buffer bytes.Buffer
	visited := []*Node{}
	traverse := g.Root
	buffer.WriteString(fmt.Sprintf("%d %s %s %d\n", traverse.ID, traverse.Name, traverse.Label, len(traverse.Children)))
	visited = append(visited, traverse)
	buffer.WriteString(g.childrenToString(traverse, visited))
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
