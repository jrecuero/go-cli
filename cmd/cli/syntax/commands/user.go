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
		Content: syntax.NewContent("", "User command", nil).(*syntax.Content),
		Syntax:  "user name [ age | id ] ?",
		//Help:   "User command",
		Arguments: []*syntax.Argument{
			{
				Content: syntax.NewContent("name", "", nil).(*syntax.Content),
				//Label:     "name",
				Type:    "string",
				Default: "",
				//Help:      "Name information",
				//Completer: &UserCompleter{syntax.NewCompleterIdent(nil)},
			},
			{
				Content: syntax.NewContent("age", "Age information", nil).(*syntax.Content),
				//Label:   "age",
				Type:    "int",
				Default: 0,
				//Help:    "Age information",
			},
			{
				Content: syntax.NewContent("id", "ID information", nil).(*syntax.Content),
				//Label:   "id",
				Type:    "int",
				Default: 0,
				//Help:    "ID information",
			},
		},
	},
}

// ManagerCmd command
var ManagerCmd = &User{
	syntax.Command{
		Content: syntax.NewContent("", "Manager command", nil).(*syntax.Content),
		Syntax:  "manager name",
		//Help:   "Manager command",
		Arguments: []*syntax.Argument{
			{
				Content: syntax.NewContent("name", "", nil).(*syntax.Content),
				//Label:   "name",
				Type:    "string",
				Default: "",
				//Help:    "Name information",
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
		Content: syntax.NewContent("", "Set command", nil).(*syntax.Content),
		Syntax:  "set speed value",
		//Help:   "Set command",
		Prefixes: []*syntax.Prefix{
			//{
			//    Label:   "speed",
			//    Type:    "string",
			//    Default: "speed",
			//    Help:    "Set the speed",
			//},
			{
				Content: syntax.NewContent("speed", "Set the speed", nil).(*syntax.Content),
				Type:    "string",
				Default: "speed",
			},
		},
		Arguments: []*syntax.Argument{
			{
				Content: syntax.NewContent("value", "Speed value", nil).(*syntax.Content),
				//Label:   "value",
				Type:    "int",
				Default: 0,
				//Help: "Speed value",
			},
		},
	},
}
