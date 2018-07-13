package syntax

import (
	"fmt"
	"strconv"
)

// CastingCall represents the function for casting the argument type.
type CastingCall func(val string) (interface{}, error)

// CastInt performs a string to int casting.
func CastInt(val string) (interface{}, error) {
	result, _ := strconv.Atoi(val)
	return result, nil
}

// CastString performs a string to string casting.
func CastString(val string) (interface{}, error) {
	return val, nil
}

var castingMap = map[string]CastingCall{
	"string": CastString,
	"int":    CastInt,
}

// Argument represents any CLI argument information.
type Argument struct {
	*Content
	Type    string
	Casting CastingCall
	Default interface{}
}

// Setup initializes all argument fields.
func (a *Argument) Setup() *Argument {
	if a.completer == nil {
		a.completer = NewCompleterArgument(a.GetLabel())
	}
	if a.Casting == nil {
		castingCall, ok := castingMap[a.Type]
		if ok {
			a.Casting = castingCall
		} else {
			panic(fmt.Sprintf("argument type %#v does not have casting call", a.Type))
		}
	}
	return a
}

// GetType returns content type.
func (a *Argument) GetType() string {
	return a.Type
}

// Cast returns the casting for the argument type.
func (a *Argument) Cast(val string) (interface{}, error) {
	if a.Casting != nil {
		return a.Casting(val)
	}
	return nil, fmt.Errorf("casting call not found for type %#v", a.Type)
}

// CreateKeywordFromSelf creates a new Argument instance that contains the
// argument in keyword format
func (a *Argument) CreateKeywordFromSelf() *Argument {
	label := fmt.Sprintf("-%s", a.label)
	return &Argument{
		Content: NewContent(label, a.help, NewCompleterIdent(label)).(*Content),
		Type:    "string",
		Default: label,
	}
}

// IsArgument returns if content is an argument..
func (a *Argument) IsArgument() bool {
	return true
}

// GetStrType returns the short string for the content type.
func (a *Argument) GetStrType() string {
	return "A"
}

var _ IContent = (*Argument)(nil)

// GetValueFromArguments returns the value for the given field in arguments
// passed to.
func GetValueFromArguments(field string, arguments interface{}) interface{} {
	argos := arguments.(map[string]interface{})
	value := argos[field]
	return value
}

// NewArgument creates a new Argument instance.
func NewArgument(label string, help string, completer ICompleter, atype string, adefault interface{}, casting CastingCall) *Argument {
	argo := &Argument{
		Content: NewContent(label, help, completer).(*Content),
		Type:    atype,
		Default: adefault,
		Casting: casting,
	}
	argo.Setup()
	return argo
}
