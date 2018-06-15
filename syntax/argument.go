package syntax

// Argument represents any CLI argument information.
type Argument struct {
	*Content
	Type    string
	Default interface{}
}

// GetValueFromArguments returns the value for the given field in arguments
// passed to.
func GetValueFromArguments(field string, arguments interface{}) interface{} {
	argos := arguments.(map[string]interface{})
	value := argos[field]
	return value
}

// NewArgument creates a new Argument instance.
func NewArgument(label string, help string, completer ICompleter, atype string, adefault interface{}) *Argument {
	return &Argument{
		Content: NewContent(label, help, completer).(*Content),
		Type:    atype,
		Default: adefault,
	}
}

// Setup initializes all argument fields.
func (a *Argument) Setup() error {
	if a.completer == nil {
		a.completer = NewCompleterAny(a.GetLabel())
	}
	return nil
}

var _ IContent = (*Argument)(nil)
