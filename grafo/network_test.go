package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
	"github.com/jrecuero/go-cli/tools"
)

// Test_Network is ...
func Test_Network(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	root := grafo.NewLeaf("origin/1")
	node1 := grafo.NewLeaf("node/1")
	node2 := grafo.NewLeaf("node/2")
	node3 := grafo.NewLeaf("node/3")
	node4 := grafo.NewLeaf("node/4")
	node5 := grafo.NewLeaf("node/5")
	network.AddNode(nil, root, 0)
	network.AddNode(root, node1, 10)
	network.AddNode(root, node2, 5)
	network.AddNode(node2, node3, 3)
	network.AddNode(node2, node4, 6)
	network.AddNode(node1, node5, 1)
	network.AddNode(node3, node5, 7)
	network.AddNode(node4, node5, 11)
	//for k, v := range network.GetLeafs() {
	//    tools.ToDisplay("%d : %#v\n\n", k, v)
	//}
	tools.ToDisplay("%s\n", network.ToMermaid())
}
