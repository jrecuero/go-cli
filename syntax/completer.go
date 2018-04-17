package syntax

// Completer represent any generic completer.
type Completer struct {
	Label   string
	Content interface{}
}

// NewCompleter returns a new Completer instance.
func NewCompleter(label string, content interface{}) *Completer {
	c := &Completer{
		Label:   label,
		Content: content,
	}
	return c
}

// GetContent returns the content for any completer.
func (c Completer) GetContent() interface{} {
	return c.Content
}

// GetLabel returns the label for any completer.
func (c Completer) GetLabel() string {
	return c.Label
}

// Ident represent the completer for any Ident node.
type Ident struct {
	Completer
}

// NewIdentCompleter returns a new Ident instance.
func NewIdentCompleter(label string, content interface{}) *Ident {
	i := &Ident{
		Completer{
			Label:   label,
			Content: content,
		},
	}
	return i
}

// Match returns the match for any ident node completer.
func (i Ident) Match(ctx *Context, line interface{}) bool {
	tokens := line.([]string)
	if tokens[0] == i.Label {
		return true
	}
	return false
}

// Help returns the help for any ident node completer.
func (i Ident) Help(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Complete returns the complete match for any ident node completer.
func (i Ident) Complete(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Joint represents the completer for any joint node.
type Joint struct {
	Completer
}

// NewJointCompleter returns a new Joint instance.
func NewJointCompleter(label string) *Joint {
	if label == "" {
		label = "JOINT"
	}
	j := &Joint{
		Completer{
			Label: label,
		},
	}
	return j
}

// GetContent returns the content for any joint node completer.
func (j Joint) GetContent() interface{} {
	return nil
}

// Match returns the match for any joint node completer.
func (j Joint) Match(ctx *Context, line interface{}) bool {
	return false
}

// Help returns the help for any joint node completer.
func (j Joint) Help(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// Complete returns the complete match for any joint node completer.
func (j Joint) Complete(ctx *Context, line interface{}) (interface{}, bool) {
	return nil, false
}

// NewStartCompleter returns a new Start instance.
func NewStartCompleter() *Joint {
	return NewJointCompleter("START")
}

// NewEndCompleter returns a new End instance.
func NewEndCompleter() *Joint {
	return NewJointCompleter("END")
}

// NewLoopCompleter returns a new Loop instance.
func NewLoopCompleter() *Joint {
	return NewJointCompleter("LOOP")
}
