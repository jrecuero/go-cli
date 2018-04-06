package syntax

// Callback represents the type for any command callback.
type Callback func(context interface{}, arguments []interface{}) bool

// Command represents any CLI command in the system.
type Command struct {
	Cb        Callback
	Syntax    *CommandSyntax
	Label     string
	Name      string
	Help      string
	Arguments []interface{}
}

// NewCommand returns a new Command instance.
func NewCommand(name string, label string, cb Callback) *Command {
	return &Command{
		Cb:    cb,
		Name:  name,
		Label: label,
	}
}
