package syntax

const _cr = "<<<_CR_>>>"

// cr represents the carrier return content.
type cr struct{}

func (c *cr) GetLabel() string {
	return _cr
}

func (c *cr) GetName() string {
	return ""
}

func (c *cr) GetType() string {
	return "string"
}

func (c *cr) GetDefault() interface{} {
	return _cr
}

func (c *cr) GetHelp() string {
	return "Carrier return"
}

// CR represents the carrier return content.
var CR = &cr{}
