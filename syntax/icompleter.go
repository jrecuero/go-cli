package syntax

// ICompleter represents the completer interface.
type ICompleter interface {
	Match(ctx *Context, line interface{}) bool
	Help(ctx *Context, line interface{}) (interface{}, bool)
	Complete(ctx *Context, line interface{}) (interface{}, bool)
	GetContent() interface{}
	GetLabel() string
}
