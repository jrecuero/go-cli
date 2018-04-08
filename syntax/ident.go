package syntax

// Ident represent the completer for any Ident node.
type Ident struct {
	Label   string
	Content interface{}
}

// GetContent returns the content for anu completer.
func (i *Ident) GetContent() interface{} {
	return i.Content
}

// Match returns the match for any completer.
func (i *Ident) Match(ctx Context, line interface{}) bool {
	tokens := line.([]string)
	if tokens[0] == i.Label {
		return true
	}
	return false
}
