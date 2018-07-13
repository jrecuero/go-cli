package prompter

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jrecuero/go-cli/syntax"
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

// executor executes any command entered in the command line.
func (pr *Prompter) executor(in string) {
	//if in == "exit" || in == "quit" {
	if in == "quit" {
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

// completer completes any token being entered in the command line.
func (pr *Prompter) completer(d prompt.Document) []prompt.Suggest {
	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, pr.NSM.GetParseTree().Graph)
	line := d.TextBeforeCursor()
	if line == "" {
		line = " "
	}
	result, _ := m.CompleteAndHelp(line)
	var s []prompt.Suggest
	var varArgs []prompt.Suggest
	for _, r := range result.([]*syntax.ComplexComplete) {
		//tools.Tracer("result: %#v\n", r)
		completeStr := r.Complete.(string)
		newSuggest := prompt.Suggest{Text: completeStr, Description: r.Help.(string)}
		if strings.HasPrefix(completeStr, "<<") && strings.HasSuffix(completeStr, ">>") {
			varArgs = append(varArgs, newSuggest)
		}
		s = append(s, newSuggest)
	}
	suggests := prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	if len(suggests) == 0 {
		suggests = varArgs
	}
	return suggests
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
		prompt.OptionTitle("go-cli"),
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
