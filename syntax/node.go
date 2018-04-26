package syntax

import (
	"bytes"
	"fmt"
)

var nodeID int

// Node represents a node in the graph.
type Node struct {
	ID       int
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
	Content  IContent
}

// Completer returns the content completer
func (n *Node) Completer() ICompleter {
	if n.Content == nil {
		return nil
	}
	return n.Content.GetCompleter()
}

// NewNode creates a new graph node.
func NewNode(label string, content IContent) *Node {
	nodeID++
	node := &Node{
		ID:      nodeID,
		Label:   label,
		BlockID: -1,
		Content: content,
	}
	return node
}

// NewJoint create a new graph joint node.
// Joint node is any node that does not contain information but is used
// to build the graph.
func NewJoint(label string, help string) *Node {
	jointCompleter := NewCompleterJoint(_joint)
	jointContent := NewContentJoint(label, help, jointCompleter)
	node := NewNode(label, jointContent)
	node.IsJoint = true
	return node
}

// NewRoot creates a new graph root node.
// Root node is at the top of the graph, it starts the graph and only
// can be one Root node in the graph.
func NewRoot() *Node {
	label := "ROOT"
	jointCompleter := NewCompleterJoint(_rooty)
	jointContent := NewContentJoint(label, "Root Node", jointCompleter)
	node := NewNode(label, jointContent)
	node.IsRoot = true
	return node
}

// NewSink creates a new graph sink node.
// Sink node is at the bottom of the graph, it terminates the graph only
// can be one sink node.
func NewSink() *Node {
	label := "SINK"
	jointCompleter := NewCompleterSink()
	jointContent := NewContentJoint(label, "Sink Node", jointCompleter)
	node := NewNode(label, jointContent)
	node.IsSink = true
	return node
}

// NewStart creates a new graph start node.
// Start node is used for building loop graphs, and it identifies the
// start point of the loop.
func NewStart(id int) *Node {
	label := "START"
	jointCompleter := NewCompleterStart()
	jointContent := NewContentJoint(label, "Start Node", jointCompleter)
	node := NewNode(label, jointContent)
	node.IsStart = true
	node.BlockID = id
	return node
}

// NewEnd creates a new graph end node.
// End node is used for building loop graphs, and it identifies the
// end point or exit of the loop.
func NewEnd(id int) *Node {
	label := "END"
	jointCompleter := NewCompleterEnd()
	jointContent := NewContentJoint(label, "End Node", jointCompleter)
	node := NewNode(label, jointContent)
	node.IsEnd = true
	node.BlockID = id
	return node
}

// NewLoop creates a new graph loop node.
// Loop node is used for building loop graphs, and it identfies the
// loop part which will point to the start of the loop.
func NewLoop(id int) *Node {
	label := "LOOP"
	jointCompleter := NewCompleterLoop()
	jointContent := NewContentJoint(label, "Loop Node", jointCompleter)
	node := NewNode(label, jointContent)
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
