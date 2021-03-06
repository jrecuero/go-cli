package main

import (
	"fmt"
	"strings"

	"github.com/jrecuero/go-cli/app/cli/syntax/commands"
	"github.com/jrecuero/go-cli/graph"
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
	g := graph.NewGraph(nil)
	var c1 = graph.NewNode("Jose Carlos", nil)
	var c2 = graph.NewNode("Marcela Veronica", nil)
	g.Root.AddChild(c1)
	g.Root.AddChild(c2)
	fmt.Println(g.Root.Label, g.Root.Children)
	for _, child := range g.Root.Children {
		fmt.Println(child.Label, child.Children)
	}
}

func runGraph() {
	g := graph.NewGraph(nil)
	c1 := graph.NewNode("Jose Carlos", nil)
	g.AddNode(c1)
	c2 := graph.NewNode("Marcela Veronica", nil)
	g.AddNode(c2)
	g.Terminate()
	fmt.Println(g.ToString())
}

func runSyntax() {
	//cs := syntax.NewCommandSyntax("SELECT name [age]? [id | passport]+")
	cs := syntax.NewCommandSyntax("SELECT name [age]?")
	fmt.Printf("%s\n%#v\n%p\n", cs.Syntax, cs.Parsed, cs.Syntax)
	cs.CreateGraph(nil)
	//cs.Graph.Explore()
	fmt.Printf("%s", cs.Graph.ToString())
}

func runMermaid() {
	cs := syntax.NewCommandSyntax("SELECT name [age]? [id | passport]+ [phone]* [male | female | other]!")
	fmt.Printf("Syntax: %s\nParsed: %#v\nGraph: %p\n", cs.Syntax, cs.Parsed, cs.Graph)
	cs.CreateGraph(nil)
	fmt.Printf("%s", cs.Graph.ToMermaid())
}

func runSimpleMatcher() {
	cs := syntax.NewCommandSyntax("SELECT name age")
	cs.CreateGraph(nil)
	fmt.Printf("%s", cs.Graph.ToString())
	m := syntax.NewMatcher(syntax.NewContext(nil), cs.Graph)
	line := []string{"SELECT", "name", "age"}
	m.Match(line)
}

func runComplexMatcher() {
	cs := syntax.NewCommandSyntax("SELECT name [ age  | id ]?")
	cs.CreateGraph(nil)
	fmt.Printf("%s", cs.Graph.ToString())
	m := syntax.NewMatcher(syntax.NewContext(nil), cs.Graph)
	//line := []string{"SELECT", "name", "age"}
	//line := []string{"SELECT", "name", "id"}
	//line := []string{"SELECT", "name"}
	line := []string{"SELECT", "name", "caca"}
	fmt.Println(m.Match(line))
}

//func cmdHello(context interface{}, arguments interface{}) bool {
//    //name := arguments[0].(string)
//    data := arguments.(map[string]string)
//    name := data["name"]
//    fmt.Printf("Context: %v, Arguments: %v, Name: %s\n", context, arguments, name)
//    return true
//}

//func runExecute() {
//    cmd := &syntax.Command{
//        Cb:        cmdHello,
//        Syntax:    nil,
//        Label:     "hello",
//        Name:      "hello",
//        Help:      "Hello command.",
//        Arguments: nil,
//    }
//    //arguments := []interface{}{"Jose Carlos", "Recuero Arias"}
//    arguments := map[string]string{"name": "Jose Carlos"}
//    syntax.Execute(cmd, arguments, nil)
//}

func testa(data interface{}) {
	fmt.Println(data)
	mapa := data.(map[string]string)
	fmt.Println(mapa["name"])
	fmt.Println(mapa["last name"])
}

func printCompleterInfo(ic syntax.ICompleter) {
	fmt.Printf("new completer is: %#v\n", ic)
	fmt.Printf("completer label is: %s\n", ic.GetLabel())
}

//func runCompleter() {
//    printCompleterInfo(syntax.NewCIdent("me"))
//    printCompleterInfo(syntax.NewCJoint(""))
//    printCompleterInfo(syntax.NewCStart())
//    printCompleterInfo(syntax.NewCEnd())
//    printCompleterInfo(syntax.NewCLoop())
//}

func runUserCmd() {
	userCmd := commands.UserCmd
	userCmd.Setup()
	fmt.Println(userCmd)
	fmt.Printf("%#v\n", userCmd.CmdSyntax)
	fmt.Printf("%#v\n", userCmd.CmdSyntax.Parsed)
	//fmt.Printf("%#v\n", userCmd.CmdSyntax.Graph)
	fmt.Printf("\n%s\n", userCmd.CmdSyntax.Graph.ToMermaid())
	fmt.Printf("\n%s\n", userCmd.CmdSyntax.Graph.ToContent())
	//fmt.Println(commands.Group)
	userCmd.Enter(nil, map[string]interface{}{"name": "Jose Carlos", "age": 51})
	managerCmd := commands.ManagerCmd
	managerCmd.Enter(nil, map[string]interface{}{"name": "Jose Carlos"})
	userCmd.CmdSyntax.Graph.Explore()
}

func runSetSpeedCmd() {
	setSpeedCmd := commands.SetSpeedCmd
	setSpeedCmd.Setup()
	fmt.Println(setSpeedCmd)
	fmt.Printf("%#v\n", setSpeedCmd.CmdSyntax)
	fmt.Printf("%#v\n", setSpeedCmd.CmdSyntax.Parsed)
	fmt.Printf("\n%s\n", setSpeedCmd.CmdSyntax.Graph.ToMermaid())
	fmt.Printf("\n%s\n", setSpeedCmd.CmdSyntax.Graph.ToContent())
	//setSpeedCmd.Enter(nil, map[string]interface{}{"name": "Jose Carlos", "age": 51})
	setSpeedCmd.CmdSyntax.Graph.Explore()
}

func runManagerCmd() {
	//command := commands.UserCmd
	command := commands.ManagerCmd
	command.Setup()
	cs := command.CmdSyntax
	//fmt.Printf("%s", cs.Graph.ToString())
	//fmt.Printf("%s", cs.Graph.ToMermaid())
	m := syntax.NewMatcher(syntax.NewContext(nil), cs.Graph)
	fmt.Printf("%#v\n", m)
	//line := []string{"user", "name", "age"}
	line := []string{"manager", "name"}
	m.Match(line)
}

func runNewUserCmd() {
	command := commands.UserCmd
	command.Setup()
	cs := command.CmdSyntax
	//fmt.Printf("%s", cs.Graph.ToString())
	m := syntax.NewMatcher(syntax.NewContext(nil), cs.Graph)
	fmt.Printf("%#v\n", m)
	line := []string{"user", "josecarlos", "-age", "51"}
	m.Match(line)
	//fmt.Printf("%s", cs.Graph.ToMermaid())
	for _, token := range m.Ctx.Matched {
		fmt.Printf("[%s] %s : %s \n", token.Node.GetContent().GetStrType(), token.Node.Label, token.Value)
	}
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
	//runComplexMatcher()
	//runCompleter()
	//runUserCmd()
	//runSetSpeedCmd()
	//runManagerCmd()
	runNewUserCmd()
}
