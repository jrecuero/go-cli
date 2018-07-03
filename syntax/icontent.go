package syntax

// IContent represents any content being added to a completer.
type IContent interface {
	GetLabel() string
	GetType() string
	GetDefault() interface{}
	GetHelp() string
	GetCompleter() ICompleter
	ToString() string
	IsCommand() bool
	IsMode() bool
	IsArgument() bool
	IsJoint() bool
	IsMatchable() bool
	GetStrType() string
}
