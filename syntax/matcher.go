package syntax

import (
	"fmt"
)

// CR represents the carrier return token
const CR = "<<<_CR_>>>"

// Matcher represents the matcher for a given graph.
type Matcher struct {
	Ctx *Context
	G   *Graph
}

// NewMatcher creates a new Matcher instance.
func NewMatcher(ctx *Context, g *Graph) *Matcher {
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

// MatchWithGraph matches the given line with the graph.
func (m *Matcher) MatchWithGraph(line interface{}) bool {
	var index int
	var ok bool
	fmt.Printf("MatchWithGraph, line: %v\n", line)
	tokens := line.([]string)
	tokens = append(tokens, CR)
	traverse := m.G.Root
	for traverse != nil {
		fmt.Printf("traverse: %v\n", traverse)
		if len(traverse.Children) == 0 {
			break
		}
		for _, n := range traverse.Children {
			index, ok = n.Completer.Match(m.Ctx, tokens[index:], index)
			if ok {
				traverse = n
				m.Ctx.AddToken(n)
				break
			}
		}
	}
	return true
}
