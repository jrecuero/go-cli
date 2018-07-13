package main

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/jrecuero/go-cli/app/cli/prompt/commands"
	"github.com/jrecuero/go-cli/prompter"
	"github.com/jrecuero/go-cli/tools"
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
	tools.ToDisplay("Your input: " + in)
	testArray := strings.Fields(in)
	tools.ToDisplay("Number of tokens: ", len(testArray))
	for _, v := range testArray {
		tools.ToDisplay("token: " + v)
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
	tools.ToDisplay("Please select table.")
	//t := prompt.Input("> ", completer)
	//tools.ToDisplay("You selected " + t)
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

func runPrompter() {
	p := &prompter.Prompter{}
	p.Setup("prompter", commands.SetupCommands())
	p.Run()
}

func main() {
	//runGoPrompt()
	//runInputSurvey()
	//runMultiSelectSurvey()
	runPrompter()
}
