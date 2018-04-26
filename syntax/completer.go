package syntax

const _joint = ""
const _rooty = "ROOT"
const _sink = "SINK"
const _start = "START"
const _end = "END"
const _loop = "LOOP"

// Completer represents any generic completer.
type Completer struct {
	label   string
	content IContent
}

// NewCompleter returns a new Completer instance.
func NewCompleter(content IContent) *Completer {
	c := &Completer{
		label:   GetLabelFromContent(content),
		content: content,
	}
	return c
}

// GetContent returns the content for any completer.
func (c *Completer) GetContent() IContent {
	return c.content
}

// GetLabel returns the label for any completer.
func (c *Completer) GetLabel() string {
	return c.label
}

// Validate checks if the content is value for the given line.
func (c *Completer) Validate(ctx *Context, line interface{}) bool {
	return true
}

// Help returns the help for any node completer.
func (c *Completer) Help(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Query returns the query for any node completer.
func (c *Completer) Query(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Complete returns the complete match for any node completer.
func (c *Completer) Complete(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Setup the completer for the given content.
func (c *Completer) Setup(content IContent) error {
	c.content = content
	return nil
}

// Match returns the match for a node completer.
func (c *Completer) Match(ctx *Context, line interface{}, index int) (int, bool) {
	return -1, false
}

var _ ICompleter = (*Completer)(nil)

// CompleterCommand represents the completer for a command node.
type CompleterCommand struct {
	*Completer
}

// NewCompleterCommand returns a new CompleterCommand instance.
func NewCompleterCommand(content IContent) *CompleterCommand {
	i := &CompleterCommand{
		NewCompleter(content),
	}
	return i
}

// Match returns the match for a command node completer.
func (cc *CompleterCommand) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == cc.GetLabel() {
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
func NewCompleterIdent(content IContent) *CompleterIdent {
	i := &CompleterIdent{
		NewCompleter(content),
	}
	return i
}

// Match returns the match for any ident node completer.
func (i *CompleterIdent) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == i.GetLabel() {
		return index + 1, true
	}
	return index, false
}

var _ ICompleter = (*CompleterIdent)(nil)

// CompleterAny represents the completer for CompleterAny character sequence node.
type CompleterAny struct {
	*Completer
}

// NewCompleterAny returns a new CompleterAny instance.
func NewCompleterAny(content IContent) *CompleterAny {
	a := &CompleterAny{
		NewCompleter(content),
	}
	return a
}

// Match returns the match for CompleterAny node completer.
func (a *CompleterAny) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == CR.GetLabel() {
		return index, false
	}
	return index + 1, true
}

var _ ICompleter = (*CompleterAny)(nil)

// CompleterCustom represents the completer for CompleterCustom node.
type CompleterCustom struct {
	*Completer
}

// NewCompleterCustom returns a new CompleterCustom instance.
func NewCompleterCustom(content IContent) *CompleterCustom {
	c := &CompleterCustom{
		NewCompleter(content),
	}
	return c
}

// Match returns the match for CompleterCustom node completer.
func (c *CompleterCustom) Match(ctx *Context, line interface{}, index int) (int, bool) {
	tokens := line.([]string)
	if tokens[0] == c.GetContent().GetLabel() {
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
	j := &CompleterJoint{
		&Completer{
			label: label,
		},
	}
	return j
}

// Match returns the match for any joint node completer.
func (j *CompleterJoint) Match(ctx *Context, line interface{}, index int) (int, bool) {
	if j.GetContent() == nil {
		return index, true
	}
	tokens := line.([]string)
	if tokens[0] == j.GetContent().GetLabel() {
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
	j := &CompleterJoint{
		&Completer{
			label:   _sink,
			content: CR,
		},
	}
	return j
}
