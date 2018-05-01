package syntax

import "errors"

// NameSpace represents a namespace for commands.
// It organizes commands in static conatiners called namespaces.
type NameSpace struct {
	Name     string
	commands []*Command
}

// NewNameSpace returns a new namespace instance.
func NewNameSpace(name string) *NameSpace {
	return &NameSpace{
		Name: name,
	}
}

// GetCommands returns all commands for the namespace.
func (ns *NameSpace) GetCommands() []*Command {
	return ns.commands
}

// Add adds a new command to the namespace.
func (ns *NameSpace) Add(c *Command) error {
	ns.commands = append(ns.commands, c)
	c.AddNameSpace(ns.Name)
	return nil
}

// Find finds a given command.
func (ns *NameSpace) Find(c *Command) (int, error) {
	for index, command := range ns.commands {
		if command == c {
			return index, nil
		}
	}
	return -1, errors.New("not found")
}

// DeleteForIndex deletes a command for the given index.
func (ns *NameSpace) DeleteForIndex(index int) (*Command, error) {
	if index >= len(ns.commands) {
		return nil, errors.New("index oout of bounds")
	}
	command := ns.commands[index]
	ns.commands = append(ns.commands[:index], ns.commands[index+1:]...)
	if command.DeleteNameSpace(ns.Name) != nil {
		return command, errors.New("namespace not found in command")
	}
	return command, nil
}

// DeleteForCommand deletes a command for the given command.
func (ns *NameSpace) DeleteForCommand(c *Command) error {
	index, ok := ns.Find(c)
	if ok == nil {
		_, ok = ns.DeleteForIndex(index)
		return ok
	}
	return errors.New("not found")
}
