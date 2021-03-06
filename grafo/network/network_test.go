package network_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
	"github.com/jrecuero/go-cli/grafo/network"
)

// TestNetwork_Network is ...
func TestNetwork_Network(t *testing.T) {
	net := network.NewNetwork("network/1")
	if net == nil {
		t.Errorf("Network:NewNetwork: network can not be <nil>")
	}
	root := grafo.NewVertex("origin/1")
	node1 := grafo.NewVertex("node/1")
	node2 := grafo.NewVertex("node/2")
	node3 := grafo.NewVertex("node/3")
	node4 := grafo.NewVertex("node/4")
	node5 := grafo.NewVertex("node/5")
	if err := net.AddNode(nil, root, 0); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := net.AddNode(root, node1, 10); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := net.AddNode(root, node2, 5); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := net.AddNode(node2, node3, 3); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := net.AddNode(node2, node4, 6); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := net.AddNode(node1, node5, 1); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := net.AddNode(node3, node5, 7); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	if err := net.AddNode(node4, node5, 11); err != nil {
		t.Errorf("Network:AddNode: return error code: %#v\n", err)
	}
	//tools.ToDisplay("%s\n", net.ToMermaid())
}

// TestNetwork_CostToNode is ...
func TestNetwork_CostToNode(t *testing.T) {
	net := network.NewNetwork("network/1")
	root := network.NewNode("origin/1", network.NewNodeContent("origin/1", 0))
	node1 := network.NewNode("node/1", network.NewNodeContent("node/1", 2))
	node2 := network.NewNode("node/2", network.NewNodeContent("node/2", 3))
	node3 := network.NewNode("node/3", network.NewNodeContent("node/3", 4))
	net.AddNode(nil, network.NodeToVertex(root), 0)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node1), 10)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node2), 5)
	net.AddNode(network.NodeToVertex(node2), network.NodeToVertex(node3), 1)
	net.SetAnchorTo(network.NodeToVertex(root))
	if w, ok := net.CostToNode(node1); ok {
		if w != 12 {
			t.Errorf("Network:CostToNode: incorrect weight mismatch: exp: %d got %d\n", 12, w)
		}
	} else {
		t.Errorf("Network:CostToNode: edge not available from: %#v to %#v\n", root.Label, node1.Label)
	}
	if _, ok := net.CostToNode(node3); ok {
		t.Errorf("Network:CostToNode: edge available from: %#v to %#v\n", root.Label, node3.Label)
	}
}

// TestNetwork_PathsFromNodeToNode is ...
func TestNetwork_PathsFromNodeToNode(t *testing.T) {
	net := network.NewNetwork("network/1")
	root := network.NewNode("root/1", network.NewNodeContent("root/1", 0))
	node1 := network.NewNode("node/1", network.NewNodeContent("node/1", 2))
	node2 := network.NewNode("node/2", network.NewNodeContent("node/2", 3))
	node3 := network.NewNode("node/3", network.NewNodeContent("node/3", 4))
	node4 := network.NewNode("node/4", network.NewNodeContent("node/4", 4))
	node5 := network.NewNode("node/5", network.NewNodeContent("node/5", 5))
	net.AddNode(nil, network.NodeToVertex(root), 0)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node1), 10)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node2), 5)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node2), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node3), network.NodeToVertex(node4), 1)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node4), 1)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node5), 1)
	if paths := net.PathsFromNodeToNode(root, node3); paths == nil {
		t.Errorf("Network:PathsFromNodeToNode: no paths found from: %#v to %#v\n", root.Label, node3.Label)
	} else {
		if len(paths) != 2 {
			t.Errorf("Network:PathsFromNodeToNode: paths length mismatch: exp: %d got: %d\n", 2, len(paths))
		}
	}
	if paths := net.PathsFromNodeToNode(root, node4); paths == nil {
		t.Errorf("Network:PathsFromNodeToNode: no paths found from: %#v to %#v\n", root.Label, node4.Label)
	} else {
		if len(paths) != 3 {
			t.Errorf("Network:PathsFromNodeToNode: paths length mismatch: exp: %d got: %d\n", 3, len(paths))
		}
		//for _, p := range paths {
		//    tools.ToDisplay("%s\n", p)
		//}
	}
	if paths := net.PathsFromNodeToNode(node2, node5); paths != nil {
		t.Errorf("Network:PathsFromNodeToNode: paths found from: %#v to %#v\n", node2.Label, node5.Label)
	}

	net.AddNode(network.NodeToVertex(node3), network.NodeToVertex(root), 1)
	//tools.ToDisplay("%s\n", net.ToMermaid())
	if paths := net.PathsFromNodeToNode(node1, root); paths == nil {
		t.Errorf("Network:PathsFromNodeToNode: no paths found from: %#v to %#v\n", node1.Label, root.Label)
	} else {
		if len(paths) != 1 {
			t.Errorf("Network:PathsFromNodeToNode: paths length mismatch: exp: %d got: %d\n", 1, len(paths))
		}
	}
}

// TestNetwork_TotalWeightInPath is ...
func TestNetwork_TotalWeightInPath(t *testing.T) {
	net := network.NewNetwork("network/1")
	root := network.NewNode("root/1", network.NewNodeContent("root/1", 0))
	node1 := network.NewNode("node/1", network.NewNodeContent("node/1", 1))
	node2 := network.NewNode("node/2", network.NewNodeContent("node/2", 2))
	node3 := network.NewNode("node/3", network.NewNodeContent("node/3", 3))
	node4 := network.NewNode("node/4", network.NewNodeContent("node/4", 4))
	net.AddNode(nil, network.NodeToVertex(root), 0)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node1), 10)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node2), 5)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node2), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node3), network.NodeToVertex(node4), 1)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node4), 1)
	if paths := net.PathsFromNodeToNode(root, node4); paths == nil {
		for _, p := range paths {
			if weight := net.TotalWeightInPath(p); weight == 0 {
				t.Errorf("Network:TotalWeightInPath: weight can not be zero")
			}
		}
	}
}

// TestNetwork_BestPathFromNodeToNode is ...
func TestNetwork_BestPathFromNodeToNode(t *testing.T) {
	net := network.NewNetwork("network/1")
	root := network.NewNode("root/1", network.NewNodeContent("root/1", 0))
	node1 := network.NewNode("node/1", network.NewNodeContent("node/1", 2))
	node2 := network.NewNode("node/2", network.NewNodeContent("node/2", 3))
	node3 := network.NewNode("node/3", network.NewNodeContent("node/3", 4))
	node4 := network.NewNode("node/4", network.NewNodeContent("node/4", 4))
	net.AddNode(nil, network.NodeToVertex(root), 0)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node1), 10)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node2), 5)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node2), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node3), network.NodeToVertex(node4), 1)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node4), 1)
	if best, weight := net.BestPathFromNodeToNode(root, node4); best == nil {
		t.Errorf("Network:BestPathFromNodeToNode: no best path found from: %#v to %#v\n", root, node3)
	} else {
		if len(best.Edges) != 2 {
			t.Errorf("Network:BestPathFromNodeToNode: best length mismatch: exp: %d got: %d\n", 2, len(best.Edges))
		}
		if weight != 17 {
			t.Errorf("Network:BestPathFromNodeToNode: weight mismatch: exp: %d got: %d\n", 17, weight)
		}
	}
}

// TestNetwork_FindLoops is ...
func TestNetwork_FindLoops(t *testing.T) {
	net := network.NewNetwork("network/1")
	root := network.NewNode("root/1", network.NewNodeContent("root/1", 0))
	node1 := network.NewNode("node/1", network.NewNodeContent("node/1", 2))
	node2 := network.NewNode("node/2", network.NewNodeContent("node/2", 3))
	node3 := network.NewNode("node/3", network.NewNodeContent("node/3", 4))
	node4 := network.NewNode("node/4", network.NewNodeContent("node/4", 4))
	node5 := network.NewNode("node/5", network.NewNodeContent("node/5", 5))
	node6 := network.NewNode("node/6", network.NewNodeContent("node/6", 6))
	net.AddNode(nil, network.NodeToVertex(root), 0)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node1), 10)
	net.AddNode(network.NodeToVertex(root), network.NodeToVertex(node2), 5)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node2), network.NodeToVertex(node3), 1)
	net.AddNode(network.NodeToVertex(node3), network.NodeToVertex(node4), 1)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node4), 1)
	net.AddNode(network.NodeToVertex(node1), network.NodeToVertex(node5), 1)
	net.AddNode(network.NodeToVertex(node3), network.NodeToVertex(root), 1)
	net.AddNode(network.NodeToVertex(node4), network.NodeToVertex(root), 1)
	net.AddNode(network.NodeToVertex(node5), network.NodeToVertex(node6), 1)
	//tools.ToDisplay("%s\n", net.ToMermaid())
	if loops := net.FindLoops(node5, nil); len(loops) != 0 {
		t.Errorf("Network:FindLoops: loops found mismatch: exp: %d got: %d\n", 0, len(loops))
	}
	if loops := net.FindLoops(root, nil); len(loops) != 5 {
		t.Errorf("Network:FindLoops: loops found mismatch: exp: %d got: %d\n", 5, len(loops))
	}
	//for _, loop := range loops {
	//    tools.ToDisplay("%s\n", loop)
	//}
}
