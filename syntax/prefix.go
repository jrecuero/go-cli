package syntax

// Prefix represents any CLI Prefix information.
type Prefix struct {
	Label     string
	Type      string
	Default   interface{}
	Help      string
	Completer ICompleter
}

// GetLabel returns user Prefix label.
func (p *Prefix) GetLabel() string {
	return p.Label
}

// GetType returns user Prefix type.
func (p *Prefix) GetType() string {
	return p.Type
}

// GetDefault returns user Prefix default value.
func (p *Prefix) GetDefault() interface{} {
	return p.Default
}

// GetHelp returns user Prefix help.
func (p *Prefix) GetHelp() string {
	return p.Help
}

// GetCompleter returns user command completer.
func (p *Prefix) GetCompleter() ICompleter {
	return p.Completer
}

// ToString returns the string with the content information.
func (p *Prefix) ToString() string {
	return p.Label
}

// Setup initializes all Prefix fields.
func (p *Prefix) Setup() error {
	if p.Completer == nil {
		p.Completer = NewCompleterAny(p)
	} else {
		p.Completer.Setup(p)
	}
	return nil
}

var _ IContent = (*Prefix)(nil)
