package syntax

// Callback represents the type for any command callback.
type Callback func(context interface{}, arguments interface{}) bool

// Uargument represents any CLI argument information.
type Uargument struct {
	Name    string
	Label   string
	Type    string
	Default interface{}
	Help    string
}

// Ucommand represents any CLI command defined by the user..
type Ucommand struct {
	Syntax    string
	Cb        Callback
	Arguments []Uargument
	Name      string
	Help      string
}

// NewUcommand returns a new Command instance.
func NewUcommand(name string, syntax string, cb Callback) *Ucommand {
	return &Ucommand{
		Syntax: syntax,
		Cb:     cb,
		Name:   name,
	}
}
