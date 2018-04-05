package syntax

// ICommand represents any CLI command definition.
type ICommand struct {
	Syntax    string
	Cb        interface{}
	Arguments []interface{}
	Name      string
	Help      string
}

// NewICommand returns a new Command instance.
func NewICommand(name string, syntax string, cb interface{}) *ICommand {
	return &ICommand{
		Syntax: syntax,
		Cb:     cb,
		Name:   name,
	}
}
