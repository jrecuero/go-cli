package syntax

import (
	"errors"
	"strings"
)

// NSManager represents the namespace manager.
// It is in charge to manage the active namespace, it will provides operation
// to settle all namespace commands properly and to find commands based on the
// command line input.
type NSManager struct {
	nsname   string
	ns       *NameSpace
	commands map[string]interface{}
	ctx      *Context
}

// GetNameSpace returns the namespace related with the namager.
func (m *NSManager) GetNameSpace() *NameSpace {
	return m.ns
}

// GetCommands returns the manager commands.
func (m *NSManager) GetCommands() map[string]interface{} {
	return m.commands
}

// NewNSManager creates a new NSManager instance.
func NewNSManager(namespace *NameSpace) *NSManager {
	if namespace == nil {
		return &NSManager{}
	}
	return &NSManager{
		nsname: namespace.Name,
		ns:     namespace,
	}
}

// add inserts a new command in the internal command map.
func (m *NSManager) add(table map[string]interface{}, name []string, value interface{}) error {
	label := name[0]
	if table[label] == nil {
		if len(name) == 1 {
			table[label] = value
		} else {
			table[label] = make(map[string]interface{})
			m.add(table[label].(map[string]interface{}), name[1:], value)
		}
	} else {
		switch v := table[label].(type) {
		case map[string]interface{}:
			if len(name) == 1 {
				return errors.New("too short")
			}
			m.add(v, name[1:], value)
		default:
			return errors.New("error")
		}
	}
	return nil
}

// Setup initializes the namespace manager.
func (m *NSManager) Setup() error {
	if m.ns == nil {
		return errors.New("no namespace")
	}
	m.commands = make(map[string]interface{})
	for _, cmd := range m.ns.GetCommands() {
		cmdSeq := strings.Split(cmd.FullCmd, " ")
		err := m.add(m.commands, cmdSeq, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search searches for the given pattern in the commands map.
func (m *NSManager) Search(pattern string) ([]*Command, error) {
	var commands []*Command
	sequence := strings.Split(pattern, " ")
	if len(sequence) == 0 {
		return nil, errors.New("no pattern")
	}
	return commands, nil
}
