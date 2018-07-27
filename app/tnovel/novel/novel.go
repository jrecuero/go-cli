package novel

import (
	"strings"

	"github.com/jrecuero/go-cli/parser"
	lexnovel "github.com/jrecuero/go-cli/parser/lex/novel"
	"github.com/jrecuero/go-cli/tools"
)

// Actor represents any actor in the app.
type Actor struct {
	Name     string
	Life     int
	Strength int
}

// NewActor creates a new actor instance.
func NewActor(name string, life int, strength int) *Actor {
	return &Actor{
		Name:     name,
		Life:     life,
		Strength: strength,
	}
}

// Novel represents the main object for the app.
type Novel struct{}

// ActionExec represents any action in the app.
type ActionExec struct {
	origins []string
	actions []string
	targets []string
}

// AddOrigin adds a new origin
func (ae *ActionExec) AddOrigin(in string) *ActionExec {
	ae.origins = append(ae.origins, in)
	return ae
}

// AddAction adds a new action.
func (ae *ActionExec) AddAction(in string) *ActionExec {
	ae.actions = append(ae.actions, in)
	return ae
}

// AddTarget adds a new target.
func (ae *ActionExec) AddTarget(in string) *ActionExec {
	ae.targets = append(ae.targets, in)
	return ae
}

// Parse parses the novel action.
func (n *Novel) Parse(line string) *lexnovel.Syntax {
	parser := parser.NewParser(strings.NewReader(line), lexnovel.NewParser())
	result, _ := parser.Parse()
	return result.(*lexnovel.Syntax)
}

// CompileStatus represents the compile status.
type CompileStatus struct {
	actual int
}

// Next moves to the next status if condition is true.
func (cs *CompileStatus) Next(condition bool) *CompileStatus {
	if condition {
		cs.actual++
	}
	return cs
}

// Value returns compile status value.
func (cs *CompileStatus) Value() int {
	return cs.actual
}

// NewCompileStatus creates a new compile status instance.
func NewCompileStatus(start int) *CompileStatus {
	return &CompileStatus{start}
}

// Compile translate the parsing action to a struct.
func (n *Novel) Compile(line string) *ActionExec {
	defer func() {
		if r := recover(); r != nil {
			tools.ToDisplay("Error: %#v\n", r)
		}
	}()
	const (
		parsingOrigin int = iota
		parsingAction
		parsingTarget
		parsingEnd
	)
	parsed := n.Parse(line)
	tools.ToDisplay("%#v\n", parsed)
	ae := &ActionExec{}
	bracketed := false
	status := NewCompileStatus(parsingOrigin)
	for i, token := range parsed.Tokens {
		str := parsed.Idents[i]
		switch token {
		case parser.IDENT:
			switch status.Value() {
			case parsingOrigin:
				ae.AddOrigin(str)
				status.Next(!bracketed)
				break
			case parsingAction:
				ae.AddAction(str)
				status.Next(!bracketed)
				break
			case parsingTarget:
				ae.AddTarget(str)
				status.Next(!bracketed)
			default:
				panic("bad action: unknown status")
				break
			}

		case lexnovel.OPENBRACKET:
			if bracketed {
				panic("bad action: too many brackets")
			}
			bracketed = true
			break
		case lexnovel.CLOSEBRACKET:
			if !bracketed {
				panic("bad action: no brackets")
			}
			bracketed = false
			status.Next(true)
			break
		default:
			panic("bad action: unknown token")
			break
		}
	}
	return ae
}
