package syntax

// ICompleter represents the completer interface.
type ICompleter interface {
	Match(ctx *Context, content IContent, line interface{}, index int) (int, bool)
	Help(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool)
	Query(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool)
	Complete(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool)
	Validate(ctx *Context, content IContent, line interface{}, index int) bool
	GetLabel() string
}
