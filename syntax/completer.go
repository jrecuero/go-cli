package syntax

// ICompleter represents the completer interface.
type ICompleter interface {
	Match(ctx Context, line interface{}) bool
	Help(ctx Context, line interface{}) (interface{}, bool)
	Complete(ctx Context, line interface{}) (interface{}, bool)
	GetContent() interface{}
}

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
	tokens = line([]string)
	if tokens[0] == i.Label {
		return true
	}
	return false
}
