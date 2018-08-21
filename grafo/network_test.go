package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
)

// TestNetwork_Network is ...
func TestNetwork_Network(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	if network == nil {
		t.Errorf("Network:NewNetwork: network can not be <nil>")
	}
	root := grafo.NewLeaf("origin/1")
	node1 := grafo.NewLeaf("node/1")
	node2 := grafo.NewLeaf("node/2")
	node3 := grafo.NewLeaf("node/3")
	node4 := grafo.NewLeaf("node/4")
	node5 := grafo.NewLeaf("node/5")
	if err := network.AddNode(nil, root, 0); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(root, node1, 10); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(root, node2, 5); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node2, node3, 3); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node2, node4, 6); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node1, node5, 1); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node3, node5, 7); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := network.AddNode(node4, node5, 11); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
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
	network.AddNode(nil, grafo.NodeToLeaf(root), 0)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node1), 10)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node2), 5)
	network.AddNode(grafo.NodeToLeaf(node2), grafo.NodeToLeaf(node3), 1)
	network.SetAnchorTo(grafo.NodeToLeaf(root))
	if w, ok := network.CostToNode(node1); ok {
		if w != 12 {
			t.Errorf("Network:CostToNode: incorrect weight mismatch: exp: %d got %d\n", 12, w)
		}
	} else {
		t.Errorf("Network:CostToNode: branch not available from: %#v to %#v\n", root.Label, node1.Label)
	}
	if _, ok := network.CostToNode(node3); ok {
		t.Errorf("Network:CostToNode: branch available from: %#v to %#v\n", root.Label, node3.Label)
	}
}

// TestNetwork_PathsFromNodeToNode is ...
func TestNetwork_PathsFromNodeToNode(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	root := grafo.NewNode("root/1", grafo.NewNodeContent("root/1", 0))
	node1 := grafo.NewNode("node/1", grafo.NewNodeContent("node/1", 2))
	node2 := grafo.NewNode("node/2", grafo.NewNodeContent("node/2", 3))
	node3 := grafo.NewNode("node/3", grafo.NewNodeContent("node/3", 4))
	node4 := grafo.NewNode("node/4", grafo.NewNodeContent("node/4", 4))
	node5 := grafo.NewNode("node/5", grafo.NewNodeContent("node/5", 5))
	network.AddNode(nil, grafo.NodeToLeaf(root), 0)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node1), 10)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node2), 5)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node2), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node3), grafo.NodeToLeaf(node4), 1)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node4), 1)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node5), 1)
	if paths := network.PathsFromNodeToNode(root, node3); paths == nil {
		t.Errorf("Network:PathsFromNodeToNode: no paths found from: %#v to %#v\n", root.Label, node3.Label)
	} else {
		if len(paths) != 2 {
			t.Errorf("Network:PathsFromNodeToNode: paths length mismatch: exp: %d got: %d\n", 2, len(paths))
		}
	}
	if paths := network.PathsFromNodeToNode(root, node4); paths == nil {
		t.Errorf("Network:PathsFromNodeToNode: no paths found from: %#v to %#v\n", root.Label, node4.Label)
	} else {
		if len(paths) != 3 {
			t.Errorf("Network:PathsFromNodeToNode: paths length mismatch: exp: %d got: %d\n", 3, len(paths))
		}
		//for _, p := range paths {
		//    tools.ToDisplay("%s\n", p)
		//}
	}
	if paths := network.PathsFromNodeToNode(node2, node5); paths != nil {
		t.Errorf("Network:PathsFromNodeToNode: paths found from: %#v to %#v\n", node2.Label, node5.Label)
	}

	network.AddNode(grafo.NodeToLeaf(node3), grafo.NodeToLeaf(root), 1)
	//tools.ToDisplay("%s\n", network.ToMermaid())
	if paths := network.PathsFromNodeToNode(node1, root); paths == nil {
		t.Errorf("Network:PathsFromNodeToNode: no paths found from: %#v to %#v\n", node1.Label, root.Label)
	} else {
		if len(paths) != 1 {
			t.Errorf("Network:PathsFromNodeToNode: paths length mismatch: exp: %d got: %d\n", 1, len(paths))
		}
	}
}

// TestNetwork_TotalWeightInPath is ...
func TestNetwork_TotalWeightInPath(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	root := grafo.NewNode("root/1", grafo.NewNodeContent("root/1", 0))
	node1 := grafo.NewNode("node/1", grafo.NewNodeContent("node/1", 1))
	node2 := grafo.NewNode("node/2", grafo.NewNodeContent("node/2", 2))
	node3 := grafo.NewNode("node/3", grafo.NewNodeContent("node/3", 3))
	node4 := grafo.NewNode("node/4", grafo.NewNodeContent("node/4", 4))
	network.AddNode(nil, grafo.NodeToLeaf(root), 0)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node1), 10)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node2), 5)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node2), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node3), grafo.NodeToLeaf(node4), 1)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node4), 1)
	if paths := network.PathsFromNodeToNode(root, node4); paths == nil {
		for _, p := range paths {
			if weight := network.TotalWeightInPath(p); weight == 0 {
				t.Errorf("Network:TotalWeightInPath: weight can not be zero")
			}
		}
	}
}

// TestNetwork_BestPathFromNodeToNode is ...
func TestNetwork_BestPathFromNodeToNode(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	root := grafo.NewNode("root/1", grafo.NewNodeContent("root/1", 0))
	node1 := grafo.NewNode("node/1", grafo.NewNodeContent("node/1", 2))
	node2 := grafo.NewNode("node/2", grafo.NewNodeContent("node/2", 3))
	node3 := grafo.NewNode("node/3", grafo.NewNodeContent("node/3", 4))
	node4 := grafo.NewNode("node/4", grafo.NewNodeContent("node/4", 4))
	network.AddNode(nil, grafo.NodeToLeaf(root), 0)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node1), 10)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node2), 5)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node2), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node3), grafo.NodeToLeaf(node4), 1)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node4), 1)
	if best, weight := network.BestPathFromNodeToNode(root, node4); best == nil {
		t.Errorf("Network:BestPathFromNodeToNode: no best path found from: %#v to %#v\n", root, node3)
	} else {
		if len(best.Branches) != 2 {
			t.Errorf("Network:BestPathFromNodeToNode: best length mismatch: exp: %d got: %d\n", 2, len(best.Branches))
		}
		if weight != 17 {
			t.Errorf("Network:BestPathFromNodeToNode: weight mismatch: exp: %d got: %d\n", 17, weight)
		}
	}
}

// TestNetwork_FindLoops is ...
func TestNetwork_FindLoops(t *testing.T) {
	network := grafo.NewNetwork("network/1")
	root := grafo.NewNode("root/1", grafo.NewNodeContent("root/1", 0))
	node1 := grafo.NewNode("node/1", grafo.NewNodeContent("node/1", 2))
	node2 := grafo.NewNode("node/2", grafo.NewNodeContent("node/2", 3))
	node3 := grafo.NewNode("node/3", grafo.NewNodeContent("node/3", 4))
	node4 := grafo.NewNode("node/4", grafo.NewNodeContent("node/4", 4))
	node5 := grafo.NewNode("node/5", grafo.NewNodeContent("node/5", 5))
	node6 := grafo.NewNode("node/6", grafo.NewNodeContent("node/6", 6))
	network.AddNode(nil, grafo.NodeToLeaf(root), 0)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node1), 10)
	network.AddNode(grafo.NodeToLeaf(root), grafo.NodeToLeaf(node2), 5)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node2), grafo.NodeToLeaf(node3), 1)
	network.AddNode(grafo.NodeToLeaf(node3), grafo.NodeToLeaf(node4), 1)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node4), 1)
	network.AddNode(grafo.NodeToLeaf(node1), grafo.NodeToLeaf(node5), 1)
	network.AddNode(grafo.NodeToLeaf(node3), grafo.NodeToLeaf(root), 1)
	network.AddNode(grafo.NodeToLeaf(node4), grafo.NodeToLeaf(root), 1)
	network.AddNode(grafo.NodeToLeaf(node5), grafo.NodeToLeaf(node6), 1)
	//tools.ToDisplay("%s\n", network.ToMermaid())
	if loops := network.FindLoops(node5, nil); len(loops) != 0 {
		t.Errorf("Network:FindLoops: loops found mismatch: exp: %d got: %d\n", 0, len(loops))
	}
	if loops := network.FindLoops(root, nil); len(loops) != 5 {
		t.Errorf("Network:FindLoops: loops found mismatch: exp: %d got: %d\n", 5, len(loops))
	}
	//for _, loop := range loops {
	//    tools.ToDisplay("%s\n", loop)
	//}
}
