package commands

import (
	"os"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// SetupCommands configures all command to run.
func SetupCommands() []*syntax.Command {
	exitCmd := syntax.NewCommand(nil, "exit", "Exit application", nil, nil)
	exitCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		os.Exit(0)
		return nil
	}

	setCmd := syntax.NewCommand(nil, "set version", "Set test help",
		[]*syntax.Argument{
			syntax.NewArgument("version", "Version number", nil, "int", 0, nil),
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
		tools.ToDisplay("executing set speed\n")
		return nil
	}

	setSpeedDeviceCmd := syntax.NewCommand(setSpeedCmd, "device name", "Set speed device help",
		[]*syntax.Argument{
			syntax.NewArgument("name", "Device name", nil, "string", "", nil),
		}, nil)
	setSpeedDeviceCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		tools.ToDisplay("executing set speed device with device: %#v\n", params["name"])
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

	commands := []*syntax.Command{
		exitCmd,
		setCmd,
		getCmd,
		syntax.NewCommand(nil, "config", "Config test help", nil, nil),
		setBaudrateCmd,
		setSpeedCmd,
		setBoolCmd,
		syntax.NewCommand(getCmd, "baudrate", "Get Baudrate test help", nil, nil),
		//syntax.NewCommand(getCmd, "speed", "Get Speed test help", nil, nil),
		getSpeedCmd,
		setSpeedDeviceCmd,
	}
	return commands
}
