package graph

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BlockType represents block type can be present in a graph.
type BlockType int

const (
	// ILLEGAL block.
	ILLEGAL BlockType = iota

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
	InPath   bool
	BlockID  int
}

// NewNode creates a new graph node.
func NewNode(name string, label string) *Node {
	nodeID++
	node := &Node{
		ID:      nodeID,
		Name:    name,
		Label:   label,
		BlockID: -1,
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
func NewStart(id int) *Node {
	node := NewJoint("START", "START")
	node.IsStart = true
	node.BlockID = id
	return node
}

// NewEnd creates a new graph end node.
// End node is used for building loop graphs, and it identifies the
// end point or exit of the loop.
func NewEnd(id int) *Node {
	node := NewJoint("END", "END")
	node.IsEnd = true
	node.BlockID = id
	return node
}

// NewLoop creast a new graph loop node.
// Loop node is used for building loop graphs, and it identfies the
// loop part which will point to the start of the loop.
func NewLoop(id int) *Node {
	node := NewJoint("LOOP", "LOOP")
	node.IsLoop = true
	node.BlockID = id
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

func (n *Node) mermaidLabel() string {
	var buffer bytes.Buffer
	if n.BlockID == -1 {
		if n.IsJoint == true {
			buffer.WriteString(fmt.Sprintf("%s((%s))", n.Label, n.Label))
		} else {
			buffer.WriteString(n.Label)
		}
	} else {
		buffer.WriteString(fmt.Sprintf("%s-%d((%s))", n.Label, n.BlockID, n.Label))
	}
	return buffer.String()
}

// ToMermaid returns the node in Mermaid graph format.
func (n *Node) ToMermaid() string {
	var buffer bytes.Buffer
	for _, child := range n.Children {
		//fmt.Printf("mermaid %s to %s\n", n.Label, child.Label)
		buffer.WriteString(fmt.Sprintf("%s --> %s\n", n.mermaidLabel(), child.mermaidLabel()))
	}
	return buffer.String()
}

// Block represents a graph block.
type Block struct {
	ID         int
	Start      *Node
	End        *Node
	Loop       *Node
	IsSkip     bool
	IsLoop     bool
	Terminated bool
}

// NewBlock creates a new graph block.
func NewBlock(id int) *Block {
	b := &Block{
		ID:         id,
		Start:      NewStart(id),
		End:        NewEnd(id),
		Loop:       NewLoop(id),
		IsLoop:     false,
		IsSkip:     false,
		Terminated: false,
	}
	return b
}

// CreateBlockNoLoopAndSkip creates a graph block without a loop
// but it can be skipped.
func (b *Block) CreateBlockNoLoopAndSkip() bool {
	b.IsLoop = false
	b.IsSkip = true
	// Next statement is required for loops that can be skipped.
	b.Start.AddChild(b.End)
	b.Loop.AddChild(b.End)
	return true
}

// CreateBlockNoLoopAndNoSkip creates a graph block without a loop
// and it can not be skipped.
func (b *Block) CreateBlockNoLoopAndNoSkip() bool {
	b.IsLoop = false
	b.IsSkip = false
	b.Loop.AddChild(b.End)
	return true
}

// CreateBlockLoopAndSkip creates a graph block with a loop
// and it can be skipped.
func (b *Block) CreateBlockLoopAndSkip() bool {
	b.IsLoop = true
	b.IsSkip = true
	// Next statement required for loops that can be skipped.
	b.Start.AddChild(b.End)
	// Next statement required for repeated loops.
	b.Loop.AddChild(b.Start)
	b.Loop.AddChild(b.End)
	return true
}

// CreateBlockLoopAndNoSkip creates a graph block with a loop
// and it can not be skipped.
func (b *Block) CreateBlockLoopAndNoSkip() bool {
	b.IsLoop = true
	b.IsSkip = false
	// Next statement required for repeated loops.
	b.Loop.AddChild(b.Start)
	b.Loop.AddChild(b.End)
	return true
}

// Terminate ends a graph loop.
func (b *Block) Terminate() bool {
	// Blocks with skip option are adding a child to END first, but it is
	// better to place that child at the end of the array.
	if b.IsSkip == true {
		childrenLen := len(b.Start.Children)
		b.Start.Children = append(b.Start.Children[1:childrenLen], b.Start.Children[0])
	}
	b.Terminated = true
	return true
}

// Graph represents a full graph.
type Graph struct {
	Root        *Node
	Sink        *Node
	Hook        *Node
	Blocks      []*Block
	ActiveBlock *Block
	Terminated  bool
	visited     []*Node
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

// newBlock creates a new generic block in the graph.
func (g *Graph) newBlock() {
	index := len(g.Blocks)
	g.ActiveBlock = NewBlock(index)
	g.Blocks = append(g.Blocks, g.ActiveBlock)
}

// setupBlock setups the Hook node to the new graph block.
func (g *Graph) setupBlock() {
	g.Hook.AddChild(g.ActiveBlock.Start)
	g.Hook = g.ActiveBlock.Start
}

// NewBlockNoLoopAndSkip creates a graph block without a loop
// but it can be skipped.
func (g *Graph) NewBlockNoLoopAndSkip() bool {
	g.newBlock()
	g.ActiveBlock.CreateBlockNoLoopAndSkip()
	g.setupBlock()
	return true
}

// NewBlockNoLoopAndNoSkip creates a graph block without a loop
// and it can not be skipped.
func (g *Graph) NewBlockNoLoopAndNoSkip() bool {
	g.newBlock()
	g.ActiveBlock.CreateBlockNoLoopAndNoSkip()
	g.setupBlock()
	return true
}

// NewBlockLoopAndSkip creates a graph block with a loop
// and it can be skipped.
func (g *Graph) NewBlockLoopAndSkip() bool {
	g.newBlock()
	g.ActiveBlock.CreateBlockLoopAndSkip()
	g.setupBlock()
	return true
}

// NewBlockLoopAndNoSkip creates a graph block with a loop
// and it can not be skipped.
func (g *Graph) NewBlockLoopAndNoSkip() bool {
	g.newBlock()
	g.ActiveBlock.CreateBlockLoopAndNoSkip()
	g.setupBlock()
	return true
}

// EndLoop ends a graph loop.
func (g *Graph) EndLoop() bool {
	g.Hook = g.ActiveBlock.End
	g.ActiveBlock.Terminate()
	g.ActiveBlock = nil
	return true
}

// AddNodeToBlock adds a node to a graph loop.
func (g *Graph) AddNodeToBlock(n *Node) bool {
	g.Hook.AddChild(n)
	n.AddChild(g.ActiveBlock.Loop)
	return true
}

// AddPathToBlock adds a node to a node path in a graph block.
func (g *Graph) AddPathToBlock(n *Node) bool {
	n.InPath = true
	g.Hook.AddChild(n)
	g.Hook = n
	return true
}

// TerminatePathToBlock terminated a node path in a graph block.
func (g *Graph) TerminatePathToBlock() bool {
	g.Hook.AddChild(g.ActiveBlock.Loop)
	return true
}

// Terminate terminates a graph.
// Graph is terminated when no more nodes can be added and the sink
// node has been linked to the last node in the graph.
func (g *Graph) Terminate() {
	g.Hook.AddChild(g.Sink)
	g.Hook = nil
	g.Terminated = true
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
func (g *Graph) childrenToString(node *Node, visited []*Node) string {
	var buffer bytes.Buffer
	for _, child := range node.Children {
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

func (g *Graph) childrenToMermaid(node *Node) string {
	var buffer bytes.Buffer
	for _, child := range node.Children {
		if child.IsIn(g.visited) == true {
			continue
		}
		//fmt.Printf("graph for %s\n", child.Label)
		buffer.WriteString(child.ToMermaid())
		g.visited = append(g.visited, child)
		//fmt.Printf("add %s to g.visited\n", child.Label)
	}
	var children []*Node
	if node.IsLoop == false || len(node.Children) == 0 {
		children = node.Children
	} else {
		children = node.Children[len(node.Children)-1:]
	}
	for _, child := range children {
		//fmt.Printf("children to process %s\n", child.Label)
		//if child.IsIn(g.visited) == true {
		//    continue
		//}
		//g.visited = append(g.visited, child)
		buffer.WriteString(g.childrenToMermaid(child))
	}
	return buffer.String()
}

// ToMermaid returns the graph in Mermaid graph format.
func (g *Graph) ToMermaid() string {
	var buffer bytes.Buffer
	g.visited = []*Node{}
	buffer.WriteString("graph TD\n")
	buffer.WriteString(g.Root.ToMermaid())
	g.visited = append(g.visited, g.Root)
	buffer.WriteString(g.childrenToMermaid(g.Root))
	return buffer.String()
}

// MapBlockToGraphFunc maps block type with method to be used.
var MapBlockToGraphFunc = map[BlockType]func(g *Graph) bool{
	NOLOOPandSKIP: func(g *Graph) bool {
		return g.NewBlockNoLoopAndSkip()
	},
	LOOPandSKIP: func(g *Graph) bool {
		return g.NewBlockLoopAndSkip()
	},
	LOOPandNOSKIP: func(g *Graph) bool {
		return g.NewBlockLoopAndNoSkip()
	},
	NOLOOPandNOSKIP: func(g *Graph) bool {
		return g.NewBlockNoLoopAndNoSkip()
	},
}
