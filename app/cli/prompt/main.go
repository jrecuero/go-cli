package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/jrecuero/go-cli/prompter"
	"github.com/jrecuero/go-cli/syntax"
	"gopkg.in/AlecAivazis/survey.v1"
)

var livePrefixState struct {
	livePrefix string
	isEnable   bool
}

func executor(in string) {
	if in == "exit" || in == "quit" {
		os.Exit(0)
	}
	fmt.Println("Your input: " + in)
	testArray := strings.Fields(in)
	fmt.Println("Number of tokens: ", len(testArray))
	for _, v := range testArray {
		fmt.Println("token: " + v)
	}
	livePrefixState.livePrefix = in + "> "
	livePrefixState.isEnable = true
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func changeLivePrefix() (string, bool) {
	return livePrefixState.livePrefix, livePrefixState.isEnable
}

func runGoPrompt() {
	fmt.Println("Please select table.")
	//t := prompt.Input("> ", completer)
	//fmt.Println("You selected " + t)
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("cli-prompt"),
		prompt.OptionHistory([]string{}),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
	)
	p.Run()
}

func runInputSurvey() {
	name := ""
	prompt := &survey.Input{
		Message: "ping",
	}
	survey.AskOne(prompt, &name, nil)
}

func runMultiSelectSurvey() {
	days := []string{}
	prompt := &survey.MultiSelect{
		Message: "What days do you prefer?",
		Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	}
	survey.AskOne(prompt, &days, nil)
}

func setupCommands() []*syntax.Command {
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
			fmt.Println("Error: %#v\n", err)
		}
		fmt.Println("executing enter with version(ctx):", version)
		params := arguments.(map[string]interface{})
		fmt.Println("executing enter wit version(args):", params["version"])
		return nil
	}
	getCmd := syntax.NewCommand(nil, "get", "Get test help", nil, nil)
	setBoolCmd := syntax.NewCommand(setCmd, "bool", "Set Bool test help", nil, nil)
	setBaudrateCmd := syntax.NewCommand(setCmd, "baudrate [speed | parity]?", "Set baudrate help",
		[]*syntax.Argument{
			syntax.NewArgument("speed", "Baudrate speed", nil, "string", "", nil),
			syntax.NewArgument("parity", "Baudrate parity value", nil, "string", "", nil),
		}, nil)
	setSpeedCmd := syntax.NewCommand(setCmd, "speed", "Set Speed test help", nil, nil)
	setSpeedDeviceCmd := syntax.NewCommand(setSpeedCmd, "device name", "Set speed device help",
		[]*syntax.Argument{
			syntax.NewArgument("name", "Device name", nil, "string", "", nil),
		}, nil)
	getSpeedCmd := syntax.NewCommand(getCmd, "speed [device name | value]?", "Get speed help",
		[]*syntax.Argument{
			syntax.NewArgument("device", "Device", nil, "string", "", nil),
			syntax.NewArgument("name", "Device name", nil, "string", "", nil),
			syntax.NewArgument("value", "Speed value", nil, "string", "", nil),
		}, nil)
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

func runPrompter() {
	p := &prompter.Prompter{}
	p.Setup("prompter", setupCommands())
	p.Run()
}

func main() {
	//runGoPrompt()
	//runInputSurvey()
	//runMultiSelectSurvey()
	runPrompter()
}
