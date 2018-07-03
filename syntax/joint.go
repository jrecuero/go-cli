package syntax

const _cr = "<<<_CR_>>>"

// ContentJoint represents any joint content.
type ContentJoint struct {
	*Content
}

// IsJoint returns if content is a joint.
func (c *ContentJoint) IsJoint() bool {
	return true
}

// GetStrType returns the short string for the content type.
func (c *ContentJoint) GetStrType() string {
	return "J"
}

var _ IContent = (*ContentJoint)(nil)

// NewContentJoint returns a new ContentJoint instance.
func NewContentJoint(label string, help string, completer ICompleter) *ContentJoint {
	content := &ContentJoint{
		NewContent(label, help, completer).(*Content),
	}
	content.matchable = false
	return content
}

// CR represents the carrier return content.
var CR *ContentJoint

// GetCR returns CR variable.
func GetCR() *ContentJoint {
	if CR == nil {
		completer := NewCompleterSink()
		CR = &ContentJoint{
			NewContent(_cr, "Carrier return", completer).(*Content),
		}
		CR.matchable = true
	}
	return CR
}
