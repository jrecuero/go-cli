package commands

import (
	"fmt"

	"github.com/jrecuero/go-cli/syntax"
)

func userCb(context interface{}, arguments interface{}) bool {
	//argos := arguments.(map[string]interface{})
	//name := argos["name"].(string)
	//age := argos["age"].(int)
	name := syntax.GetValueFromArguments("name", arguments).(string)
	age := syntax.GetValueFromArguments("age", arguments).(int)
	fmt.Printf("name: %s age: %d", name, age)
	return true
}

// User command
var User = syntax.UCommand{
	Syntax: "user name age",
	Cb:     userCb,
	Arguments: []syntax.UArgument{
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
	Label: "user",
	Help:  "User command",
}

func groupCb(context interface{}, arguments interface{}) bool {
	return true
}

// Group command
var Group = syntax.UCommand{
	Syntax: "group name",
	Cb:     groupCb,
	Arguments: []syntax.UArgument{
		{
			Label:   "name",
			Type:    "string",
			Default: "",
			Help:    "Name information",
		},
	},
	Label: "group",
	Help:  "Group information",
}
