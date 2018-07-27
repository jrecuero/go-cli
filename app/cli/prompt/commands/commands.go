package commands

import (
	"strconv"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// NamesCompleter represents the name completer
type versionCompleter struct {
	*syntax.CompleterArgument
}

// Validate checks if the content is value for the given line.
func (vc *versionCompleter) Validate(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) bool {
	token := strings.Fields(line.(string))[index]
	if data, err := strconv.Atoi(token); err == nil {
		if data >= 1 && data <= 9 {
			return true
		}
	}
	tools.ToDisplay("\ninvalid version: %s. 1 <= version <= 9\n", token)
	return false
}

// NamesCompleter represents the name completer
type NamesCompleter struct {
	*syntax.CompleterArgument
}

type boolCmdCompleter struct {
	*syntax.CompleterCommand
}

var boolFlag = true

func (bc *boolCmdCompleter) Match(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (int, bool) {
	if boolFlag {
		return bc.CompleterCommand.Match(ctx, content, line, index)
	}
	return index, false
}

func (bc *boolCmdCompleter) Complete(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	if boolFlag {
		return bc.CompleterCommand.Complete(ctx, content, line, index)
	}
	return nil, false
}

func (bc *boolCmdCompleter) Help(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	if boolFlag {
		return bc.CompleterCommand.Help(ctx, content, line, index)
	}
	return nil, false
}

// Query returns the query for any node completer.
func (nc *NamesCompleter) Query(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	data := []*syntax.CompleteHelp{
		syntax.NewCompleteHelp("Coke", "Coke device"),
		syntax.NewCompleteHelp("Pepsi", "Pepsi device"),
	}
	return data, true
}

// Validate checks if the content is value for the given line.
func (nc *NamesCompleter) Validate(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) bool {
	tokens := strings.Fields(line.(string))
	return tokens[index] == "Coke" || tokens[index] == "Pepsi"
}

// SetupCommands configures all command to run.
func SetupCommands() []*syntax.Command {
	tenantCmd := syntax.NewCommand(nil, "tenant tname", "Create a new tenant",
		[]*syntax.Argument{
			syntax.NewArgument("tname", "Tenant name", nil, "string", "none", nil),
		}, nil)
	tenantCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tname := tools.GetValFromArgs(arguments, "tname").(string)
		tools.ToDisplay("Create tenant %#v\n", tname)
		return nil
	}
	tenantCmd.IsBuiltIn = true

	setCmd := syntax.NewCommand(nil, "set version", "Set test help",
		[]*syntax.Argument{
			syntax.NewArgument("version", "Version number", &versionCompleter{syntax.NewCompleterArgument("version")}, "int", 0, nil),
		}, nil)
	setCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		version, err := ctx.GetArgValueForArgLabel(nil, "version")
		if err != nil {
			tools.ToDisplay("Error: %#v\n", err)
		}
		tools.ToDisplay("executing enter with version(ctx): %#v\n", version)
		params := arguments.(map[string]interface{})
		tools.ToDisplay("executing enter wit version(args): %#v\n", params["version"])
		return nil
	}

	getCmd := syntax.NewCommand(nil, "get", "Get test help", nil, nil)
	getCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("executing get\n")
		return nil
	}

	setBoolCmd := syntax.NewCommand(setCmd, "bool", "Set Bool test help", nil, nil)
	setBoolCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("executing bool command\n")
		return nil
	}
	setBoolCmd.SetCompleter(&boolCmdCompleter{syntax.NewCompleterCommand("bool")})

	setBaudrateCmd := syntax.NewCommand(setCmd, "baudrate [speed | parity]?", "Set baudrate help",
		[]*syntax.Argument{
			syntax.NewArgument("speed", "Baudrate speed", nil, "string", "", nil),
			syntax.NewArgument("parity", "Baudrate parity value", nil, "string", "", nil),
		}, nil)
	setBaudrateCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("executing baudrate with speed: %#v parity: %#v\n", params["speed"], params["parity"])
		return nil
	}

	setSpeedCmd := syntax.NewCommand(setCmd, "speed", "Set Speed test help", nil, nil)
	setSpeedCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		if ok, _ := ctx.GetProcess().Check(syntax.RUNASNOFINAL); ok {
			tools.ToDisplay("executing set speed as-no-final!!!!\n")
		} else {
			tools.ToDisplay("executing set speed\n")
		}
		return nil
	}
	setSpeedCmd.RunAsNoFinal = true

	setSpeedDeviceCmd := syntax.NewCommand(setSpeedCmd, "device name", "Set speed device help",
		[]*syntax.Argument{
			syntax.NewArgument("name", "Device name", &NamesCompleter{syntax.NewCompleterArgument("name")}, "string", "", nil),
		}, nil)
	setSpeedDeviceCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("executing set speed device with device: %#v\n", params["name"])
		//if argo, err := setSpeedDeviceCmd.LookForArgument("name"); err == nil {
		//    q, _ := argo.GetCompleter().Query(ctx, nil, nil, 0)
		//    tools.ToDisplay("Query: %#v\n", q)
		//}
		return nil
	}

	getBaudrateCmd := syntax.NewCommand(getCmd, "baudrate", "Get Baudrate test help", nil, nil)
	getBaudrateCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("get baudrate command\n")
		return nil
	}

	getSpeedCmd := syntax.NewCommand(getCmd, "speed [device name | value]?", "Get speed help",
		[]*syntax.Argument{
			syntax.NewArgument("device", "Device", nil, "string", "", nil),
			syntax.NewArgument("name", "Device name", nil, "string", "", nil),
			syntax.NewArgument("value", "Speed value", nil, "string", "", nil),
		}, nil)
	getSpeedCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("executing get speed with device: %#v name: %#v value:%#v\n",
			params["device"], params["name"], params["value"])
		return nil
	}

	configCmd := syntax.NewMode(nil, "config", "Config test help", nil, nil)
	configCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("executing config mode\n")
		if _prompt, err := ctx.Cache.Get("prompt"); err == nil {
			prompt.OptionInputTextColor(prompt.Red)(_prompt.(*prompt.Prompt))
		}
		return nil
	}
	configCmd.Prompt = "config>>> "

	configDescCmd := syntax.NewCommand(configCmd, "desc", "Description config test help", nil, nil)
	configDescCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("config description\n")
		return nil
	}

	terminalCmd := syntax.NewCommand(nil, "terminal [device | remote]!", "Configure terminal",
		[]*syntax.Argument{
			syntax.NewArgument("device", "Terminal device", nil, "string", "device", nil),
			syntax.NewArgument("remote", "Remote Device", nil, "string", "terminal", nil),
		}, nil)
	terminalCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("terminal executed on device: %#v remote: %#v\n", params["device"], params["remote"])
		return nil
	}

	systemCmd := syntax.NewCommand(nil, "system <ip>", "Configure terminal",
		[]*syntax.Argument{
			syntax.NewArgument("ip", "IP address", nil, "string", "", nil),
		}, nil)
	systemCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("system for ip: %#v\n", params["ip"])
		return nil
	}

	helpCmd := syntax.NewCommand(nil, "helper [<set> | <get>]?", "Application help",
		[]*syntax.Argument{
			syntax.NewArgument("set", "Help for set command", nil, "string", "", nil),
			syntax.NewArgument("speed", "Help for set speed command", nil, "string", "", nil),
			syntax.NewArgument("get", "Help for get command", nil, "string", "", nil),
		}, nil)
	helpCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("help for command: %#v %#v %#v\n", params["set"], params["speed"], params["get"])
		return nil
	}

	configIPCmd := syntax.NewCommand(configCmd, "ip [<IPV4> addr4 | <IPV6> addr6]!", "Configure ip for IPV4 or IPV6",
		[]*syntax.Argument{
			syntax.NewArgument("IPV4", "IPV4 address schema", nil, "string", "IPV4", nil),
			syntax.NewArgument("addr4", "IPV4 address", nil, "string", "0.0.0.0", nil),
			syntax.NewArgument("IPV6", "IPV6 address schema", nil, "string", "IPV6", nil),
			syntax.NewArgument("addr6", "IPV6 address", nil, "string", "0:0:0:0:0:0", nil),
		}, nil)
	configIPCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("config ip address IPV4: %#v:%#v IPV6: %#v:%#v\n", params["IPV4"], params["addr4"], params["IPV6"], params["addr6"])
		return nil
	}

	configEthCmd := syntax.NewCommand(configCmd, "eth [mac]@", "Configure ethernet mac address",
		[]*syntax.Argument{
			syntax.NewArgument("mac", "mac address", nil, "freeform", "nothing to declare", nil),
		}, nil)
	configEthCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("config mac address: %#v\n", params["mac"])
		return nil
	}

	configDbaseCmd := syntax.NewCommand(configCmd, "dbase [fields]@", "Configure database",
		[]*syntax.Argument{
			syntax.NewArgument("fields", "dbase fields", nil, "map", "base:none", nil),
		}, nil)
	configDbaseCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("configure database: %#v\n", params["fields"])
		return nil
	}

	commands := []*syntax.Command{
		setCmd,
		getCmd,
		setBaudrateCmd,
		setSpeedCmd,
		setBoolCmd,
		getBaudrateCmd,
		getSpeedCmd,
		setSpeedDeviceCmd,
		terminalCmd,
		systemCmd,
		helpCmd,
		configCmd,
		configDescCmd,
		configIPCmd,
		configDbaseCmd,
		tenantCmd,
	}
	return commands
}
