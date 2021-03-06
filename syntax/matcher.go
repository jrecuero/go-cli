package syntax

import (
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
)

// Worker represent the function for any complete or help worker.
type Worker = func(cn *ContentNode, tokens []string) interface{}

// Matcher represents the matcher for a given graph.
type Matcher struct {
	Ctx     *Context     // context instance.
	Grapher *graph.Graph // parsing tree graph instance.
	Rooter  *graph.Node  // parsing tree root for any handling.
}

// matchCommandLine matches the given command line with the graph.
func (m *Matcher) matchCommandLine(line interface{}) error {
	//tools.Tracer("line: %v\n", line)
	tokens := line.([]string)
	tokens = append(tokens, GetCR().GetLabel())
	index, ok := m.matchWithGraph(tokens)
	if !ok || index != len(tokens) {
		return tools.ERROR(nil, false, "Command line %s failed at index %d => %s\n", line, index, tokens[index:index+1])
	}
	//for _, mt := range m.Ctx.Matched {
	//    tools.Tracer("Context matched: %s %s %v\n", mt.Node.GetContent().GetLabel(), mt.Value, mt)
	//}
	return nil
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
		// TODO: Validation should be called at this point before adding the
		// token to the Match..
		cn := NodeToContentNode(traverse)
		//tools.Debug("add token to context: %#v %s\n", cn.GetContent().GetLabel(), tokens[index-1])
		if cn.Validate(m.Ctx, strings.Join(tokens, " "), index-1) {
			m.Ctx.AddToken(index-1, cn, tokens[index-1])
		} else {
			return index - 1, false
		}
	}
	m.Ctx.UpdateCommandBox()
	return index, true
}

// workerComplete gets all complete options for the given node.
func (m *Matcher) workerComplete(cn *ContentNode, tokens []string) interface{} {
	result := []*CompleteHelp{}
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		var completeIf interface{}
		if qresult, ok := childCN.Query(m.Ctx, tokens, 0); ok && qresult != nil {
			completeIf = qresult
		} else {
			completeIf, _ = childCN.Complete(m.Ctx, tokens, 0)
		}
		complete := completeIf.([]interface{})
		//tools.Debug("childCN: %#v complete: %#v\n", childCN.GetContent().GetLabel(), complete)
		for _, c := range complete {
			result = append(result, NewCompleteHelp(c, nil))
		}
	}
	return result
}

// workerHelp gets all help options for the given node.
func (m *Matcher) workerHelp(cn *ContentNode, tokens []string) interface{} {
	result := []*CompleteHelp{}
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		var helpIf interface{}
		if qresult, ok := childCN.Query(m.Ctx, tokens, 0); ok && qresult != nil {
			helpIf = qresult
		} else {
			helpIf, _ = childCN.Help(m.Ctx, tokens, 0)
		}
		help := helpIf.([]interface{})
		//tools.Debug("childCN: %#v help: %#v\n", childCN.GetContent().GetLabel(), help)
		for _, h := range help {
			result = append(result, NewCompleteHelp(nil, h))
		}
	}
	return result
}

// workerCompleteAndHelp gets all complete and help options for the given node.
func (m *Matcher) workerCompleteAndHelp(cn *ContentNode, tokens []string) interface{} {
	result := []*CompleteHelp{}
	//tools.Debug("cn: %#v\n", cn.GetContent().GetLabel())
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		if qresult, ok := childCN.Query(m.Ctx, tokens, 0); ok && qresult != nil {
			for _, r := range qresult.([]*CompleteHelp) {
				result = append(result, r)
			}
		} else {
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
				result = append(result, NewCompleteHelp(complete[i], help[i]))
			}
		}
	}
	return result
}

//// workerQuery gets all possible values for the given node
func (m *Matcher) workerQuery(cn *ContentNode, tokens []string) interface{} {
	result := []*CompleteHelp{}
	//tools.Debug("cn: %#v\n", cn.GetContent().GetLabel())
	for _, childNode := range cn.Children {
		childCN := NodeToContentNode(childNode)
		if qresult, ok := childCN.Query(m.Ctx, tokens, 0); ok && qresult != nil {
			for _, r := range qresult.([]*CompleteHelp) {
				result = append(result, r)
			}
		}
	}
	return result
}

// processCompleteAndHelp returns possible complete string or help for the
// command line being entered.
func (m *Matcher) processCompleteAndHelp(in interface{}, worker Worker) (interface{}, error) {
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
		return []*CompleteHelp{}, tools.ERROR(nil, false, "not-a-proper-match tokens: %#v index: %d len: %d\n", tokens, index, len(tokens))
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
	return result, nil
}

func (m *Matcher) removeDuplicated(in interface{}) interface{} {
	tokens := in.([]*CompleteHelp)
	result := []*CompleteHelp{}
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
func (m *Matcher) Match(line interface{}) (interface{}, error) {
	tokens := m.tokenizeLine(line)
	//tools.Tracer("glue-line: %v\n", tokens)
	m.Ctx.GetProcess().Set(MATCH)
	err := m.matchCommandLine(tokens)
	m.Ctx.GetProcess().Clean()
	return nil, err
}

// Execute executes the command for the given command line.
func (m *Matcher) Execute(line interface{}) (interface{}, error) {
	var retcode interface{}
	m.Ctx.GetProcess().Set(EXECUTE)
	if _, err := m.Match(line); err != nil {
		return nil, tools.ERROR(err, false, "match return %#v for line: %#v\n", err, line)
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
		// Is commans is RunAsNoFinal=true it means the command callback has to
		// be executed even if it is not the final command, it means the final
		// command is a child command.
		if (i < lenCommandBox && cmd.RunAsNoFinal) || (i == lenCommandBox) {
			if i < lenCommandBox && cmd.RunAsNoFinal {
				m.Ctx.GetProcess().Append(RUNASNOFINAL)
			}
			args, err := m.Ctx.GetArgValuesForCommandLabel(tools.PString(cmd.GetLabel()))
			if err != nil {
				return nil, tools.ERROR(err, true, "line: %#v arguments not found: %#v\n", line, err)
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
		retcode = POPMODE
	} else if m.Ctx.GetLastCommand().IsMode() {
		m.Ctx.PushMode(m.Rooter)
		lastAnchor := m.Ctx.GetLastAnchor()
		m.Rooter = lastAnchor
	}
	m.Ctx.GetProcess().Clean()
	m.Ctx.Clean()
	return retcode, nil
}

// Complete returns possible complete string for command line being entered.
func (m *Matcher) Complete(line interface{}) (interface{}, error) {
	//tools.Tracer("line: %v\n", line)
	m.Ctx.GetProcess().Set(COMPLETE)
	completeAndHelp, err := m.processCompleteAndHelp(line, m.workerComplete)
	result := GetCompletes(completeAndHelp.([]*CompleteHelp))
	m.Ctx.GetProcess().Clean()
	return result, err
}

// Help returns the help for a node if it is matched.
func (m *Matcher) Help(line interface{}) (interface{}, error) {
	//tools.Tracer("line: %v\n", line)
	m.Ctx.GetProcess().Set(HELP)
	completeAndHelp, err := m.processCompleteAndHelp(line, m.workerHelp)
	result := GetHelps(completeAndHelp.([]*CompleteHelp))
	m.Ctx.GetProcess().Clean()
	return result, err
}

// CompleteAndHelp returns possible complete string for command line being entered.
func (m *Matcher) CompleteAndHelp(line interface{}) (interface{}, error) {
	//tools.Tracer("line: %v\n", line)
	m.Ctx.GetProcess().Set(COMPLETE)
	result, err := m.processCompleteAndHelp(line, m.workerCompleteAndHelp)
	m.Ctx.GetProcess().Clean()
	return result, err
}

// Query returns possible values for the given token.
func (m *Matcher) Query(line interface{}) (interface{}, error) {
	//tools.Tracer("line: %v\n", line)
	m.Ctx.GetProcess().Set(QUERY)
	result, err := m.processCompleteAndHelp(line, m.workerQuery)
	m.Ctx.GetProcess().Clean()
	return result, err
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
