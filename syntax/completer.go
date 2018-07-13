package syntax

import (
	"fmt"
	"strings"

	"github.com/jrecuero/go-cli/tools"
)

const _joint = ""
const _rooty = "ROOT"
const _sink = "SINK"
const _start = "START"
const _end = "END"
const _loop = "LOOP"

// Completer represents any generic completer.
type Completer struct {
	label string
}

// GetLabel returns the label for any completer.
func (c *Completer) GetLabel() string {
	return c.label
}

// Validate checks if the content is value for the given line.
func (c *Completer) Validate(ctx *Context, content IContent, line interface{}, index int) bool {
	return true
}

// Help returns the help for any node completer.
func (c *Completer) Help(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	return nil, false
}

// Query returns the query for any node completer.
func (c *Completer) Query(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	return nil, false
}

// Complete returns the complete match for any node completer.
func (c *Completer) Complete(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	return nil, true
}

// Match returns the match for a node completer.
func (c *Completer) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	return -1, false
}

// completeLabel returns the complete match for any node completer.
func (c *Completer) completeLabel(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tokens := line.([]string)
	ilast := len(tokens) - 1
	if tokens[ilast] == "" {
		return content.GetLabel(), true
	} else if strings.HasPrefix(content.GetLabel(), tokens[ilast]) {
		return content.GetLabel(), true
	}
	return nil, false
}

// helpLabel returns the complete match for any node completer.
func (c *Completer) helpLabel(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tokens := line.([]string)
	ilast := len(tokens) - 1
	if tokens[ilast] == "" {
		return content.GetHelp(), true
	} else if strings.HasPrefix(content.GetLabel(), tokens[ilast]) {
		return content.GetHelp(), true
	}
	return nil, false
}

// NewCompleter returns a new Completer instance.
func NewCompleter(label string) *Completer {
	return &Completer{
		label: label,
	}
}

var _ ICompleter = (*Completer)(nil)

// CompleterCommand represents the completer for a command node.
type CompleterCommand struct {
	*Completer
}

// Match returns the match for a command node completer.
func (cc *CompleterCommand) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		return index + 1, true
	}
	return index, false
}

// Complete returns the complete match for any node completer.
func (cc *CompleterCommand) Complete(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	return cc.completeLabel(ctx, content, line, index)
}

// Help returns the help for any node completer.
func (cc *CompleterCommand) Help(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	return cc.helpLabel(ctx, content, line, index)
}

// NewCompleterCommand returns a new CompleterCommand instance.
func NewCompleterCommand(label string) *CompleterCommand {
	return &CompleterCommand{
		NewCompleter(label),
	}
}

var _ ICompleter = (*CompleterCommand)(nil)

// CompleterIdent represents the completer for any CompleterIdent node.
type CompleterIdent struct {
	*Completer
}

// Match returns the match for any ident node completer.
func (ci *CompleterIdent) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	tokens := line.([]string)
	if tokens[index] == ci.GetLabel() {
		return index + 1, true
	}
	return index, false
}

// Complete returns the complete match for any node completer.
func (ci *CompleterIdent) Complete(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	return ci.completeLabel(ctx, content, line, index)
}

// Help returns the help for any node completer.
func (ci *CompleterIdent) Help(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	return ci.helpLabel(ctx, content, line, index)
}

// NewCompleterIdent returns a new CompleterIdent instance.
func NewCompleterIdent(label string) *CompleterIdent {
	return &CompleterIdent{
		NewCompleter(label),
	}
}

var _ ICompleter = (*CompleterIdent)(nil)

// CompleterArgument represents the completer for CompleterArgument character sequence node.
type CompleterArgument struct {
	*Completer
}

// Match returns the match for CompleterArgument node completer.
func (ca *CompleterArgument) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	tokens := line.([]string)
	if tokens[index] == "" {
		return index, false
	}
	// When complete or help process, it should not match if it is still
	// entering the argument.
	proc := ctx.GetProcess()
	if proc == COMPLETE || proc == HELP {
		toklen := len(tokens)
		if index == (toklen - 1) {
			return index, false
		}
	}
	if ok, err := content.Validate(tokens[index]); !ok {
		tools.ERROR(err, false, "Validation ERROR: %#v\n", err)
		return index, false
	}
	return index + 1, true
}

// Complete returns the complete match for any node completer.
func (ca *CompleterArgument) Complete(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	// TODO: Query should be checked here.
	//return "<<WORD>>", true
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	return fmt.Sprintf("<<%s>>", content.GetType()), true
}

// Help returns the help for any node completer.
func (ca *CompleterArgument) Help(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	//tokens := line.([]string)
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	return content.GetHelp(), true
}

// NewCompleterArgument returns a new CompleterArgument instance.
func NewCompleterArgument(label string) *CompleterArgument {
	return &CompleterArgument{
		NewCompleter(label),
	}
}

var _ ICompleter = (*CompleterArgument)(nil)

// CompleterCustom represents the completer for CompleterCustom node.
type CompleterCustom struct {
	*Completer
}

// Match returns the match for CompleterCustom node completer.
func (c *CompleterCustom) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		return index + 1, true
	}
	return index, false
}

// NewCompleterCustom returns a new CompleterCustom instance.
func NewCompleterCustom(label string) *CompleterCustom {
	return &CompleterCustom{
		NewCompleter(label),
	}
}

var _ ICompleter = (*CompleterCustom)(nil)

// CompleterJoint represents the completer for any joint node.
type CompleterJoint struct {
	*Completer
}

// Match returns the match for any joint node completer.
func (cj *CompleterJoint) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	if content == nil || !content.IsMatchable() {
		return index, true
	}
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		return index + 1, true
	}
	return index, false
}

// NewCompleterJoint returns a new CompleterJoint instance.
func NewCompleterJoint(label string) *CompleterJoint {
	if label == "" {
		label = "JOINT"
	}
	return &CompleterJoint{
		NewCompleter(label),
	}
}

var _ ICompleter = (*CompleterJoint)(nil)

// NewCompleterStart returns a new Start instance.
func NewCompleterStart() *CompleterJoint {
	return NewCompleterJoint(_start)
}

// NewCompleterEnd returns a new End instance.
func NewCompleterEnd() *CompleterJoint {
	return NewCompleterJoint(_end)
}

// NewCompleterLoop returns a new Loop instance.
func NewCompleterLoop() *CompleterJoint {
	return NewCompleterJoint(_loop)
}

// CompleterSink represents the completer for any joint node.
type CompleterSink struct {
	*Completer
}

// Match returns the match for any joint node completer.
func (cs *CompleterSink) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		return index + 1, true
	}
	return index, false
}

// Complete returns the complete match for any node completer.
func (cs *CompleterSink) Complete(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	tokens := line.([]string)
	if tokens[len(tokens)-1] == "" {
		return content.GetLabel(), true
	}
	return nil, false
}

// Help returns the help match for any node completer.
func (cs *CompleterSink) Help(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	tokens := line.([]string)
	if tokens[len(tokens)-1] == "" {
		return content.GetHelp(), true
	}
	return nil, false
}

// NewCompleterSink returns a new sinkn completer instance.
func NewCompleterSink() *CompleterSink {
	return &CompleterSink{
		NewCompleter(_sink),
	}
}
