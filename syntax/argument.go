package syntax

//// Argument represents any CLI argument information.
//type Argument struct {
//    Label     string
//    Type      string
//    Default   interface{}
//    Help      string
//    Completer ICompleter
//}

// Argument represents any CLI argument information.
type Argument struct {
	*Content
	Type    string
	Default interface{}
}

//// GetLabel returns user argument label.
//func (a *Argument) GetLabel() string {
//    return a.Label
//}

//// GetType returns user argument type.
//func (a *Argument) GetType() string {
//    return a.Type
//}

//// GetDefault returns user argument default value.
//func (a *Argument) GetDefault() interface{} {
//    return a.Default
//}

//// GetHelp returns user argument help.
//func (a *Argument) GetHelp() string {
//    return a.Help
//}

//// GetCompleter returns user command completer.
//func (a *Argument) GetCompleter() ICompleter {
//    return a.Completer
//}

//// ToString returns the string with the content information.
//func (a *Argument) ToString() string {
//    return a.Label
//}

//// IsCommand returns if content is a command.
//func (a *Argument) IsCommand() bool {
//    return false
//}

// GetValueFromArguments returns the value for the given field in arguments
// passed to.
func GetValueFromArguments(field string, arguments interface{}) interface{} {
	argos := arguments.(map[string]interface{})
	value := argos[field]
	return value
}

// Setup initializes all argument fields.
func (a *Argument) Setup() error {
	if a.completer == nil {
		a.completer = NewCompleterAny(a)
	} else {
		a.completer.Setup(a)
	}
	return nil
}

var _ IContent = (*Argument)(nil)
