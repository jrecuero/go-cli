package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/jrecuero/go-cli/graph"
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

func runNode() {
	g := graph.NewGraph()
	var c1 = graph.NewNode("Jose Carlos", "JC")
	var c2 = graph.NewNode("Marcela Veronica", "MV")
	g.Root.AddChild(c1)
	g.Root.AddChild(c2)
	fmt.Println(g.Root.ID, g.Root.Name, g.Root.Label, g.Root.Children)
	for _, child := range g.Root.Children {
		fmt.Println(child.ID, child.Name, child.Label, child.Children)
	}
}

func runGraph() {
	g := graph.NewGraph()
	c1 := graph.NewNode("Jose Carlos", "JC")
	g.AddNode(c1)
	c2 := graph.NewNode("Marcela Veronica", "MV")
	g.AddNode(c2)
	g.Terminate()
	fmt.Println(g.ToString())
}

func runSyntax() {
	cs := syntax.NewCommandSyntax("SELECT name [age]? [id | passport]+")
	fmt.Printf("%s\n%#v\n%p\n", cs.Syntax, cs.Parsed, cs.Graph)
	cs.CreateGraph()
	//cs.Graph.Explore()
	fmt.Printf("%s", cs.Graph.ToString())
}

func runMermaid() {
	cs := syntax.NewCommandSyntax("SELECT name [age]? [id | passport]+ [phone]* [male | female | other]!")
	fmt.Printf("Syntax: %s\nParsed: %#v\nGraph: %p\n", cs.Syntax, cs.Parsed, cs.Graph)
	cs.CreateGraph()
	fmt.Printf("%s", cs.Graph.ToMermaid())
}

func main() {
	//runGoPrompt()
	//runInputSurvey()
	//runMultiSelectSurvey()
	//runNode()
	//runGraph()
	//runSyntax()
	runMermaid()
}
