package syntax

import "errors"

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
	a.Activate(name, ns, NewNSManager(name, ns))
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
	var oldNS *NameSpace
	if h.active != nil {
		oldNS = h.active.NS
	}
	ns, ok := h.ActivateNameSpace(name)
	if ok == nil {
		return oldNS, ns, nil
	}
	return nil, nil, errors.New("not found")
}

// SwitchBackToNameSpace switches back to the latest namespace.
func (h *NSHandler) SwitchBackToNameSpace() (*NameSpace, *NameSpace, error) {
	return nil, nil, nil
}
