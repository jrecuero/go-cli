package syntax

// Argument represents any CLI argument information.
type Argument struct {
	Label     string
	Type      string
	Default   interface{}
	Help      string
	Completer ICompleter
}

// GetLabel returns user argument label.
func (a *Argument) GetLabel() string {
	return a.Label
}

// GetType returns user argument type.
func (a *Argument) GetType() string {
	return a.Type
}

// GetDefault returns user argument default value.
func (a *Argument) GetDefault() interface{} {
	return a.Default
}

// GetHelp returns user argument help.
func (a *Argument) GetHelp() string {
	return a.Help
}

// GetCompleter returns user command completer.
func (a *Argument) GetCompleter() ICompleter {
	return a.Completer
}

// GetValueFromArguments returns the value for the given field in arguments
// passed to.
func GetValueFromArguments(field string, arguments interface{}) interface{} {
	argos := arguments.(map[string]interface{})
	value := argos[field]
	return value
}
