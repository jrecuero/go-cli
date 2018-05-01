package syntax

import (
	"errors"
)

// NSActive represents the active namespace.
type NSActive struct {
	Name    string
	NS      *NameSpace
	Mgr     *NSManager
	enabled bool
}

// NewNSActive creates a new NSActive instance.
func NewNSActive(name string, ns *NameSpace) *NSActive {
	a := &NSActive{}
	a.Activate(name, ns, NewNSManager(ns))
	return a
}

// Activate activates a namespace
func (a *NSActive) Activate(name string, ns *NameSpace, mgr *NSManager) error {
	a.Name = name
	a.NS = ns
	a.Mgr = mgr
	a.enabled = true
	return nil
}

// Deactivate deactivates a namespace
func (a *NSActive) Deactivate() error {
	a.Name = ""
	a.NS = nil
	a.Mgr = nil
	a.enabled = false
	return nil
}

// NSHandler represents the namespace handler.
// It is charge to handler all namespaces available in the application, it will
// active one or other namespace, switch between them and control how to switch
// back to the previous namespace properly.
type NSHandler struct {
	active     *NSActive
	namespaces map[string]*NameSpace
	stack      []*NSActive
}

// NewNSHandler creates a new NSHandler instance.
func NewNSHandler() *NSHandler {
	h := &NSHandler{}
	h.namespaces = make(map[string]*NameSpace)
	return h
}

// GetNameSpaces returns all namespaces created in the handler.
func (h *NSHandler) GetNameSpaces() map[string]*NameSpace {
	return h.namespaces
}

// GetActive returns the active namespace.
func (h *NSHandler) GetActive() *NSActive {
	return h.active
}

// GetStack returns the namespace stack.
func (h *NSHandler) GetStack() []*NSActive {
	return h.stack
}

// FindNameSpace looks for a namespace for the given name.
func (h *NSHandler) FindNameSpace(name string) (*NameSpace, error) {
	ns, ok := h.namespaces[name]
	if ok {
		return ns, nil
	}
	return nil, errors.New("not found")
}

// CreateNameSpace creates a new namespace.
func (h *NSHandler) CreateNameSpace(name string) (*NameSpace, error) {
	ns, ok := h.namespaces[name]
	if ok {
		return nil, errors.New("found")
	}
	ns = &NameSpace{
		Name: name,
	}
	h.namespaces[name] = ns
	return ns, nil
}

// DeleteNameSpace deletes a namespace.
func (h *NSHandler) DeleteNameSpace(name string) (*NameSpace, error) {
	ns, ok := h.namespaces[name]
	if ok {
		delete(h.namespaces, name)
		return ns, nil
	}
	return nil, errors.New("not found")
}

// RegisterCommandToNameSpace registers a command to the given namespace.
func (h *NSHandler) RegisterCommandToNameSpace(name string, c *Command) error {
	ns, ok := h.namespaces[name]
	if ok {
		return ns.Add(c)
	}
	return errors.New("not found")
}

// UnregisterCommandFromNameSpace unregisters a command from the given
// namespace.
func (h *NSHandler) UnregisterCommandFromNameSpace(name string, c *Command) error {
	ns, ok := h.namespaces[name]
	if ok {
		return ns.DeleteForCommand(c)
	}
	return errors.New("Not found")
}

// ActivateNameSpace activates namespace for the given name.
func (h *NSHandler) ActivateNameSpace(name string) (*NameSpace, error) {
	ns, ok := h.namespaces[name]
	if ok {
		h.active = NewNSActive(name, ns)
		return ns, nil
	}
	return nil, errors.New("Not found")
}

// DeactivateNameSpace deactivates namespace for the given name.
func (h *NSHandler) DeactivateNameSpace(name string) (*NameSpace, error) {
	ns, ok := h.namespaces[name]
	if ok {
		h.active.Deactivate()
		h.active = nil
		return ns, nil
	}
	return nil, errors.New("Not found")
}

// SwitchToNameSpace switches the active namespace and store the all one.
func (h *NSHandler) SwitchToNameSpace(name string) (*NameSpace, *NameSpace, error) {
	var oldActive *NSActive
	if h.active != nil {
		oldActive = h.active
	}
	ns, ok := h.ActivateNameSpace(name)
	if ok == nil {
		h.stack = append(h.stack, oldActive)
		return oldActive.NS, ns, nil
	}
	return nil, nil, errors.New("not found")
}

// SwitchBackToNameSpace switches back to the latest namespace.
func (h *NSHandler) SwitchBackToNameSpace() (*NameSpace, *NameSpace, error) {
	stackLen := len(h.stack)
	if stackLen == 0 {
		return nil, nil, errors.New("empty stack")
	}
	// lastActive contains the last active namespace, that will be the next
	// active namespace now.
	lastActive := h.stack[stackLen-1]
	// active contains the actual actual namesapace, that will be removed.
	active := h.active
	ns, ok := h.ActivateNameSpace(lastActive.Name)
	if ok == nil && ns == lastActive.NS {
		h.stack = h.stack[:stackLen-1]
		return active.NS, lastActive.NS, nil
	}
	return nil, nil, errors.New("not found")
}
