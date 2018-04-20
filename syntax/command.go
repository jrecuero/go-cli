package syntax

// Command represents any CLI command internally in the system.
type Command struct {
	Cb          Callback
	Syntax      string
	CmdSyntax   *CommandSyntax
	Label       string
	Help        string
	Arguments   []Argument
	Completer   ICompleter
	Namespace   string
	ToNamespace string
}

// GetLabel returns user command label.
func (c *Command) GetLabel() string {
	return c.Label
}

// GetType returns user command type.
func (c *Command) GetType() string {
	return ""
}

// GetDefault returns user command default value.
func (c *Command) GetDefault() interface{} {
	return c.Label
}

// GetHelp returns user command help.
func (c *Command) GetHelp() string {
	return c.Help
}

// GetCompleter returns user command completer.
func (c *Command) GetCompleter() ICompleter {
	return c.Completer
}

// NewCommand returns a new Command instance.
func NewCommand(label string, cb Callback) *Command {
	return &Command{
		Cb:    cb,
		Label: label,
	}
}

// Setup initializes all command fields.
func (c *Command) Setup() bool {
	return true
}
