package prompter

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// Prompter represents the CLI propmpt.
type Prompter struct {
	NSH *syntax.NSHandler
	NSM *syntax.NSManager
	NS  *syntax.NameSpace
}

var livePrefixState struct {
	livePrefix string
	isEnable   bool
}

func (pr *Prompter) executor(in string) {
	if in == "exit" || in == "quit" {
		os.Exit(0)
	}
	fmt.Println("Your input: " + in)
	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, pr.NSM.GetParseTree().Graph)
	if _, ok := m.Execute(in); !ok {
		fmt.Errorf("execute return %#v for line: %s", ok, in)
		return
	}
}

func (pr *Prompter) completer(d prompt.Document) []prompt.Suggest {
	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, pr.NSM.GetParseTree().Graph)
	line := d.TextBeforeCursor()
	if line == "" {
		line = " "
	}
	//result, _ := m.Complete(line)
	result, _ := m.CompleteAndHelp(line)
	var s []prompt.Suggest
	for _, r := range result.([]*syntax.ComplexComplete) {
		//for _, r := range result.([]interface{}) {
		//s = append(s, prompt.Suggest{Text: r.(string), Description: r.(string)})
		tools.Tracer("%#v\n", r)
		s = append(s, prompt.Suggest{Text: r.Complete.(string), Description: r.Help.(string)})
	}
	//s := []prompt.Suggest{
	//    {Text: "users", Description: "Store the username and age"},
	//    {Text: "articles", Description: "Store the article text posted by user"},
	//    {Text: "comments", Description: "Store the text commented to articles"},
	//}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (pr *Prompter) changeLivePrefix() (string, bool) {
	return livePrefixState.livePrefix, livePrefixState.isEnable
}

// Setup initializes the prompter with the given commands.
func (pr *Prompter) Setup(nsname string, commands []*syntax.Command) error {
	if nsh, err := syntax.CreateNSHandler(nsname, commands); err == nil {
		pr.NSH = nsh
		pr.NSM = nsh.GetActive().NSMgr
		pr.NS = nsh.GetActive().NS
		return nil
	}
	return fmt.Errorf("Prompter setup error")
}

// Run runs the prompt.
func (pr *Prompter) Run() {
	fmt.Println("Please select:")
	p := prompt.New(
		pr.executor,
		pr.completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(pr.changeLivePrefix),
		prompt.OptionTitle("cli-prompt"),
		prompt.OptionHistory([]string{}),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionSuggestionFilter(func(suggest prompt.Suggest) (prompt.Suggest, bool) {
			if text := suggest.Text; strings.HasPrefix(text, "<<") && strings.HasSuffix(text, ">>") {
				return prompt.Suggest{}, false
			}
			return suggest, true
		}),
	)
	p.Run()
}
