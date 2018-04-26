package syntax

const _cr = "<<<_CR_>>>"

// GetLabelFromContent returns the label for any content interface.
func GetLabelFromContent(content IContent) string {
	if content == nil {
		return "<nil>"
	}
	return content.GetLabel()
}

// Content represents any  content.
type Content struct {
	label     string
	help      string
	completer ICompleter
}

var _ IContent = (*Content)(nil)

// GetLabel returns content label.
func (c *Content) GetLabel() string {
	return c.label
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

// ToString returns the string with the content information.
func (c *Content) ToString() string {
	return c.label
}

// NewContent returns a new Content instance.
func NewContent(label string, help string, completer ICompleter) IContent {
	return &Content{
		label:     label,
		help:      help,
		completer: completer,
	}
}

// ContentJoint represents any joint content.
type ContentJoint struct {
	*Content
}

var _ IContent = (*ContentJoint)(nil)

// NewContentJoint returns a new ContentJoint instance.
func NewContentJoint(label string, help string, completer ICompleter) *ContentJoint {
	return &ContentJoint{
		NewContent(label, help, completer).(*Content),
	}
}

// CR represents the carrier return content.
var CR *ContentJoint

// GetCR returns CR variable.
func GetCR() *ContentJoint {
	if CR == nil {
		CR = &ContentJoint{
			NewContent(_cr, "Carrier return", nil).(*Content),
		}
		CR.completer = &CompleterJoint{
			&Completer{
				label:   _sink,
				content: CR,
			},
		}
	}
	return CR
}
