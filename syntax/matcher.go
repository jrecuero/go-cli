package syntax

import (
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
)

// ComplexComplete represents complete and help together.
type ComplexComplete struct {
	Complete interface{} // token complete value.
	Help     interface{} // token help value.
}

// Worker represent the function for any complete or help worker.
type Worker = func(cn *ContentNode, tokens []string) interface{}

// Matcher represents the matcher for a given graph.
type Matcher struct {
	Ctx     *Context     // context instance.
	Grapher *graph.Graph // parsing tree graph instance.
	Rooter  *graph.Node  // parsing tree root for any handling.
}

// matchCommandLine matches the given command line with the graph.
func (m *Matcher) matchCommandLine(line interface{}) bool {
	//tools.Tracer("line: %v\n", line)
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
		//tools.Debug("node: %#v check for matching: %d %s => %#v\n",
		//    NodeToContentNode(node).GetContent().GetLabel(), index, tokens[index], cn.GetContent().GetLabel())
		if indexMatched, ok := cn.Match(m.Ctx, tokens, index); ok {
			//tools.Debug("node matched: %d:%d %v %s => %v\n", indexMatched, index, ok, tokens[index], cn.GetContent().GetLabel())
			child := n
			for indexMatched == index {
				if child, indexMatched, ok = m.traverseAndMatchGraph(child, tokens, indexMatched); !ok {
					break
				}
			}
			if indexMatched != index {
				//tools.Debug("confirmed matched: %d %s => %v\n", indexMatched, tokens[index], cn.GetContent().GetLabel())
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
	//tools.Tracer("tokens: %v\n", tokens)
	traverse := m.Rooter
	for traverse != nil && len(traverse.Children) != 0 {
		if traverse, index, ok = m.traverseAndMatchGraph(traverse, tokens, index); !ok {
			return index, false
		}
		//tools.Debug("add token to context: %#v %s\n", NodeToContentNode(traverse).GetContent().GetLabel(), tokens[index-1])
		m.Ctx.AddToken(index-1, NodeToContentNode(traverse), tokens[index-1])
	}
	m.Ctx.UpdateCommandBox()
	return index, true
}

// workerComplete gets all complete options for the given node.
func (m *Matcher) workerComplete(cn *ContentNode, tokens []string) interface{} {
	result := []*ComplexComplete{}
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		completeIf, _ := childCN.Complete(m.Ctx, tokens, 0)
		complete := completeIf.([]interface{})
		//tools.Debug("childCN: %#v complete: %#v\n", childCN.GetContent().GetLabel(), complete)
		for _, c := range complete {
			result = append(result, &ComplexComplete{
				Complete: c,
			})
		}
	}
	return result
}

// workerHelp gets all help options for the given node.
func (m *Matcher) workerHelp(cn *ContentNode, tokens []string) interface{} {
	result := []*ComplexComplete{}
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		helpIf, _ := childCN.Help(m.Ctx, tokens, 0)
		help := helpIf.([]interface{})
		//tools.Debug("childCN: %#v help: %#v\n", childCN.GetContent().GetLabel(), help)
		for _, h := range help {
			result = append(result, &ComplexComplete{
				Help: h,
			})
		}
	}
	return result
}

// workerCompleteAndHelp gets all complete and help options for the given node.
func (m *Matcher) workerCompleteAndHelp(cn *ContentNode, tokens []string) interface{} {
	result := []*ComplexComplete{}
	//tools.Debug("cn: %#v\n", cn.GetContent().GetLabel())
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		completeIf, _ := childCN.Complete(m.Ctx, tokens, 0)
		helpIf, _ := childCN.Help(m.Ctx, tokens, 0)
		complete := completeIf.([]interface{})
		help := helpIf.([]interface{})
		//tools.Debug("childCN: %#v complete: %#v help: %#v\n", childCN.GetContent().GetLabel(), complete, help)
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
	var lastCN *ContentNode
	extendLine := false
	tokens := m.tokenizeLine(in)
	//tools.Tracer("tokenize-line: %#v\n", tokens)
	m.Ctx.UpdateMatched(len(tokens))
	if tools.LastChar(in.(string)) == " " {
		tokens = append(tokens, "")
		extendLine = true
	}
	index, _ := m.matchWithGraph(tokens)
	if index < (len(tokens) - 1) {
		tools.Debug("not-a-proper-match tokens: %#v index: %d len: %d\n", tokens, index, len(tokens))
		return []*ComplexComplete{}, false
	}
	//tools.Debug("len(matched): %d extended: %v index: %d\n", len(m.Ctx.Matched), extendLine, index)
	if len(m.Ctx.Matched) == 0 {
		// There is not match, this happens when it is being entered the first
		// command or the command line is empty.
		lastCN = NodeToContentNode(m.Rooter)
	} else if !extendLine && index == len(tokens) {
		ilastCN := len(m.Ctx.Matched) - 2
		if ilastCN < 0 {
			lastCN = NodeToContentNode(m.Rooter)
		} else {
			lastCN = m.Ctx.Matched[ilastCN].Node
		}

	} else {
		ilastCN := len(m.Ctx.Matched) - 1
		lastCN = m.Ctx.Matched[ilastCN].Node
	}
	result := worker(lastCN, tokens)
	//tools.Debug("line: %#v tokens: %#v results (%#v): %#v\n", line, tokens, lastCN.GetContent().GetLabel(), result)
	result = m.removeDuplicated(result)
	return result, true
}

func (m *Matcher) removeDuplicated(in interface{}) interface{} {
	tokens := in.([]*ComplexComplete)
	result := []*ComplexComplete{}
	keys := []string{}
	for _, tok := range tokens {
		//tools.Debug("tok: %#v\n", tok)
		complete := tools.ToString(tok.Complete)
		help := tools.ToString(tok.Help)
		key := complete + " " + help
		if tools.SearchKeyInTable(keys, key) != nil {
			keys = append(keys, key)
			result = append(result, tok)
		}
	}
	return result
}

func (m *Matcher) tokenizeLine(in interface{}) []string {
	//tools.Tracer("%#v\n", in)
	// This is the behavior that does not process quotes.
	//tokens := strings.Fields(line.(string))
	if in == " " {
		return []string{}
	}
	tokens := strings.Split(strings.TrimSpace(in.(string)), " ")
	//tools.Tracer("%#v\n", tokens)
	toglue := false
	var aux string
	result := []string{}
	for _, d := range tokens {
		if toglue {
			if strings.Contains(d, "\"") {
				toglue = false
				if d == "\"" {
					aux += " "
				} else {
					aux += " " + strings.Replace(d, "\"", "", -1)
				}
				result = append(result, aux)
			} else if d == "" {
				aux += " "
			} else {
				aux += " " + d
			}
		} else if strings.Contains(d, "\"") {
			toglue = true
			if d == "\"" {
				aux = " "
			} else {
				aux = strings.Replace(d, "\"", "", -1)
			}
		} else {
			result = append(result, d)
		}
	}
	if toglue {
		result = append(result, aux)
	}
	return result
}

// Match matches if a node is matched for a token.
func (m *Matcher) Match(line interface{}) (interface{}, bool) {
	tokens := m.tokenizeLine(line)
	//tools.Tracer("glue-line: %v\n", tokens)
	m.Ctx.GetProcess().Set(MATCH)
	result := m.matchCommandLine(tokens)
	m.Ctx.GetProcess().Clean()
	return nil, result
}

// Execute executes the command for the given command line.
func (m *Matcher) Execute(line interface{}) (interface{}, bool) {
	m.Ctx.GetProcess().Set(EXECUTE)
	if _, ok := m.Match(line); !ok {
		tools.ERROR(nil, true, "match return %#v for line: %#v\n", ok, line)
		return nil, false
	}
	//for _, t := range m.Ctx.Matched {
	//    tools.Debug("matched %#v\n", t)
	//}
	//for i, token := range m.Ctx.GetCommandBox() {
	//    for _, a := range token.ArgValues {
	//        tools.Debug("%d cmdbox.argvalue %#v\n", i, a)
	//    }
	//}
	for i, token := range m.Ctx.GetCommandBox() {
		cmd := token.Cmd
		//tools.Debug("%d command %#v run-as-no-final: %#v\n", i, cmd.GetLabel(), cmd.RunAsNoFinal)
		lenCommandBox := len(m.Ctx.GetCommandBox()) - 1
		if (i < lenCommandBox && cmd.RunAsNoFinal) || (i == lenCommandBox) {
			if i < lenCommandBox && cmd.RunAsNoFinal {
				m.Ctx.GetProcess().Append(RUNASNOFINAL)
			}
			args, err := m.Ctx.GetArgValuesForCommandLabel(tools.PString(cmd.GetLabel()))
			if err != nil {
				tools.ERROR(err, true, "line: %#v arguments not found: %#v\n", line, err)
				return nil, false
			}
			cmd.Enter(m.Ctx, args)
			if i < lenCommandBox && cmd.RunAsNoFinal {
				m.Ctx.GetProcess().Remove(RUNASNOFINAL)
			}
		}
	}
	if ok, _ := m.Ctx.GetProcess().Check(POPMODE); ok {
		modeBox := m.Ctx.PopMode()
		m.Rooter = modeBox.Anchor
	} else if m.Ctx.GetLastCommand().IsMode() {
		m.Ctx.PushMode(m.Rooter)
		lastAnchor := m.Ctx.GetLastAnchor()
		m.Rooter = lastAnchor
	}
	m.Ctx.GetProcess().Clean()
	m.Ctx.Clean()
	return nil, true
}

// Complete returns possible complete string for command line being entered.
func (m *Matcher) Complete(line interface{}) (interface{}, bool) {
	//tools.Tracer("line: %v\n", line)
	m.Ctx.GetProcess().Set(COMPLETE)
	complexResult, ok := m.processCompleteAndHelp(line, m.workerComplete)
	var result []interface{}
	for _, c := range complexResult.([]*ComplexComplete) {
		result = append(result, c.Complete)
	}
	m.Ctx.GetProcess().Clean()
	return result, ok
}

// Help returns the help for a node if it is matched.
func (m *Matcher) Help(line interface{}) (interface{}, bool) {
	//tools.Tracer("line: %v\n", line)
	m.Ctx.GetProcess().Set(HELP)
	complexResult, ok := m.processCompleteAndHelp(line, m.workerHelp)
	var result []interface{}
	for _, c := range complexResult.([]*ComplexComplete) {
		result = append(result, c.Help)
	}
	m.Ctx.GetProcess().Clean()
	return result, ok
}

// CompleteAndHelp returns possible complete string for command line being entered.
func (m *Matcher) CompleteAndHelp(line interface{}) (interface{}, bool) {
	//tools.Tracer("line: %v\n", line)
	m.Ctx.GetProcess().Set(COMPLETE)
	result, ok := m.processCompleteAndHelp(line, m.workerCompleteAndHelp)
	m.Ctx.GetProcess().Clean()
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
