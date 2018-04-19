package syntax

// Callback represents the type for any command callback.
type Callback func(context interface{}, arguments interface{}) bool

// UCommand represents any CLI command defined by the user..
type UCommand struct {
	Syntax    string
	Cb        Callback
	Arguments []UArgument
	Label     string
	Name      string
	Help      string
	Completer ICompleter
}

// GetLabel returns user command label.
func (uc *UCommand) GetLabel() string {
	return uc.Label
}

// GetName returns user command name.
func (uc *UCommand) GetName() string {
	return uc.Name
}

// GetType returns user command type.
func (uc *UCommand) GetType() string {
	return ""
}

// GetDefault returns user command default value.
func (uc *UCommand) GetDefault() interface{} {
	return uc.Label
}

// GetHelp returns user command help.
func (uc *UCommand) GetHelp() string {
	return uc.Help
}

// NewUcommand returns a new Command instance.
func NewUcommand(name string, syntax string, cb Callback) *UCommand {
	return &UCommand{
		Syntax: syntax,
		Cb:     cb,
		Name:   name,
	}
}
