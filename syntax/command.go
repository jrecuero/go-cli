package syntax

// Command represents any CLI command in the system.
type Command struct {
	Cb        interface{}
	Syntax    *CommandSyntax
	Label     string
	Name      string
	Help      string
	Arguments []interface{}
}

// NewCommand returns a new Command instance.
func NewCommand(name string, label string, cb interface{}) *Command {
	return &Command{
		Cb:    cb,
		Name:  name,
		Label: label,
	}
}
