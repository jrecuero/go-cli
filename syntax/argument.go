package syntax

import (
	"fmt"
	"strconv"
)

// CastingCall represents the function for casting the argument type.
type CastingCall func(val string) (interface{}, error)

// ValidateCall represents the function for validating the argument value.
type ValidateCall func(val interface{}) (bool, error)

// CastInt performs a string to int casting.
func CastInt(val string) (interface{}, error) {
	result, err := strconv.Atoi(val)
	return result, err
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
	Type      string
	Caster    CastingCall
	Validator ValidateCall
	Default   interface{}
}

// Setup initializes all argument fields.
func (a *Argument) Setup() *Argument {
	if a.completer == nil {
		a.completer = NewCompleterArgument(a.GetLabel())
	}
	if a.Caster == nil {
		castingCall, ok := castingMap[a.Type]
		if ok {
			a.Caster = castingCall
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
	if a.Caster != nil {
		return a.Caster(val)
	}
	return val, nil
}

// Validate checks in the argument value is a valid one.
func (a *Argument) Validate(val interface{}) (bool, error) {
	if a.Validator != nil {
		return a.Validator(val)
	}
	return true, nil
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
		Content:   NewContent(label, help, completer).(*Content),
		Type:      atype,
		Default:   adefault,
		Caster:    casting,
		Validator: nil,
	}
	argo.Setup()
	return argo
}
