package commands

import (
	"github.com/jrecuero/go-cli/dbase"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// StructuredDatabase represents the active dataabase
//var StructuredDatabase *dbase.DB

// Dbase represents the active dataabase
var Dbase *dbase.DataBase

//// StructuredSetupCommands configures all command to run.
//func StructuredSetupCommands() []*syntax.Command {
//    createCommand := syntax.NewCommand(nil, "CREATE", "Create ...", nil, nil)

//    createDbaseCommand := syntax.NewCommand(createCommand, "DBASE dbname", "Create Database",
//        []*syntax.Argument{
//            syntax.NewArgument("dbname", "Database name", nil, "string", "none", nil),
//        }, nil)
//    createDbaseCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
//        params := arguments.(map[string]interface{})
//        dbname := params["dbname"].(string)
//        tools.ToDisplay("CREATE DBASE %#v\n", dbname)
//        StructuredDatabase = dbase.NewDB(dbname)
//        return nil
//    }

//    createTableCommand := syntax.NewCommand(createCommand, "TABLE tbname", "Create table",
//        []*syntax.Argument{
//            syntax.NewArgument("tbname", "Table name", nil, "string", "none", nil),
//        }, nil)
//    createTableCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
//        params := arguments.(map[string]interface{})
//        tbname := params["tbname"].(string)
//        tools.ToDisplay("CREATE TABLE %#v\n", tbname)
//        Database.CreateTable(tbname, dbase.PersonLayout())
//        return nil
//    }

//    commands := []*syntax.Command{
//        createCommand,
//        createDbaseCommand,
//        createTableCommand,
//    }
//    return commands
//}

// SetupCommands configures all command to run.
func SetupCommands() []*syntax.Command {
	createCommand := syntax.NewCommand(
		nil,
		"CREATE",
		"Create ...",
		nil,
		nil)

	displayCommand := syntax.NewCommand(
		nil,
		"DISPLAY",
		"Display ...",
		nil,
		nil)

	createDbaseCommand := syntax.NewCommand(
		createCommand,
		"DBASE dbname",
		"Create Database",
		[]*syntax.Argument{
			syntax.NewArgument(
				"dbname",
				"Database name",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	createDbaseCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		dbname := params["dbname"].(string)
		tools.ToDisplay("CREATE DBASE %#v\n", dbname)
		Dbase = dbase.NewDataBase(dbname)
		return nil
	}

	createTableCommand := syntax.NewCommand(
		createCommand,
		"TABLE tbname [layout]@",
		"Create table",
		[]*syntax.Argument{
			syntax.NewArgument(
				"tbname",
				"Table name",
				nil,
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"layout",
				"Table layout",
				nil,
				"map",
				"none",
				nil),
		},
		nil)
	createTableCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tbname := params["tbname"].(string)
		layout := params["layout"].(map[string]string)
		tools.ToDisplay("CREATE TABLE %#v %#v\n", tbname, layout)
		tbLayout := dbase.NewTableLayout()
		for k, v := range layout {
			tbLayout.AddColumn(dbase.NewColumn(k, v))
		}
		tb := dbase.NewTable(tbname).SetLayout(tbLayout)
		Dbase.AddTable(tb)
		return nil
	}

	createRowCommand := syntax.NewCommand(
		createCommand,
		"ROW tbname [row]@",
		"Create table",
		[]*syntax.Argument{
			syntax.NewArgument(
				"tbname",
				"Table name",
				nil,
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"row",
				"New row",
				nil,
				"list",
				"none",
				nil),
		},
		nil)
	createRowCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tbname := params["tbname"].(string)
		row := params["row"].([]interface{})
		tb := Dbase.GetTable(tbname)
		tb.AddRow(dbase.NewRow(row...))
		return nil
	}

	displayTableCommand := syntax.NewCommand(
		displayCommand,
		"TABLE tbname",
		"Display table",
		[]*syntax.Argument{
			syntax.NewArgument(
				"tbname",
				"Table name",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	displayTableCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tbname := params["tbname"].(string)
		tb := Dbase.GetTable(tbname)
		layout := tb.GetLayout()
		for i, col := range layout.Columns {
			tools.ToDisplay("%d Column: %#v\n", i, col)
		}
		for _, row := range tb.Rows {
			tools.ToDisplay("Row: %#v\n", row)
		}
		return nil
	}

	commands := []*syntax.Command{
		createCommand,
		createDbaseCommand,
		createTableCommand,
		createRowCommand,
		displayCommand,
		displayTableCommand,
	}
	return commands
}
