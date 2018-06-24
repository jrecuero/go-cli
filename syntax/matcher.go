package syntax

import (
	"github.com/jrecuero/go-cli/graph"
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
	return nil, true
}

// Help returns the help for a node if it is matched.
func (m *Matcher) Help(line interface{}) (interface{}, bool) {
	return nil, true
}

// Complete returns the complete value if a node is matched.
func (m *Matcher) Complete(line interface{}) (interface{}, bool) {
	return nil, true
}

// MatchCommandLine matches the given command line with the graph.
func (m *Matcher) MatchCommandLine(line interface{}) bool {
	//fmt.Printf("MatchCommandLine, line: %v\n", line)
	tokens := line.([]string)
	tokens = append(tokens, GetCR().GetLabel())
	index, result := m.MatchWithGraph(tokens)
	if index != len(tokens) {
		//fmt.Printf("Command line %s failed at index %d => %s\n", line, index, tokens[index:index+1])
		return false
	}
	return result
}

// MatchWithGraph matches the given token sequence with the graph.
func (m *Matcher) MatchWithGraph(tokens []string) (int, bool) {
	var index int
	var ok bool
	//fmt.Printf("MatchWithGraph, tokens: %v\n", tokens)
	traverse := m.G.Root
	for traverse != nil && len(traverse.Children) != 0 {
		var found bool
		for _, n := range traverse.Children {
			cn := NodeToContentNode(n)
			indexMatched := index
			if index, ok = cn.Match(m.Ctx, tokens, index); ok {
				valueMatched := tokens[indexMatched]
				traverse = n
				//fmt.Printf("traverse matched: %d %s => %v\n", indexMatched, valueMatched, traverse)
				m.Ctx.AddToken(cn, valueMatched)
				found = true
				break
			}
		}
		if !found {
			return index, false
		}
	}
	return index, true
}
