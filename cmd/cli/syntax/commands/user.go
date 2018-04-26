package commands

import (
	"fmt"

	"github.com/jrecuero/go-cli/syntax"
)

// UserCompleter is a custon completer for user arguments.
type UserCompleter struct {
	*syntax.CompleterIdent
}

// User represents a command
type User struct {
	syntax.Command
}

// Enter is the default User, it executes with the running context.
func (u *User) Enter(ctx *syntax.Context, arguments interface{}) error {
	name := syntax.GetValueFromArguments("name", arguments).(string)
	age := syntax.GetValueFromArguments("age", arguments)
	if age == nil {
		fmt.Printf("Your name is %s\n", name)
	} else {
		age = age.(int)
		fmt.Printf("Mr./Mrs. %s is %d old\n", name, age)
	}
	return nil
}

// UserCmd command
var UserCmd = &User{
	syntax.Command{
		Syntax: "user name [ age | id ] ?",
		Help:   "User command",
		Arguments: []*syntax.Argument{
			{
				Label:     "name",
				Type:      "string",
				Default:   "",
				Help:      "Name information",
				Completer: &UserCompleter{syntax.NewCompleterIdent(nil)},
			},
			{
				Label:   "age",
				Type:    "int",
				Default: 0,
				Help:    "Age information",
			},
			{
				Label:   "id",
				Type:    "int",
				Default: 0,
				Help:    "ID information",
			},
		},
	},
}

// ManagerCmd command
var ManagerCmd = &User{
	syntax.Command{
		Syntax: "manager name",
		Help:   "Manager command",
		Arguments: []*syntax.Argument{
			{
				Label:   "name",
				Type:    "string",
				Default: "",
				Help:    "Name information",
			},
		},
	},
}

// Set represents a command
type Set struct {
	syntax.Command
}

// Enter is the default User, it executes with the running context.
func (r *Set) Enter(ctx *syntax.Context, arguments interface{}) error {
	value := syntax.GetValueFromArguments("value", arguments).(int)
	fmt.Printf("Set speed to %d\n", value)
	return nil
}

// SetSpeedCmd command
var SetSpeedCmd = &Set{
	syntax.Command{
		Syntax: "set speed value",
		Help:   "Set command",
		Prefixes: []*syntax.Prefix{
			{
				Label: "speed",
				Type:  "string",
				Help:  "Set the speed",
			},
		},
		Arguments: []*syntax.Argument{
			{
				Label: "value",
				Type:  "int",
				Help:  "Speed value",
			},
		},
	},
}
