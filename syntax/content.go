package syntax

// Content represents any  content.
type Content struct {
	label     string
	help      string
	completer ICompleter
	matchable bool
}

var _ IContent = (*Content)(nil)

// GetLabel returns content label.
func (c *Content) GetLabel() string {
	return c.label
}

// GetString returns the string for the content node.
func (c *Content) GetString() string {
	return c.GetLabel()
}

// GetType returns content type.
func (c *Content) GetType() string {
	return "string"
}

// GetDefault returns content default value.
func (c *Content) GetDefault() interface{} {
	return c.label
}

// GetHelp return content help.
func (c *Content) GetHelp() string {
	return c.help
}

// GetCompleter returns user command completer.
func (c *Content) GetCompleter() ICompleter {
	return c.completer
}

// Validate returns user command completer.
func (c *Content) Validate(val interface{}) (bool, error) {
	return true, nil
}

// ToString returns the string with the content information.
func (c *Content) ToString() string {
	return c.label
}

// IsCommand returns if content is a command.
func (c *Content) IsCommand() bool {
	return false
}

// IsMode returns if content is a mode.
func (c *Content) IsMode() bool {
	return false
}

// IsArgument returns if content is a argument.
func (c *Content) IsArgument() bool {
	return false
}

// IsJoint returns if content is a joint.
func (c *Content) IsJoint() bool {
	return false
}

// IsMatchable returns if content is matchable.
func (c *Content) IsMatchable() bool {
	return c.matchable
}

// GetStrType returns the short string for the content type.
func (c *Content) GetStrType() string {
	return "X"
}

// SetCompleter sets a new completer for the content.
func (c *Content) SetCompleter(completer ICompleter) bool {
	c.completer = completer
	return true
}

// GetGraphPattern returns the string with the graphical pattern.
func (c *Content) GetGraphPattern() *string {
	return nil
}

// NewContent returns a new Content instance.
func NewContent(label string, help string, completer ICompleter) IContent {
	return &Content{
		label:     label,
		help:      help,
		completer: completer,
	}
}

// GetLabelFromContent returns the label for any content interface.
func GetLabelFromContent(content IContent) string {
	if content == nil {
		return "<nil>"
	}
	return content.GetLabel()
}
