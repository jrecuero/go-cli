package prompter

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// LivePrefixState represent the prompt.
type LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

// Prompter represents the CLI propmpt.
type Prompter struct {
	NSH         *syntax.NSHandler
	NSM         *syntax.NSManager
	NS          *syntax.NameSpace
	PrefixState *LivePrefixState
	_prompt     *prompt.Prompt
}

// executor executes any command entered in the command line.
func (pr *Prompter) executor(in string) {
	if in == "quit" {
		tools.CloseLog()
		os.Exit(0)
	}
	fmt.Println("Your input: " + in)
	tools.Info("Running command line %#v\n", in)
	// Pass the NSManager to the context cache to be used by internal commands.
	pr.NSM.GetContext().Cache.Add("nsm", pr.NSM)
	pr.NSM.GetContext().Cache.Add("prompt", pr._prompt)
	if retcode, err := pr.NSM.Execute(in); err == nil {
		if retcode == syntax.POPMODE {
			prompt.OptionInputTextColor(prompt.DefaultColor)(pr._prompt)
		}
	} else {
		fmt.Errorf("execute return %#v for line: %s", err, in)
		return
	}
	//tools.ToDisplay(pr.NSM.GetCommandTree().ToMermaid())
	//tools.ToDisplay(pr.NSM.GetParseTree().ToMermaid())
	pr.PrefixState.LivePrefix = pr.NSM.GetContext().GetPrompt()
	pr.PrefixState.IsEnable = true
}

// completer completes any token being entered in the command line.
func (pr *Prompter) completer(d prompt.Document) []prompt.Suggest {
	//tools.Tracer("document: %#v\n", d)
	if d.IsReset {
		return []prompt.Suggest{}
	}
	line := d.TextBeforeCursor()
	if line == "" {
		line = " "
	}
	result, _ := pr.NSM.CompleteAndHelp(line)
	var s []prompt.Suggest
	var varArgs []prompt.Suggest
	for _, r := range result.([]*syntax.CompleteHelp) {
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
	return pr.PrefixState.LivePrefix, pr.PrefixState.IsEnable
}

// Setup initializes the prompter with the given commands.
func (pr *Prompter) Setup(nsname string, commands ...*syntax.Command) error {
	if nsh, err := syntax.CreateNSHandler(nsname, commands); err == nil {
		pr.NSH = nsh
		pr.NSM = nsh.GetActive().NSMgr
		pr.NS = nsh.GetActive().NS
		pr.PrefixState = &LivePrefixState{
			LivePrefix: syntax.DEFAULTPROMPT,
			IsEnable:   true,
		}
		return nil
	}
	return fmt.Errorf("Prompter setup error")
}

// Run runs the prompt.
func (pr *Prompter) Run() {
	fmt.Println("Please select:")
	pr._prompt = prompt.New(
		pr.executor,
		pr.completer,
		prompt.OptionPrefix(syntax.DEFAULTPROMPT),
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
	pr._prompt.Run()
}

// NewPrompter creates a new Prompter instance.
func NewPrompter() *Prompter {
	return &Prompter{}
}
