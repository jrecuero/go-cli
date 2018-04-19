package syntax

// UArgument represents any CLI argument information.
type UArgument struct {
	Name      string
	Label     string
	Type      string
	Default   interface{}
	Help      string
	Completer ICompleter
}

// GetLabel returns user argument label.
func (ua *UArgument) GetLabel() string {
	return ua.Label
}

// GetName returns user argument name.
func (ua *UArgument) GetName() string {
	return ua.Name
}

// GetType returns user argument type.
func (ua *UArgument) GetType() string {
	return ua.Type
}

// GetDefault returns user argument default value.
func (ua *UArgument) GetDefault() interface{} {
	return ua.Default
}

// GetHelp returns user argument help.
func (ua *UArgument) GetHelp() string {
	return ua.Help
}

// GetCompleter returns user command completer.
func (ua *UArgument) GetCompleter() ICompleter {
	return ua.Completer
}
