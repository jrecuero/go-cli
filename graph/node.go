package graph

import (
	"bytes"
	"fmt"

	"github.com/jrecuero/go-cli/tools"
)

var masterID int

// Node represents a node in the graph.
type Node struct {
	id            int
	Label         string
	Children      []*Node
	IsRoot        bool
	IsSink        bool
	IsStart       bool
	IsEnd         bool
	IsLoop        bool
	IsJoint       bool
	IsNext        bool
	InPath        bool
	BlockID       int
	AllowChildren bool
	Content       interface{}
}

// IsContent checks if a Node contains data content information or not.
func (n *Node) IsContent() bool {
	return !(n.IsRoot || n.IsSink || n.IsStart || n.IsEnd || n.IsLoop || n.IsJoint || n.IsNext)
}

// AddChild adds a new child node.
func (n *Node) AddChild(child *Node) bool {
	tools.Tracer("AddChild:\n\tn: %#v\n\tchild: %#v\n", n, child)
	if n.AllowChildren {
		n.Children = append(n.Children, child)
		return true
	}
	return false
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
			buffer.WriteString(fmt.Sprintf("%s-%d((%s))", n.Label, n.id, n.Label))
		} else {
			buffer.WriteString(fmt.Sprintf("%s-%d[%s]", n.Label, n.id, n.Label))
		}
	} else {
		if n.IsJoint == true {
			buffer.WriteString(fmt.Sprintf("%s-%d-%d((%s))", n.Label, n.id, n.BlockID, n.Label))
		} else {
			buffer.WriteString(fmt.Sprintf("%s-%d-%d[%s]", n.Label, n.id, n.BlockID, n.Label))
		}
	}
	return buffer.String()
}

// ToMermaidChildren returns node in Mermaid graph format.
func (n *Node) ToMermaidChildren() string {
	var buffer bytes.Buffer
	for _, child := range n.Children {
		buffer.WriteString(fmt.Sprintf("%s --> %s\n", n.mermaidLabel(), child.mermaidLabel()))
	}
	return buffer.String()
}

// ToContent returns node content information.
func (n *Node) ToContent() string {
	var buffer bytes.Buffer
	c := n.Content
	pattern := "[%-20s]\t%#v\n"
	buffer.WriteString(fmt.Sprintf(pattern, tools.GetReflectType(c), c))
	return buffer.String()
}

// ToContentChildren returns children node content information.
func (n *Node) ToContentChildren() string {
	var buffer bytes.Buffer
	for _, child := range n.Children {
		buffer.WriteString(child.ToContent())
	}
	return buffer.String()
}

// Match returns the match for a node.
func (n *Node) Match(ctx interface{}, line interface{}, index int) (int, bool) {
	return index, true
}

// Help returns the help for any node.
func (n *Node) Help(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	return n.Content, true
}

// Query returns the query for any node.
func (n *Node) Query(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	return nil, true
}

// Complete returns the complete match for any node.
func (n *Node) Complete(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	return n.Content, true
}

// Validate checks if the content is value for the given line.
func (n *Node) Validate(ctx interface{}, line interface{}, index int) bool {
	return true
}

// NewNode creates a new graph node.
func NewNode(label string, content interface{}) *Node {
	masterID++
	node := &Node{
		id:            masterID,
		Label:         label,
		BlockID:       -1,
		Content:       content,
		AllowChildren: true,
	}
	return node
}

// NewNodeJoint create a new graph joint node.
// Joint node is any node that does not contain information but is used
// to build the graph.
func NewNodeJoint(label string, content interface{}) *Node {
	node := NewNode(label, content)
	node.IsJoint = true
	return node
}

// NewNodeRoot creates a new graph root node.
// Root node is at the top of the graph, it starts the graph and only
// can be one Root node in the graph.
func NewNodeRoot(content interface{}) *Node {
	node := NewNodeJoint("ROOT", content)
	node.IsRoot = true
	return node
}

// NewNodeSink creates a new graph sink node.
// Sink node is at the bottom of the graph, it terminates the graph only
// can be one sink node.
func NewNodeSink(content interface{}) *Node {
	if content != nil {
		node := NewNodeJoint("SINK", content)
		node.IsSink = true
		node.AllowChildren = false
		return node
	}
	return nil
}

// NewNodeNext creates a new graph next node.
// Next node is at the botton of the graph, it allows the command to be
// chained with the next command.
func NewNodeNext(content interface{}) *Node {
	if content != nil {
		node := NewNodeJoint("NEXT", content)
		node.IsNext = true
		return node
	}
	return nil
}

// NewNodeStart creates a new graph start node.
// Start node is used for building loop graphs, and it identifies the
// start point of the loop.
func NewNodeStart(id int, content interface{}) *Node {
	node := NewNodeJoint("START", content)
	node.IsStart = true
	node.BlockID = id
	return node
}

// NewNodeEnd creates a new graph end node.
// End node is used for building loop graphs, and it identifies the
// end point or exit of the loop.
func NewNodeEnd(id int, content interface{}) *Node {
	node := NewNodeJoint("END", content)
	node.IsEnd = true
	node.BlockID = id
	return node
}

// NewNodeLoop creates a new graph loop node.
// Loop node is used for building loop graphs, and it identfies the
// loop part which will point to the start of the loop.
func NewNodeLoop(id int, content interface{}) *Node {
	node := NewNodeJoint("LOOP", content)
	node.IsLoop = true
	node.BlockID = id
	return node
}
