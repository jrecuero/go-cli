package api

// ICompleter represents the completer interface.
type ICompleter interface {
	Match(ctx interface{}, line interface{}) bool
	Help(ctx interface{}, line interface{}) (interface{}, bool)
	Complete(ctx interface{}, line interface{}) (interface{}, bool)
	GetContent() interface{}
}
