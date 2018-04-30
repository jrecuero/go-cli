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

// Setup initializes all argument fields.
func (a *Argument) Setup() error {
	if a.completer == nil {
		a.completer = NewCompleterAny(a)
	} else {
		a.completer.Setup(a)
	}
	return nil
}

var _ IContent = (*Argument)(nil)
