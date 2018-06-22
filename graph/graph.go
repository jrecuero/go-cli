package graph

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const _mermaid = "MERMAID"
const _content = "CONTENT"

// SetupGraph represents all customizable fields for the Graph. Content
// information should be enclosed in these instances.
type SetupGraph struct {
	RootContent  interface{}
	SinkContent  interface{}
	NextContent  interface{}
	JointContent interface{}
	StartContent interface{}
	EndContent   interface{}
	LoopContent  interface{}
}

// Graph represents a full graph.
type Graph struct {
	Root        *Node
	Sink        *Node
	Next        *Node
	Hook        *Node
	Blocks      []*Block
	ActiveBlock *Block
	Terminated  bool
	visited     []*Node
	Setup       *SetupGraph
}

// NewGraph creates a new graph.
func NewGraph(setupG *SetupGraph) *Graph {
	if setupG == nil {
		setupG = &SetupGraph{}
	}
	g := &Graph{
		Root:  NewNodeRoot(setupG.RootContent),
		Sink:  NewNodeSink(setupG.SinkContent),
		Next:  NewNodeNext(setupG.NextContent),
		Setup: setupG,
	}
	g.Hook = g.Root
	return g
}

// AddNode adds a new node to the graph. It moves the Hook to the new Node
// being added.
func (g *Graph) AddNode(n *Node) bool {
	g.Hook.AddChild(n)
	g.Hook = n
	return true
}

// newBlock creates a new generic block in the graph.
func (g *Graph) newBlock() {
	index := len(g.Blocks)
	g.ActiveBlock = NewBlock(index, g.Setup.StartContent, g.Setup.EndContent, g.Setup.LoopContent)
	g.Blocks = append(g.Blocks, g.ActiveBlock)
}

// setupBlock setups the Hook node to the new graph block.
func (g *Graph) setupBlock() {
	g.Hook.AddChild(g.ActiveBlock.Start)
	g.setupHookToBlockStart()
}

// setupHookToBlockStart setups the graph block to point to the block start.
func (g *Graph) setupHookToBlockStart() {
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

// AddNodeToBlock adds a node to a graph loop. It adds the node to the Hook,
// but it does update the Hook value. Loop is added as a Node child.
func (g *Graph) AddNodeToBlock(n *Node) bool {
	g.Hook.AddChild(n)
	n.AddChild(g.ActiveBlock.Loop)
	return true
}

// AddPathToBlock adds a node to a node path in a graph block. It moves the
// Hookd to the new Node being added.
func (g *Graph) AddPathToBlock(n *Node) bool {
	n.InPath = true
	g.Hook.AddChild(n)
	g.Hook = n
	return true
}

// CreatePathToBlock creates a path in a block with the list of given nodes.
func (g *Graph) CreatePathToBlock(ln []*Node) bool {
	for _, n := range ln {
		g.AddPathToBlock(n)
	}
	return g.TerminatePathToBlock()
}

// AddIdentAndAnyToBlock adds an ident and an any node to the graph block. This
// can be used to add nodes that contains keyword-value pairs.
func (g *Graph) AddIdentAndAnyToBlock(ident *Node, any *Node) bool {
	return g.CreatePathToBlock([]*Node{ident, any})
}

// TerminatePathToBlock terminated a node path in a graph block.
func (g *Graph) TerminatePathToBlock() bool {
	g.Hook.AddChild(g.ActiveBlock.Loop)
	g.setupHookToBlockStart()
	return true
}

// Terminate terminates a graph.
// Graph is terminated when no more nodes can be added and the sink
// node has been linked to the last node in the graph.
func (g *Graph) Terminate() {
	if g.Sink != nil {
		g.Hook.AddChild(g.Sink)
	}
	if g.Next != nil {
		g.Hook.AddChild(g.Next)
	}
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
		fmt.Printf(fmt.Sprintf("\n\nNode Information: %s\n", traverse.ToContent()))
		fmt.Printf(fmt.Sprintf("Nbr of children: %d\n", len(traverse.Children)))
		if len(traverse.Children) > 0 {
			for i, child := range traverse.Children {
				//fmt.Printf("\t> %d %s\n", i, child.Label)
				fmt.Printf("\t> %d %s\n", i, child.ToContent())
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
			fmt.Printf("Parent selected %s\n", traverse.Label)
		} else {
			index, _ = strconv.Atoi(text)
			parents = append(parents, traverse)
			traverse = traverse.Children[index]
			fmt.Printf("Children selected %d - %s\n", index, traverse.Label)
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
		visited = append(visited, child)
		buffer.WriteString(fmt.Sprintf("%s %d\n", child.Label, len(child.Children)))
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
	buffer.WriteString(fmt.Sprintf("%s %d\n", traverse.Label, len(traverse.Children)))
	visited = append(visited, traverse)
	buffer.WriteString(g.childrenToString(traverse, visited))
	return buffer.String()
}

// childrenToCustom allows to build a mermaid graph or a content graph.
func (g *Graph) childrenToCustom(node *Node, custom string) string {
	var buffer bytes.Buffer
	for _, child := range node.Children {
		if child.IsIn(g.visited) == true {
			continue
		}
		switch custom {
		case _mermaid:
			buffer.WriteString(child.ToMermaidChildren())
		case _content:
			buffer.WriteString(child.ToContent())
		}
		g.visited = append(g.visited, child)
	}
	var children []*Node
	if node.IsLoop == false || len(node.Children) == 0 {
		children = node.Children
	} else {
		children = node.Children[len(node.Children)-1:]
	}
	for _, child := range children {
		switch custom {
		case _mermaid:
			buffer.WriteString(g.childrenToMermaid(child))
		case _content:
			buffer.WriteString(g.childrenToContent(child))
		}
	}
	return buffer.String()
}

// childrenToMermaid calls the custom build for mermaid graph.
func (g *Graph) childrenToMermaid(node *Node) string {
	return g.childrenToCustom(node, _mermaid)
}

// childrenToContent calls the custom build for content graph.
func (g *Graph) childrenToContent(node *Node) string {
	return g.childrenToCustom(node, _content)
}

// ToMermaid returns the graph in Mermaid graph format.
func (g *Graph) ToMermaid() string {
	var buffer bytes.Buffer
	g.visited = []*Node{}
	buffer.WriteString("graph TD\n")
	buffer.WriteString(g.Root.ToMermaidChildren())
	g.visited = append(g.visited, g.Root)
	buffer.WriteString(g.childrenToMermaid(g.Root))
	return buffer.String()
}

// ToContent returns node graph content information.
func (g *Graph) ToContent() string {
	var buffer bytes.Buffer
	g.visited = []*Node{}
	buffer.WriteString("Content\n")
	buffer.WriteString(g.Root.ToContent())
	g.visited = append(g.visited, g.Root)
	buffer.WriteString(g.childrenToContent(g.Root))
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
