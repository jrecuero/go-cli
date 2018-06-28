package syntax

import (
	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
)

// ContentNode represents any node which content is IContent.
type ContentNode struct {
	*graph.Node
}

// GetContent returns the Content field as IContent.
func (cn *ContentNode) GetContent() IContent {
	return cn.Content.(IContent)
}

// Match returns the match for content node.
func (cn *ContentNode) Match(ctx interface{}, line interface{}, index int) (int, bool) {
	content := cn.GetContent()
	tools.Log().Printf("%s %v %d %#v\n", content.GetLabel(), line, index, content.GetCompleter())
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		return completer.Match(context, content, line, index)
	}
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		return index + 1, true
	}
	return index, false
}

// Help returns the help for content node.
func (cn *ContentNode) Help(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	content := cn.GetContent()
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		return completer.Help(context, content, line, index)
	}
	return content.GetLabel(), true
}

// Query returns the query for content node.
func (cn *ContentNode) Query(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	content := cn.GetContent()
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		return completer.Query(context, content, line, index)
	}
	return nil, true
}

// Complete returns the complete match for content node.
func (cn *ContentNode) Complete(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	content := cn.GetContent()
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		return completer.Complete(context, content, line, index)
	}
	return content.GetLabel(), true
}

// Validate checks if the content is value for the given line.
func (cn *ContentNode) Validate(ctx interface{}, line interface{}, index int) bool {
	content := cn.GetContent()
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		return completer.Validate(context, content, line, index)
	}
	return true
}

// NewContentNode creates a new content node instance.
func NewContentNode(label string, content IContent) *ContentNode {
	return &ContentNode{
		graph.NewNode(label, content),
	}
}

// ContentNodeToNode casts a ContentNode to a Node
func ContentNodeToNode(cn *ContentNode) *graph.Node {
	return cn.Node
}

// NodeToContentNode casts a Node to a ContentNode
func NodeToContentNode(n *graph.Node) *ContentNode {
	return &ContentNode{
		Node: n,
	}
}
