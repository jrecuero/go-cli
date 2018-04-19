package syntax

// IContent represents any content being added to a completer.
type IContent interface {
	GetLabel() string
	GetName() string
	GetType() string
	GetDefault() interface{}
	GetHelp() string
	GetCompleter() ICompleter
}
