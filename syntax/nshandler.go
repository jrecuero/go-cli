package syntax

import (
	"errors"
	"fmt"
)

// NSActive represents the active namespace.
type NSActive struct {
	Name    string     // active namespace name.
	NS      *NameSpace // active namespace instance.
	NSMgr   *NSManager // acitve namespace manager
	enabled bool       // is active namespace enabled?
}

// Activate activates a namespace
func (nsa *NSActive) Activate(nsaname string, ns *NameSpace, nsmgr *NSManager) error {
	nsa.Name = nsaname
	nsa.NS = ns
	nsa.NSMgr = nsmgr
	nsa.enabled = true
	return nil
}

// Deactivate deactivates a namespace
func (nsa *NSActive) Deactivate() error {
	nsa.Name = ""
	nsa.NS = nil
	nsa.NSMgr = nil
	nsa.enabled = false
	return nil
}

// NewNSActive creates a new NSActive instance.
func NewNSActive(nsaname string, ns *NameSpace) (*NSActive, error) {
	var err error
	nsa := &NSActive{}
	if nsm, err := CreateNSManager(ns); err == nil {
		if err = nsa.Activate(nsaname, ns, nsm); err != nil {
			return nil, err
		}
		return nsa, nil
	}
	return nil, err
}

// NSHandler represents the namespace handler.
// It is charge to handler all namespaces available in the application, it will
// active one or other namespace, switch between them and control how to switch
// back to the previous namespace properly.
type NSHandler struct {
	active     *NSActive             // active namespace instance.
	namespaces map[string]*NameSpace // map with all available namespaces..
	stack      []*NSActive           // stack with used acitve namespaces.
}

// GetNameSpaces returns all namespaces created in the handler.
func (nsh *NSHandler) GetNameSpaces() map[string]*NameSpace {
	return nsh.namespaces
}

// GetActive returns the active namespace.
func (nsh *NSHandler) GetActive() *NSActive {
	return nsh.active
}

// GetStack returns the namespace stack.
func (nsh *NSHandler) GetStack() []*NSActive {
	return nsh.stack
}

// FindNameSpace looks for a namespace for the given name.
func (nsh *NSHandler) FindNameSpace(nsname string) (*NameSpace, error) {
	if ns, ok := nsh.namespaces[nsname]; ok {
		return ns, nil
	}
	return nil, fmt.Errorf("NameSpace %#v not found", nsname)
}

// CreateNameSpace creates a new namespace.
func (nsh *NSHandler) CreateNameSpace(nsname string) (*NameSpace, error) {
	if ns, ok := nsh.namespaces[nsname]; !ok {
		ns = NewNameSpace(nsname)
		nsh.namespaces[nsname] = ns
		return ns, nil
	}
	return nil, fmt.Errorf("NameSpace %#v not found", nsname)
}

// DeleteNameSpace deletes a namespace.
func (nsh *NSHandler) DeleteNameSpace(nsname string) (*NameSpace, error) {
	if ns, ok := nsh.namespaces[nsname]; ok {
		delete(nsh.namespaces, nsname)
		return ns, nil
	}
	return nil, fmt.Errorf("NameSpace %#v not found", nsname)
}

// RegisterCommandsToNameSpace registers a list of commands to the given namespace.
func (nsh *NSHandler) RegisterCommandsToNameSpace(nsname string, commands []*Command) error {
	if ns, ok := nsh.namespaces[nsname]; ok {
		for _, cmd := range commands {
			if err := ns.Add(cmd); err != nil {
				return fmt.Errorf("Register command %#v in NameSpace %#v failed with %#v", cmd.GetLabel(), nsname, err)
			}
		}
		return nil
	}
	return fmt.Errorf("NameSpace %#v not found", nsname)
}

// RegisterCommandToNameSpace registers a command to the given namespace.
func (nsh *NSHandler) RegisterCommandToNameSpace(nsname string, c *Command) error {
	if ns, ok := nsh.namespaces[nsname]; ok {
		return ns.Add(c)
	}
	return fmt.Errorf("NameSpace %#v not found", nsname)
}

// UnregisterCommandFromNameSpace unregisters a command from the given
// namespace.
func (nsh *NSHandler) UnregisterCommandFromNameSpace(nsname string, c *Command) error {
	if ns, ok := nsh.namespaces[nsname]; ok {
		return ns.DeleteForCommand(c)
	}
	return fmt.Errorf("NameSpace %#v not found", nsname)
}

// ActivateNameSpace activates namespace for the given name.
func (nsh *NSHandler) ActivateNameSpace(nsname string) (*NameSpace, error) {
	if ns, ok := nsh.namespaces[nsname]; ok {
		var err error
		if nsh.active, err = NewNSActive(nsname, ns); err == nil {
			return ns, nil
		}
		return nil, err
	}
	return nil, fmt.Errorf("NameSpace %#v not found", nsname)
}

// DeactivateNameSpace deactivates namespace for the given name.
func (nsh *NSHandler) DeactivateNameSpace(nsname string) (*NameSpace, error) {
	if ns, ok := nsh.namespaces[nsname]; ok {
		nsh.active.Deactivate()
		nsh.active = nil
		return ns, nil
	}
	return nil, fmt.Errorf("NameSpace %#v not found", nsname)
}

// SwitchToNameSpace switches the active namespace and store the all one.
func (nsh *NSHandler) SwitchToNameSpace(nsname string) (*NameSpace, *NameSpace, error) {
	var oldActive *NSActive
	if nsh.active != nil {
		oldActive = nsh.active
	}
	if ns, ok := nsh.ActivateNameSpace(nsname); ok == nil {
		nsh.stack = append(nsh.stack, oldActive)
		return oldActive.NS, ns, nil
	}
	return nil, nil, fmt.Errorf("NameSpace %#v not found", nsname)
}

// SwitchBackToNameSpace switches back to the latest namespace.
func (nsh *NSHandler) SwitchBackToNameSpace() (*NameSpace, *NameSpace, error) {
	stackLen := len(nsh.stack)
	if stackLen == 0 {
		return nil, nil, errors.New("empty stack")
	}
	// lastActive contains the last active namespace, that will be the next
	// active namespace now.
	lastActive := nsh.stack[stackLen-1]
	// active contains the actual actual namesapace, that will be removed.
	active := nsh.active
	if ns, ok := nsh.ActivateNameSpace(lastActive.Name); ok == nil && ns == lastActive.NS {
		nsh.stack = nsh.stack[:stackLen-1]
		return active.NS, lastActive.NS, nil
	}
	return nil, nil, fmt.Errorf("NameSpace %#v not found", lastActive.Name)
}

// NewNSHandler creates a new NSHandler instance.
func NewNSHandler() *NSHandler {
	nsh := &NSHandler{}
	nsh.namespaces = make(map[string]*NameSpace)
	return nsh
}

// CreateNSHandler creates a new NSHandler with a new namespace.
func CreateNSHandler(nsname string, commands []*Command) (*NSHandler, error) {
	nsh := NewNSHandler()
	if _, err := nsh.CreateNameSpace(nsname); err != nil {
		return nil, err
	}
	// Add builtins commands
	commandsToProcess := []*Command{NewExitCommand()}
	for _, cmd := range commands {
		commandsToProcess = append(commandsToProcess, cmd)
	}
	if err := nsh.RegisterCommandsToNameSpace(nsname, commandsToProcess); err != nil {
		return nil, err
	}
	if _, err := nsh.ActivateNameSpace(nsname); err != nil {
		return nil, err
	}
	return nsh, nil
}
