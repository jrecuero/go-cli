package syntax

import (
	"fmt"
	"strconv"
	"strings"
)

// CastingCall represents the function for casting the argument type.
type CastingCall func(val string) (interface{}, error)

// AssignerCall represents the function for casting the argument type.
type AssignerCall func(argcontent interface{}, val interface{}) interface{}

// CastInt performs a string to int casting.
func CastInt(val string) (interface{}, error) {
	result, err := strconv.Atoi(val)
	return result, err
}

// CastString performs a string to string casting.
func CastString(val string) (interface{}, error) {
	return val, nil
}

// AssignInt performs a value assignment.
func AssignInt(argcontent interface{}, val interface{}) interface{} {
	return val
}

// AssignString performs a value assignment.
func AssignString(argcontent interface{}, val interface{}) interface{} {
	return val
}

// AssignStringList performs a value assignment.
func AssignStringList(argcontent interface{}, val interface{}) interface{} {
	if argcontent == nil {
		return []string{val.(string)}
	}
	return append(argcontent.([]string), val.(string))
}

// AssignFreeForm performs a value assignment.
func AssignFreeForm(argcontent interface{}, val interface{}) interface{} {
	if argcontent == nil {
		return val.(string)
	}
	return argcontent.(string) + " " + val.(string)
}

// AssignMap performs a value assignment.
func AssignMap(argcontent interface{}, val interface{}) interface{} {
	entry := strings.Split(val.(string), ":")
	mapEntry := map[string]string{entry[0]: entry[1]}
	if argcontent == nil {
		return mapEntry
	}
	argcontent.(map[string]string)[entry[0]] = entry[1]
	return argcontent
}

// argmapper represents the structure with all mapper functions.
type argmapper struct {
	casting  CastingCall
	assigner AssignerCall
}

var callerMap = map[string]argmapper{
	"string":   argmapper{CastString, AssignString},
	"int":      argmapper{CastInt, AssignInt},
	"list":     argmapper{CastString, AssignStringList},
	"freeform": argmapper{CastString, AssignFreeForm},
	"map":      argmapper{CastString, AssignMap},
}

// Argument represents any CLI argument information.
type Argument struct {
	*Content              // argument content
	Type     string       // identifies the type of argument.
	Caster   CastingCall  // caster method to obtein proper argument value.
	Assigner AssignerCall // assigner will copy the data passed.
	Default  interface{}  // default argument value
}

var argumentGraphPattern = ">%s]"

// Setup initializes all argument fields.
func (arg *Argument) Setup() *Argument {
	if arg.completer == nil {
		arg.completer = NewCompleterArgument(arg.GetLabel())
	}
	if arg.Caster == nil {
		mapper, ok := callerMap[arg.Type]
		if ok {
			arg.Caster = mapper.casting
		} else {
			panic(fmt.Sprintf("argument type %#v does not have casting call", arg.Type))
		}
	}
	if arg.Assigner == nil {
		mapper, ok := callerMap[arg.Type]
		if ok {
			//tools.Debug("mapping arg: %#v type: %#v assigner: %#v\n", arg.GetLabel(), arg.GetType(), mapper.assigner)
			arg.Assigner = mapper.assigner
		} else {
			panic(fmt.Sprintf("argument type %#v does not have assigner call", arg.Type))
		}
	}
	return arg
}

// GetType returns content type.
func (arg *Argument) GetType() string {
	return arg.Type
}

// GetGraphPattern returns the string with the graphical pattern.
func (arg *Argument) GetGraphPattern() *string {
	return &argumentGraphPattern
}

// Cast returns the casting for the argument type.
func (arg *Argument) Cast(val string) (interface{}, error) {
	if arg.Caster != nil {
		return arg.Caster(val)
	}
	return val, nil
}

// Assign returns the casting for the argument type.
func (arg *Argument) Assign(argcontent interface{}, val interface{}) interface{} {
	//tools.Debug("arg: %#v type: %#v assigner: %#v\n", arg.GetLabel(), arg.GetType(), arg.Assigner)
	if arg.Assigner != nil {
		return arg.Assigner(argcontent, val)
	}
	return nil
}

// CreateKeywordFromSelf creates a new Argument instance that contains the
// argument in keyword format
func (arg *Argument) CreateKeywordFromSelf() *Argument {
	label := fmt.Sprintf("-%s", arg.label)
	return &Argument{
		Content: NewContent(label, arg.help, NewCompleterIdent(label)).(*Content),
		Type:    "string",
		Default: label,
	}
}

// IsArgument returns if content is an argument..
func (arg *Argument) IsArgument() bool {
	return true
}

// GetStrType returns the short string for the content type.
func (arg *Argument) GetStrType() string {
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
		Caster:  casting,
	}
	argo.Setup()
	return argo
}
