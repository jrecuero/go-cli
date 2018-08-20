package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
)

// TestNetwork_Network is ...
func TestNetwork_Network(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	if network == nil {
		t.Errorf("Tree:NewNetwork: network can not be <nil>")
	}
	root := grafo.NewLeaf("origin/1")
	node1 := grafo.NewLeaf("node/1")
	node2 := grafo.NewLeaf("node/2")
	node3 := grafo.NewLeaf("node/3")
	node4 := grafo.NewLeaf("node/4")
	node5 := grafo.NewLeaf("node/5")
	if err := network.AddNode(nil, root, 0); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(root, node1, 10); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(root, node2, 5); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node2, node3, 3); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node2, node4, 6); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node1, node5, 1); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node3, node5, 7); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node4, node5, 11); err != nil {
		t.Errorf("Tree:AddNode: return error code: %#v\n", err)
	}
	//tools.ToDisplay("%s\n", network.ToMermaid())
}

// TestNetwork_CostToNode is ...
func TestNetwork_CostToNode(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	root := grafo.NewNode("origin/1", grafo.NewNodeContent("origin/1", 0))
	node1 := grafo.NewNode("node/1", grafo.NewNodeContent("node/1", 2))
	node2 := grafo.NewNode("node/2", grafo.NewNodeContent("node/2", 3))
	node3 := grafo.NewNode("node/3", grafo.NewNodeContent("node/3", 4))
	network.AddNode(nil, grafo.ToLeaf(root), 0)
	network.AddNode(grafo.ToLeaf(root), grafo.ToLeaf(node1), 10)
	network.AddNode(grafo.ToLeaf(root), grafo.ToLeaf(node2), 5)
	network.AddNode(grafo.ToLeaf(node2), grafo.ToLeaf(node3), 1)
	network.SetAnchorTo(grafo.ToLeaf(root))
	if w, ok := network.CostToNode(node1); ok {
		if w != 12 {
			t.Errorf("Tree:CostToNode: incorrect weight: exp: %d got %d\n", 12, w)
		}
	} else {
		t.Errorf("Tree:CostToNode: branch not available from: %#v to %#v\n", root, node1)
	}
	if _, ok := network.CostToNode(node3); ok {
		t.Errorf("Tree:CostToNode: branch available from: %#v to %#v\n", root, node3)
	}
}

// TestNetwork_PathFromTo is ...
func TestNetwork_PathFromTo(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	root := grafo.NewNode("root/1", grafo.NewNodeContent("root/1", 0))
	node1 := grafo.NewNode("node/1", grafo.NewNodeContent("node/1", 2))
	node2 := grafo.NewNode("node/2", grafo.NewNodeContent("node/2", 3))
	node3 := grafo.NewNode("node/3", grafo.NewNodeContent("node/3", 4))
	node4 := grafo.NewNode("node/4", grafo.NewNodeContent("node/4", 4))
	network.AddNode(nil, grafo.ToLeaf(root), 0)
	network.AddNode(grafo.ToLeaf(root), grafo.ToLeaf(node1), 10)
	network.AddNode(grafo.ToLeaf(root), grafo.ToLeaf(node2), 5)
	network.AddNode(grafo.ToLeaf(node1), grafo.ToLeaf(node3), 1)
	network.AddNode(grafo.ToLeaf(node2), grafo.ToLeaf(node3), 1)
	network.AddNode(grafo.ToLeaf(node3), grafo.ToLeaf(node4), 1)
	network.AddNode(grafo.ToLeaf(node3), grafo.ToLeaf(root), 1)
	network.AddNode(grafo.ToLeaf(node1), grafo.ToLeaf(node4), 1)
	//tools.ToDisplay("%s\n", network.ToMermaid())
	if paths := network.PathsFromNodeToNode(root, node3); paths == nil {
		t.Errorf("Tree:PathsFromNodeToNode: no paths found from: %#v to %#v\n", root, node3)
	} else {
		if len(paths) != 2 {
			t.Errorf("Tree:PathsFromNodeToNode: paths lenfgth: exp: %d got: %d\n", 2, len(paths))
		}
	}
	if paths := network.PathsFromNodeToNode(root, node4); paths == nil {
		t.Errorf("Tree:PathsFromNodeToNode: no paths found from: %#v to %#v\n", root, node4)
	} else {
		if len(paths) != 3 {
			t.Errorf("Tree:PathsFromNodeToNode: paths lenfgth: exp: %d got: %d\n", 3, len(paths))
		}
		//for _, p := range paths {
		//    tools.ToDisplay("%s\n", p.Label)
		//    for _, t := range p.Traversed {
		//        tools.ToDisplay("\t%s -> %s\n", t.Parent.Label, t.Child.Label)
		//    }
		//}
	}
	if paths := network.PathsFromNodeToNode(node1, root); paths == nil {
		t.Errorf("Tree:PathsFromNodeToNode: no paths found from: %#v to %#v\n", node1, root)
	} else {
		if len(paths) != 1 {
			t.Errorf("Tree:PathsFromNodeToNode: paths lenfgth: exp: %d got: %d\n", 1, len(paths))
		}
	}
}
