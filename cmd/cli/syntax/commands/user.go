package commands

import (
	"fmt"

	"github.com/jrecuero/go-cli/syntax"
)

func userCb(context interface{}, arguments interface{}) bool {
	argos := arguments.(map[string]interface{})
	name := argos["name"].(string)
	age := argos["age"].(int)
	fmt.Printf("name: %s age: %d", name, age)
	return true
}

// User command
var User = syntax.Ucommand{
	Syntax: "user name age",
	Cb:     userCb,
	Arguments: []syntax.Uargument{
		{
			Name:    "name",
			Label:   "name",
			Type:    "string",
			Default: "",
			Help:    "Name information",
		},
		{
			Name:    "age",
			Label:   "age",
			Type:    "int",
			Default: 0,
			Help:    "Age information",
		},
	},
	Name: "user",
	Help: "User command",
}

func groupCb(context interface{}, arguments interface{}) bool {
	return true
}

// Group command
var Group = syntax.Ucommand{
	Syntax: "group name",
	Cb:     groupCb,
	Arguments: []syntax.Uargument{
		{
			Name:    "name",
			Label:   "name",
			Type:    "string",
			Default: "",
			Help:    "Name information",
		},
	},
	Name: "group",
	Help: "Group information",
}
