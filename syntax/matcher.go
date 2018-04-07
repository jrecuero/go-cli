package syntax

import (
	"github.com/jrecuero/go-cli/graph"
)

// Matcher represents the matcher for a given graph.
type Matcher struct {
	Ctx Context
	G   *graph.Graph
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
	tokens := line.([]string)
	traverse := m.G.Root
	for traverse != nil {
		for _, n := range traverse.Children {
			if n.Completer.Match(m.Ctx, line) {
				traverse = n
				m.Ctx.AddToken(n)
				break
			}
		}
	}
	return true
}
