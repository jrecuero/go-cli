package syntax

import (
	"fmt"
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
)

// Matcher represents the matcher for a given graph.
type Matcher struct {
	Ctx *Context
	G   *graph.Graph
}

// NewMatcher creates a new Matcher instance.
func NewMatcher(ctx *Context, g *graph.Graph) *Matcher {
	m := &Matcher{
		Ctx: ctx,
		G:   g,
	}
	return m
}

// Match matches if a node is matched for a token.
func (m *Matcher) Match(line interface{}) (interface{}, bool) {
	slice := strings.Fields(line.(string))
	return nil, m.matchCommandLine(slice)
}

// Help returns the help for a node if it is matched.
func (m *Matcher) Help(line interface{}) (interface{}, bool) {
	return nil, true
}

// matchCommandLine matches the given command line with the graph.
func (m *Matcher) matchCommandLine(line interface{}) bool {
	tools.Tracer("line: %v\n", line)
	tokens := line.([]string)
	tokens = append(tokens, GetCR().GetLabel())
	index, result := m.matchWithGraph(tokens)
	if index != len(tokens) {
		tools.Tracer("Command line %s failed at index %d => %s\n", line, index, tokens[index:index+1])
		return false
	}
	//for _, mt := range m.Ctx.Matched {
	//    tools.Tracer("Context matched: %s %s %v\n", mt.Node.GetContent().GetLabel(), mt.Value, mt)
	//}
	return result
}

// traverseAndMatchGraph finds a match in the graph for the given tokens.
func (m *Matcher) traverseAndMatchGraph(node *graph.Node, tokens []string, index int) (*graph.Node, int, bool) {
	if index >= len(tokens) {
		return nil, index, false
	}
	for _, n := range node.Children {
		cn := NodeToContentNode(n)
		tools.Tracer("node check for matching: %d %s => %#v\n", index, tokens[index], cn.GetContent().GetLabel())
		if indexMatched, ok := cn.Match(m.Ctx, tokens, index); ok {
			tools.Tracer("node matched: %d:%d %v %s => %v\n", indexMatched, index, ok, tokens[index], cn.GetContent().GetLabel())
			child := n
			for indexMatched == index {
				if child, indexMatched, ok = m.traverseAndMatchGraph(child, tokens, indexMatched); !ok {
					break
				}
			}
			if indexMatched != index {
				tools.Tracer("confirmed matched: %d %s => %v\n", indexMatched, tokens[index], cn.GetContent().GetLabel())
				return child, indexMatched, true
			}
		}
	}
	return nil, index, false
}

// matchWithGraph matches the given token sequence with the graph.
func (m *Matcher) matchWithGraph(tokens []string) (int, bool) {
	var index int
	var ok bool
	tools.Tracer("tokens: %v\n", tokens)
	traverse := m.G.Root
	for traverse != nil && len(traverse.Children) != 0 {
		if traverse, index, ok = m.traverseAndMatchGraph(traverse, tokens, index); !ok {
			return index, false
		}
		tools.Tracer("add token to context: %#v %s\n", NodeToContentNode(traverse).GetContent().GetLabel(), tokens[index-1])
		m.Ctx.AddToken(NodeToContentNode(traverse), tokens[index-1])
	}
	return index, true
}

// Complete returns possible complete string for command line being entered.
func (m *Matcher) Complete(in interface{}) (interface{}, bool) {
	line := in.(string)
	tokens := strings.Fields(line)
	if tools.LastChar(line) == " " {
		tokens = append(tokens, "")
	}
	m.matchWithGraph(tokens)
	ilastCN := len(m.Ctx.Matched) - 1
	lastCN := m.Ctx.Matched[ilastCN].Node
	//result := []interface{}{}
	//ilast := len(tokens) - 1
	//if tokens[ilast] == "" {
	//result, _ := lastCN.Complete(m.Ctx, tokens, 0)
	result := []interface{}{}
	for _, childNode := range lastCN.Children {
		childCN := NodeToContentNode(childNode)
		//childContent := childCN.GetContent()
		//complet, _ := childContent.GetCompleter().Complete(m.Ctx, childContent, tokens, 0)
		//result = append(result, complet)
		complet, _ := childCN.Complete(m.Ctx, tokens, 0)
		for _, c := range complet.([]interface{}) {
			result = append(result, c)
		}
	}
	fmt.Printf("line: %#v\n", line)
	fmt.Printf("tokens: %#v\n", tokens)
	fmt.Printf("complete result (%#v): %#v\n", lastCN.GetContent().GetLabel(), result)
	return result, true
	//} else {
	//    result, _ := lastCN.GetContent().GetCompleter().Complete(m.Ctx, lastCN.GetContent(), tokens, 0)
	//    fmt.Printf("complete result: %#v\n", result)
	//    return result, true
	//}
	//return nil, false
}
