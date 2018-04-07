package main

import (
	"fmt"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/syntax"
)

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

func main() {
	//runNode()
	//runGraph()
	//runSyntax()
	//runMermaid()
	runExecute()
	testa(map[string]string{"name": "Jose Carlos", "last name": "Recuero Arias"})
}
