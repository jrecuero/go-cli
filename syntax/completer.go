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

// GetContent returns the content for anu completer.
func (c *Completer) GetContent() interface{} {
	return c.Content
}

// Ident represent the completer for any Ident node.
type Ident struct {
	Completer
}

// Match returns the match for any completer.
func (i *Ident) Match(ctx Context, line interface{}) bool {
	tokens := line.([]string)
	if tokens[0] == i.Label {
		return true
	}
	return false
}

// Joint represents the completer for any joint node.
type Joint struct {
	Completer
}
