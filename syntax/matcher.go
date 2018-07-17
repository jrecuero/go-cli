package syntax

import (
	"errors"
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
)

// ComplexComplete represents complete and help together.
type ComplexComplete struct {
	Complete interface{}
	Help     interface{}
}

// Worker represent the function for any complete or help worker.
type Worker = func(cn *ContentNode, tokens []string) interface{}

// Matcher represents the matcher for a given graph.
type Matcher struct {
	Ctx     *Context
	Grapher *graph.Graph
	Rooter  *graph.Node
}

// Match matches if a node is matched for a token.
func (m *Matcher) Match(line interface{}) (interface{}, bool) {
	tools.Tracer("line: %v\n", line)
	slice := strings.Fields(line.(string))
	m.Ctx.SetProcess(tools.PString(MATCH))
	result := m.matchCommandLine(slice)
	m.Ctx.SetProcess(nil)
	return nil, result
}

// matchCommandLine matches the given command line with the graph.
func (m *Matcher) matchCommandLine(line interface{}) bool {
	tools.Tracer("line: %v\n", line)
	tokens := line.([]string)
	tokens = append(tokens, GetCR().GetLabel())
	index, result := m.matchWithGraph(tokens)
	if index != len(tokens) {
		tools.Error("Command line %s failed at index %d => %s\n", line, index, tokens[index:index+1])
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
		tools.Debug("node check for matching: %d %s => %#v\n", index, tokens[index], cn.GetContent().GetLabel())
		if indexMatched, ok := cn.Match(m.Ctx, tokens, index); ok {
			tools.Debug("node matched: %d:%d %v %s => %v\n", indexMatched, index, ok, tokens[index], cn.GetContent().GetLabel())
			child := n
			for indexMatched == index {
				if child, indexMatched, ok = m.traverseAndMatchGraph(child, tokens, indexMatched); !ok {
					break
				}
			}
			if indexMatched != index {
				tools.Debug("confirmed matched: %d %s => %v\n", indexMatched, tokens[index], cn.GetContent().GetLabel())
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
	//traverse := m.Grapher.Root
	traverse := m.Rooter
	for traverse != nil && len(traverse.Children) != 0 {
		if traverse, index, ok = m.traverseAndMatchGraph(traverse, tokens, index); !ok {
			return index, false
		}
		tools.Debug("add token to context: %#v %s\n", NodeToContentNode(traverse).GetContent().GetLabel(), tokens[index-1])
		m.Ctx.AddToken(NodeToContentNode(traverse), tokens[index-1])
	}
	return index, true
}

// Execute executes the command for the given command line.
func (m *Matcher) Execute(line interface{}) (interface{}, bool) {
	tools.Tracer("executing line %#v\n", line)
	m.Ctx.SetProcess(tools.PString(EXECUTE))
	if _, ok := m.Match(line); !ok {
		tools.ERROR(errors.New("token match error"), true, "match return %#v for line: %s", ok, line)
		return nil, false
	}
	args, err := m.Ctx.GetArgValuesForCommandLabel(nil)
	if err != nil {
		tools.ERROR(err, true, "line: %#v arguments not found: %#v\n", line, err)
		return nil, false
	}
	command := m.Ctx.GetLastCommand()
	command.Enter(m.Ctx, args)
	if m.Ctx.GetProcess() == POPMODE {
		modeBox := m.Ctx.PopMode()
		m.Rooter = modeBox.Anchor
	} else if m.Ctx.GetLastCommand().IsMode() {
		m.Ctx.PushMode(m.Rooter)
		lastAnchor := m.Ctx.GetLastAnchor()
		m.Rooter = lastAnchor
	}
	m.Ctx.SetProcess(nil)
	m.Ctx.Clean()
	return nil, true
}

// workerComplete gets all complete options for the given node.
func (m *Matcher) workerComplete(cn *ContentNode, tokens []string) interface{} {
	result := []interface{}{}
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		completeIf, _ := childCN.Complete(m.Ctx, tokens, 0)
		complete := completeIf.([]interface{})
		tools.Debug("childCN: %#v complete: %#v\n", childCN.GetContent().GetLabel(), complete)
		for _, c := range complete {
			result = append(result, c)
		}
	}
	return result
}

// workerHelp gets all help options for the given node.
func (m *Matcher) workerHelp(cn *ContentNode, tokens []string) interface{} {
	result := []interface{}{}
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		helpIf, _ := childCN.Help(m.Ctx, tokens, 0)
		help := helpIf.([]interface{})
		tools.Debug("childCN: %#v help: %#v\n", childCN.GetContent().GetLabel(), help)
		for _, c := range help {
			result = append(result, c)
		}
	}
	return result
}

// workerCompleteAndHelp gets all complete and help options for the given node.
func (m *Matcher) workerCompleteAndHelp(cn *ContentNode, tokens []string) interface{} {
	result := []*ComplexComplete{}
	tools.Debug("cn: %#v\n", cn.GetContent().GetLabel())
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		completeIf, _ := childCN.Complete(m.Ctx, tokens, 0)
		helpIf, _ := childCN.Help(m.Ctx, tokens, 0)
		complete := completeIf.([]interface{})
		help := helpIf.([]interface{})
		tools.Debug("childCN: %#v complete: %#v help: %#v\n", childCN.GetContent().GetLabel(), complete, help)
		limit := len(complete)
		if limit > len(help) {
			limit = len(help)
		}
		for i := 0; i < limit; i++ {
			result = append(result, &ComplexComplete{
				Complete: complete[i],
				Help:     help[i],
			})
		}
	}
	return result
}

// processCompleteAndHelp returns possible complete string or help for the
// command line being entered.
func (m *Matcher) processCompleteAndHelp(in interface{}, worker Worker) (interface{}, bool) {
	line := in.(string)
	var tokens []string
	var lastCN *ContentNode
	tokens = strings.Fields(line)
	if tools.LastChar(line) == " " {
		tokens = append(tokens, "")
	}
	m.matchWithGraph(tokens)
	if len(m.Ctx.Matched) == 0 {
		// There is not match, this happens when it is being entered the first
		// command or the command line is empty.
		//lastCN = NodeToContentNode(m.Grapher.Root)
		lastCN = NodeToContentNode(m.Rooter)
	} else {
		ilastCN := len(m.Ctx.Matched) - 1
		lastCN = m.Ctx.Matched[ilastCN].Node
	}
	result := worker(lastCN, tokens)
	tools.Debug("line: %#v\n", line)
	tools.Debug("tokens: %#v\n", tokens)
	tools.Debug("results (%#v): %#v\n", lastCN.GetContent().GetLabel(), result)
	return result, true
}

// Complete returns possible complete string for command line being entered.
func (m *Matcher) Complete(in interface{}) (interface{}, bool) {
	tools.Tracer("line: %v\n", in)
	m.Ctx.SetProcess(tools.PString(COMPLETE))
	result, ok := m.processCompleteAndHelp(in, m.workerComplete)
	m.Ctx.SetProcess(nil)
	return result, ok
}

// Help returns the help for a node if it is matched.
func (m *Matcher) Help(in interface{}) (interface{}, bool) {
	tools.Tracer("line: %v\n", in)
	m.Ctx.SetProcess(tools.PString(HELP))
	result, ok := m.processCompleteAndHelp(in, m.workerHelp)
	m.Ctx.SetProcess(nil)
	return result, ok
}

// CompleteAndHelp returns possible complete string for command line being entered.
func (m *Matcher) CompleteAndHelp(in interface{}) (interface{}, bool) {
	tools.Tracer("line: %v\n", in)
	m.Ctx.SetProcess(tools.PString(COMPLETE))
	result, ok := m.processCompleteAndHelp(in, m.workerCompleteAndHelp)
	m.Ctx.SetProcess(nil)
	return result, ok
}

// NewMatcher creates a new Matcher instance.
func NewMatcher(ctx *Context, grapher *graph.Graph) *Matcher {
	m := &Matcher{
		Ctx:     ctx,
		Grapher: grapher,
		Rooter:  grapher.Root,
	}
	return m
}
