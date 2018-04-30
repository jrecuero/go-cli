package syntax

// Prefix represents any CLI Prefix information.
type Prefix struct {
	*Content
	Type    string
	Default interface{}
}

// Setup initializes all Prefix fields.
func (p *Prefix) Setup() error {
	if p.completer == nil {
		p.completer = NewCompleterAny(p)
	} else {
		p.completer.Setup(p)
	}
	return nil
}

var _ IContent = (*Prefix)(nil)
