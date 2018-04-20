package syntax

const _cr = "<<<_CR_>>>"

// AContent represents any  content.
type AContent struct {
	label     string
	help      string
	completer ICompleter
}

// NewAContent returns a new AContent instance.
func NewAContent(label string, completer ICompleter) IContent {
	return &AContent{
		label:     label,
		help:      "Joint content",
		completer: completer,
	}
}

// AJoint represents any joint content.
type AJoint struct {
	*AContent
}

// GetLabel returns joint content label.
func (j *AJoint) GetLabel() string {
	return j.label
}

// GetType returns joint content type.
func (j *AJoint) GetType() string {
	return "string"
}

// GetDefault returns joint content default value.
func (j *AJoint) GetDefault() interface{} {
	return j.label
}

// GetHelp return joint content help.
func (j *AJoint) GetHelp() string {
	return j.help
}

// GetCompleter returns user command completer.
func (j *AJoint) GetCompleter() ICompleter {
	return j.completer
}

// NewAJoint returns a new AJoint instance.
func NewAJoint(label string, completer ICompleter) *AJoint {
	return &AJoint{
		&AContent{
			label:     label,
			help:      "Joint content",
			completer: completer,
		},
	}
}

// CR represents the carrier return content.
var CR *AJoint

// GetCR returns CR variable.
func GetCR() *AJoint {
	if CR == nil {
		CR = &AJoint{
			&AContent{
				label: _cr,
				help:  "Carrier return",
			},
		}
		CR.completer = &CJoint{
			&Completer{
				label:   _sink,
				Content: CR,
			},
		}
	}
	return CR
}
