package commands

import (
	"github.com/jrecuero/go-cli/dbase"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// Database represents the active dataabase
var Database *dbase.DB

// SetupCommands configures all command to run.
func SetupCommands() []*syntax.Command {
	createCommand := syntax.NewCommand(nil, "CREATE", "Create ...", nil, nil)

	createDbaseCommand := syntax.NewCommand(createCommand, "DBASE dbname", "Create Database",
		[]*syntax.Argument{
			syntax.NewArgument("dbname", "Database name", nil, "string", "none", nil),
		}, nil)
	createDbaseCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		dbname := params["dbname"].(string)
		tools.ToDisplay("CREATE DBASE %#v\n", dbname)
		Database = dbase.NewDB(dbname)
		return nil
	}

	createTableCommand := syntax.NewCommand(createCommand, "TABLE tbname", "Create table",
		[]*syntax.Argument{
			syntax.NewArgument("tbname", "Table name", nil, "string", "none", nil),
		}, nil)
	createTableCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tbname := params["tbname"].(string)
		tools.ToDisplay("CREATE TABLE %#v\n", tbname)
		Database.CreateTable(tbname, dbase.PersonLayout())
		return nil
	}

	commands := []*syntax.Command{
		createCommand,
		createDbaseCommand,
		createTableCommand,
	}
	return commands
}
