package syntax

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

// NewCompleter returns a new Completer instance.
func NewCompleter(label string) *Completer {
	return &Completer{
		label: label,
	}
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
	return nil, false
}

// Match returns the match for a node completer.
func (c *Completer) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	return -1, false
}

var _ ICompleter = (*Completer)(nil)

// CompleterCommand represents the completer for a command node.
type CompleterCommand struct {
	*Completer
}

// NewCompleterCommand returns a new CompleterCommand instance.
func NewCompleterCommand(label string) *CompleterCommand {
	return &CompleterCommand{
		NewCompleter(label),
	}
}

// Match returns the match for a command node completer.
func (cc *CompleterCommand) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		//ctx.SetLastCommand(content.(*Command))
		return index + 1, true
	}
	return index, false
}

var _ ICompleter = (*CompleterCommand)(nil)

// CompleterIdent represents the completer for any CompleterIdent node.
type CompleterIdent struct {
	*Completer
}

// NewCompleterIdent returns a new CompleterIdent instance.
func NewCompleterIdent(label string) *CompleterIdent {
	return &CompleterIdent{
		NewCompleter(label),
	}
}

// Match returns the match for any ident node completer.
func (i *CompleterIdent) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[index] == i.GetLabel() {
		return index + 1, true
	}
	return index, false
}

var _ ICompleter = (*CompleterIdent)(nil)

// CompleterArgument represents the completer for CompleterArgument character sequence node.
type CompleterArgument struct {
	*Completer
}

// NewCompleterArgument returns a new CompleterArgument instance.
func NewCompleterArgument(label string) *CompleterArgument {
	return &CompleterArgument{
		NewCompleter(label),
	}
}

// Match returns the match for CompleterArgument node completer.
func (ca *CompleterArgument) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	//tokens := line.([]string)
	//ctx.SetLastArgument(content.(*Argument))
	return index + 1, true
}

var _ ICompleter = (*CompleterArgument)(nil)

// CompleterCustom represents the completer for CompleterCustom node.
type CompleterCustom struct {
	*Completer
}

// NewCompleterCustom returns a new CompleterCustom instance.
func NewCompleterCustom(label string) *CompleterCustom {
	return &CompleterCustom{
		NewCompleter(label),
	}
}

// Match returns the match for CompleterCustom node completer.
func (c *CompleterCustom) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		return index + 1, true
	}
	return index, false
}

var _ ICompleter = (*CompleterCustom)(nil)

// CompleterJoint represents the completer for any joint node.
type CompleterJoint struct {
	*Completer
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

// Match returns the match for any joint node completer.
func (j *CompleterJoint) Match(ctx *Context, content IContent, line interface{}, index int) (int, bool) {
	//tools.Tracer("CompleterJoint:Match:content: %#v\n", content.GetLabel())
	if content == nil || !content.IsMatchable() {
		return index, true
	}
	tokens := line.([]string)
	if tokens[index] == content.GetLabel() {
		return index + 1, true
	}
	return index, false
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

// NewCompleterSink returns a new sinkn completer instance.
func NewCompleterSink() *CompleterJoint {
	return NewCompleterJoint(_sink)
}
