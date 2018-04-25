package commands

import (
	"fmt"

	"github.com/jrecuero/go-cli/syntax"
)

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
		Syntax: "user name age",
		Help:   "User command",
		Arguments: []syntax.Argument{
			{
				Label:   "name",
				Type:    "string",
				Default: "",
				Help:    "Name information",
			},
			{
				Label:   "age",
				Type:    "int",
				Default: 0,
				Help:    "Age information",
			},
		},
	},
}

// ManagerCmd command
var ManagerCmd = &User{
	syntax.Command{
		Syntax: "manager name",
		Help:   "Manager command",
		Arguments: []syntax.Argument{
			{
				Label:   "name",
				Type:    "string",
				Default: "",
				Help:    "Name information",
			},
		},
	},
}

//func groupCb(context interface{}, arguments interface{}) bool {
//    return true
//}

//// Group command
//var Group = syntax.Command{
//    Syntax: "group name",
//    Cb:     groupCb,
//    Arguments: []syntax.Argument{
//        {
//            Label:   "name",
//            Type:    "string",
//            Default: "",
//            Help:    "Name information",
//        },
//    },
//    Label: "group",
//    Help:  "Group information",
//}
