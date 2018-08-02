package commands

import (
	"fmt"

	"github.com/jrecuero/go-cli/dbase"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// Dbase represents the active dataabase
var Dbase *dbase.DataBase

// SetupCommands configures all command to run.
func SetupCommands() []*syntax.Command {
	createCommand := syntax.NewCommand(
		nil,
		"CREATE",
		"Create ...",
		nil,
		nil)

	updateCommand := syntax.NewCommand(
		nil,
		"UPDATE",
		"Update ...",
		nil,
		nil)

	deleteCommand := syntax.NewCommand(
		nil,
		"DELETE",
		"Delete ...",
		nil,
		nil)

	displayCommand := syntax.NewCommand(
		nil,
		"DISPLAY",
		"Display ...",
		nil,
		nil)

	saveCommand := syntax.NewCommand(
		nil,
		"SAVE",
		"Save database...",
		nil,
		nil)
	saveCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("saving database %#v...\n", fmt.Sprintf("%s.db", Dbase.Name))
		return Dbase.Save()
	}

	loadCommand := syntax.NewCommand(
		nil,
		"LOAD filename",
		"Load database...",
		[]*syntax.Argument{
			syntax.NewArgument(
				"filename",
				"Database filename",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	loadCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		filename := params["filename"].(string)
		var err error
		tools.ToDisplay("loading database %#v...\n", filename)
		Dbase, err = dbase.Load(filename)
		return err
	}

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
		"Create row in table",
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

	updateRowCommand := syntax.NewCommand(
		updateCommand,
		"ROW tbname irow [row]@",
		"Update row in table",
		[]*syntax.Argument{
			syntax.NewArgument(
				"tbname",
				"Table name",
				nil,
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"irow",
				"Row index",
				nil,
				"int",
				"0",
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
	updateRowCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tbname := params["tbname"].(string)
		irow, _ := params["irow"].(int)
		row := params["row"].([]interface{})
		tb := Dbase.GetTable(tbname)
		tb.UpdateRow(irow, dbase.NewRow(row...))
		return nil
	}

	deleteRowCommand := syntax.NewCommand(
		deleteCommand,
		"ROW tbname irow",
		"Delete row from table",
		[]*syntax.Argument{
			syntax.NewArgument(
				"tbname",
				"Table name",
				nil,
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"irow",
				"Row index",
				nil,
				"int",
				"0",
				nil),
		},
		nil)
	deleteRowCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tbname := params["tbname"].(string)
		irow := params["irow"].(int)
		tb := Dbase.GetTable(tbname)
		tb.DeleteRow(irow)
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
		layout := tb.Layout
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
		updateCommand,
		updateRowCommand,
		deleteCommand,
		deleteRowCommand,
		displayCommand,
		displayTableCommand,
		saveCommand,
		loadCommand,
	}
	return commands
}
