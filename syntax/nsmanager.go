package syntax

// NSManager represents the namespace manager.
// It is in charge to manage the active namespace, it will provides operation
// to settle all namespace commands properly and to find commands based on the
// command line input.
type NSManager struct {
	nsname   string
	ns       *NameSpace
	commands []interface{}
	ctx      *Context
}

// NewNSManager creates a new NSManager instance.
func NewNSManager(name string, namespace *NameSpace) *NSManager {
	mgr := &NSManager{
		nsname: name,
		ns:     namespace,
	}
	return mgr
}
