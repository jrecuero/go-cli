package main

import (
	"fmt"
	"strings"

	"github.com/jrecuero/go-cli/cmd/cli/syntax/commands"
	"github.com/jrecuero/go-cli/parser"
	"github.com/jrecuero/go-cli/syntax"
)

func runParser() {
	s := "SELECT name [ age ]?"
	syntax, err := parser.NewParser(strings.NewReader(s)).Parse()
	fmt.Println("command: ", syntax.Command)
	fmt.Println("arguments: ", syntax.Arguments)
	fmt.Println(err)
}

func runNode() {
	g := syntax.NewGraph()
	var c1 = syntax.NewNode("Jose Carlos", "JC", nil)
	var c2 = syntax.NewNode("Marcela Veronica", "MV", nil)
	g.Root.AddChild(c1)
	g.Root.AddChild(c2)
	fmt.Println(g.Root.ID, g.Root.Name, g.Root.Label, g.Root.Children)
	for _, child := range g.Root.Children {
		fmt.Println(child.ID, child.Name, child.Label, child.Children)
	}
}

func runGraph() {
	g := syntax.NewGraph()
	c1 := syntax.NewNode("Jose Carlos", "JC", nil)
	g.AddNode(c1)
	c2 := syntax.NewNode("Marcela Veronica", "MV", nil)
	g.AddNode(c2)
	g.Terminate()
	fmt.Println(g.ToString())
}

func runSyntax() {
	//cs := syntax.NewCommandSyntax("SELECT name [age]? [id | passport]+")
	cs := syntax.NewCommandSyntax("SELECT name [age]?")
	fmt.Printf("%s\n%#v\n%p\n", cs.Syntax, cs.Parsed, cs.Syntax)
	cs.CreateGraph()
	//cs.syntax.Explore()
	fmt.Printf("%s", cs.Graph.ToString())
}

func runMermaid() {
	cs := syntax.NewCommandSyntax("SELECT name [age]? [id | passport]+ [phone]* [male | female | other]!")
	fmt.Printf("Syntax: %s\nParsed: %#v\nGraph: %p\n", cs.Syntax, cs.Parsed, cs.Graph)
	cs.CreateGraph()
	fmt.Printf("%s", cs.Graph.ToMermaid())
}

func runSimpleMatcher() {
	cs := syntax.NewCommandSyntax("SELECT name age")
	cs.CreateGraph()
	fmt.Printf("%s", cs.Graph.ToString())
	m := syntax.NewMatcher(syntax.NewContext(), cs.Graph)
	line := []string{"SELECT", "name", "age"}
	m.MatchCommandLine(line)
}

func runComplexMatcher() {
	cs := syntax.NewCommandSyntax("SELECT name [ age  | id ]?")
	cs.CreateGraph()
	fmt.Printf("%s", cs.Graph.ToString())
	m := syntax.NewMatcher(syntax.NewContext(), cs.Graph)
	//line := []string{"SELECT", "name", "age"}
	//line := []string{"SELECT", "name", "id"}
	//line := []string{"SELECT", "name"}
	line := []string{"SELECT", "name", "caca"}
	fmt.Println(m.MatchCommandLine(line))
}

func cmdHello(context interface{}, arguments interface{}) bool {
	//name := arguments[0].(string)
	data := arguments.(map[string]string)
	name := data["name"]
	fmt.Printf("Context: %v, Arguments: %v, Name: %s\n", context, arguments, name)
	return true
}

func runExecute() {
	cmd := &syntax.Command{
		Cb:        cmdHello,
		Syntax:    nil,
		Label:     "hello",
		Name:      "hello",
		Help:      "Hello command.",
		Arguments: nil,
	}
	//arguments := []interface{}{"Jose Carlos", "Recuero Arias"}
	arguments := map[string]string{"name": "Jose Carlos"}
	syntax.Execute(cmd, arguments, nil)
}

func testa(data interface{}) {
	fmt.Println(data)
	mapa := data.(map[string]string)
	fmt.Println(mapa["name"])
	fmt.Println(mapa["last name"])
}

func printCompleterInfo(ic syntax.ICompleter) {
	fmt.Printf("new completer is: %#v\n", ic)
	fmt.Printf("completer label is: %s\n", ic.GetLabel())
	fmt.Printf("completer content is: %#v\n", ic.GetContent())
}

func runCompleter() {
	printCompleterInfo(syntax.NewIdentCompleter("me", nil))
	printCompleterInfo(syntax.NewJointCompleter(""))
	printCompleterInfo(syntax.NewStartCompleter())
	printCompleterInfo(syntax.NewEndCompleter())
	printCompleterInfo(syntax.NewLoopCompleter())
}

func main() {
	//runParser()
	//runNode()
	//runGraph()
	//runSyntax()
	//runMermaid()
	//runExecute()
	//testa(map[string]string{"name": "Jose Carlos", "last name": "Recuero Arias"})
	//runSimpleMatcher()
	runComplexMatcher()
	//runCompleter()
	fmt.Println(commands.User)
	fmt.Println(commands.Group)
	commands.User.Cb(nil, map[string]interface{}{"name": "Jose Carlos", "age": 51})
}
