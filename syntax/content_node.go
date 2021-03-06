package syntax

import (
	"bytes"
	"fmt"
	"reflect"

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

// slicing transforms data into a slice.
func (cn *ContentNode) slicing(in interface{}) []interface{} {
	result := []interface{}{}
	rt := reflect.TypeOf(in)
	switch rt.Kind() {
	case reflect.Slice:
		for _, r := range in.([]interface{}) {
			result = append(result, r)
		}
		break
	case reflect.Array:
		for _, r := range in.([]interface{}) {
			result = append(result, r)
		}
		break
	default:
		result = append(result, in)
		break
	}
	return result
}

// Match returns the match for content node.
func (cn *ContentNode) Match(ctx interface{}, line interface{}, index int) (int, bool) {
	content := cn.GetContent()
	//tools.Tracer("%s %v %d %#v\n", content.GetLabel(), line, index, content.GetCompleter())
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

// Query returns the query for content node.
func (cn *ContentNode) Query(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	content := cn.GetContent()
	//tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		return completer.Query(context, content, line, index)
	}
	return nil, true
}

// Help returns the help for content node.
func (cn *ContentNode) Help(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	content := cn.GetContent()
	//tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		result := []interface{}{}
		if cn.IsContent() || cn.IsSink {
			if ret, ok := cn.GetContent().GetCompleter().Help(context, cn.GetContent(), line, 0); ok {
				result = cn.slicing(ret)
			}
		} else {
			for _, childNode := range cn.Children {
				childCN := NodeToContentNode(childNode)
				help, _ := childCN.Help(context, line, 0)
				for _, c := range help.([]interface{}) {
					result = append(result, c)
				}
			}
		}
		//tools.Debug("result is %#v\n", result)
		return result, true
	}
	return content.GetLabel(), true
}

// Complete returns the complete match for content node.
func (cn *ContentNode) Complete(ctx interface{}, line interface{}, index int) (interface{}, bool) {
	content := cn.GetContent()
	//tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	if completer := content.GetCompleter(); completer != nil {
		//tools.Debug("cn: %#v\n", content.GetLabel())
		context := ctx.(*Context)
		result := []interface{}{}
		if cn.IsContent() || cn.IsSink {
			if ret, ok := cn.GetContent().GetCompleter().Complete(context, cn.GetContent(), line, 0); ok {
				result = cn.slicing(ret)
			}
		} else {
			for _, childNode := range cn.Children {
				childCN := NodeToContentNode(childNode)
				complet, _ := childCN.Complete(context, line, 0)
				for _, c := range complet.([]interface{}) {
					result = append(result, c)
				}
			}
		}
		//tools.Debug("results: %#v\n", result)
		return result, true
	}
	return content.GetLabel(), true
}

// Validate checks if the content is value for the given line.
func (cn *ContentNode) Validate(ctx interface{}, line interface{}, index int) bool {
	content := cn.GetContent()
	//tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	if completer := content.GetCompleter(); completer != nil {
		context := ctx.(*Context)
		return completer.Validate(context, content, line, index)
	}
	return true
}

// ToContent returns node content information.
func (cn *ContentNode) ToContent() string {
	var buffer bytes.Buffer
	pattern := "[%-20s]\t%#v\n"
	if cn == nil || cn.Content == nil {
		buffer.WriteString(fmt.Sprintf(pattern, "nil", "nil"))
	} else {
		c := cn.GetContent()
		buffer.WriteString(fmt.Sprintf(pattern, tools.GetReflectType(c), c.GetLabel()))

	}
	return buffer.String()
}

// ToContentChildren returns children node content information.
func (cn *ContentNode) ToContentChildren() string {
	var buffer bytes.Buffer
	for _, child := range cn.Children {
		buffer.WriteString(NodeToContentNode(child).ToContent())
	}
	return buffer.String()
}

// NewContentNode creates a new content node instance.
func NewContentNode(label string, content IContent) *ContentNode {
	cn := &ContentNode{
		graph.NewNode(label, content),
	}
	cn.GraphPattern = content.GetGraphPattern()
	return cn
}

// ContentNodeToNode casts a ContentNode to a Node
func ContentNodeToNode(cn *ContentNode) *graph.Node {
	if cn == nil {
		return nil
	}
	return cn.Node
}

// NodeToContentNode casts a Node to a ContentNode
func NodeToContentNode(n *graph.Node) *ContentNode {
	if n == nil {
		return nil
	}
	return &ContentNode{
		Node: n,
	}
}
