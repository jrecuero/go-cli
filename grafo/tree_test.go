package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
)

// TestTree_NewTree is ...
func TestTree_NewTree(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	if tree == nil {
		t.Errorf("NewTree: tree can not be <nil>")
	}
	if tree.Label != "tree/1" {
		t.Errorf("NewTree: label mistmatch: exp: %#v got: %#v\n", "tree/1", tree.Label)
	}
	if tree.GetRoot() == nil {
		t.Errorf("NewTree: root can not be <nil>")
	}
	if tree.GetAnchor() == nil {
		t.Errorf("NewTree: anchor can not be <nil>")
	}
	if tree.GetAnchor() != tree.GetRoot() {
		t.Errorf("NewTree: anchor mismatch: exp: %v got: %v\n", tree.GetRoot(), tree.GetAnchor())
	}
}

// TestTree_AddBranch is ...
func TestTree_AddBranch(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	root := tree.GetRoot()
	parent := grafo.NewLeaf("parent/1")
	rootBranch := grafo.StaticBranch(nil, parent)
	if err := tree.AddBranch(nil, rootBranch); err != nil {
		t.Errorf("AddBranch: return code error: %#v\n", err)
	}
	if rootBranch.Parent == nil {
		t.Errorf("AddBranch: branch parent mismatch: exp not <nil>")
	}
	if rootBranch.Parent != root {
		t.Errorf("AddBranch: branch parent label mismatch: exp: %v got: %v\n", root, rootBranch.Parent)
	}
	if len(parent.Parents) != 1 {
		t.Errorf("Tree:AddBranch: child parents length mismatch: exp %d got: %d\n", 1, len(parent.Parents))
	}
	if parent.Parents[0] != root {
		t.Errorf("Tree:AddBranch: child parent mismatch: exp %v got: %v\n", root, parent.Parents[0])
	}
}

// TestTree_AddChild is ...
func TestTree_AddChild(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	root := tree.GetRoot()
	child1 := grafo.NewLeaf("child/1")
	if err := tree.AddChild(nil, child1); err != nil {
		t.Errorf("Tree:AddChild: return code error: %#v\n", err)
	}
	if len(child1.Parents) != 1 {
		t.Errorf("Tree:AddChild: child parents length mismatch: exp %d got: %d\n", 1, len(child1.Parents))
	}
	if child1.Parents[0] != root {
		t.Errorf("Tree:AddChild: child mismatch: exp %v got: %v\n", root, child1.Parents[0])
	}

	child2 := grafo.NewLeaf("child/2")
	if err := tree.AddChild(nil, child2); err != nil {
		t.Errorf("Tree:AddChild: return code error: %#v\n", err)
	}
	if len(child2.Parents) != 1 {
		t.Errorf("Tree:AddChild: child parents length mismatch: exp %d got: %d\n", 1, len(child2.Parents))
	}
	if child2.Parents[0] != root {
		t.Errorf("Tree:AddChild: child mismatch: exp %v got: %v\n", root, child2.Parents[0])
	}
	if len(root.Branches) != 2 {
		t.Errorf("Tree:AdChild: root branches length mismatch: exp: %d got: %d\n", 2, len(root.Branches))
	}
}

// TestTree_ExistPathTo is ...
func TestTree_ExistPathTo(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	parent := grafo.NewLeaf("parent/1")
	child1 := grafo.NewLeaf("child/1")
	child2 := grafo.NewLeaf("child/2")
	child3 := grafo.NewLeaf("child/3")
	tree.AddChild(nil, parent)
	tree.AddChild(parent, child1)
	tree.AddChild(parent, child2)
	tree.AddChild(child1, child3)
	if branch, ok := tree.ExistPathTo(parent, child1); ok {
		if branch != parent.Branches[0] {
			t.Errorf("Tree:ExistPathTo: incorrect branch: exp: %#v got: %#v\n", parent.Branches[0], branch)
		}
	} else {
		t.Errorf("Tree:ExistPathTo: branch not found from: %#v to %#v\n", parent, child1)
	}
	if _, ok := tree.ExistPathTo(parent, child3); ok {
		t.Errorf("Tree:ExistPathTo: branch found from: %#v to %#v\n", parent, child3)
	}
}

// TestTree_IsPathTo is ...
func TestTree_IsPathTo(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	parent := grafo.NewLeaf("parent/1")
	child1 := grafo.NewLeaf("child/1")
	child2 := grafo.NewLeaf("child/2")
	child3 := grafo.NewLeaf("child/3")
	tree.AddChild(nil, parent)
	tree.AddChild(parent, child1)
	tree.AddChild(parent, child2)
	tree.AddChild(child1, child3)
	if branch, ok := tree.IsPathTo(parent, child1); ok {
		if branch != parent.Branches[0] {
			t.Errorf("Tree:IsPathTo: incorrect branch: exp: %#v got: %#v\n", parent.Branches[0], branch)
		}
	} else {
		t.Errorf("Tree:IsPathTo: branch not found from: %#v to %#v\n", parent, child1)
	}
	if _, ok := tree.IsPathTo(parent, child3); ok {
		t.Errorf("Tree:IsPathTo: branch found from: %#v to %#v\n", parent, child3)
	}
}

// TestTree_PathsFrom is ...
func TestTree_PathsFrom(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	parent := grafo.NewLeaf("parent/1")
	child1 := grafo.NewLeaf("child/1")
	child2 := grafo.NewLeaf("child/2")
	child3 := grafo.NewLeaf("child/3")
	tree.AddChild(nil, parent)
	tree.AddChild(parent, child1)
	tree.AddChild(parent, child2)
	tree.AddChild(child1, child3)
	if paths := tree.PathsFrom(parent); len(paths) != 2 {
		t.Errorf("Tree:PathsFrom: path length mismatch: exp: %d got: %d\n", 2, len(paths))
	}
	if paths := tree.PathsFrom(child3); len(paths) != 0 {
		t.Errorf("Tree:PathsFrom: path length mismatch: exp: %d got: %d\n", 0, len(paths))
	}
}

// TestTree_AddTraverse
func TestTree_AddTraverse(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	parent := grafo.NewLeaf("parent/1")
	child1 := grafo.NewLeaf("child/1")
	child2 := grafo.NewLeaf("child/2")
	tree.AddChild(nil, parent)
	tree.AddChild(parent, child1)
	tree.AddChild(parent, child2)
	if err := tree.AddTraverse(grafo.StaticBranch(nil, parent)); err != nil {
		t.Errorf("Tree:AddTraverse: return error code: %#v\n", err)
	}
}

// TestTree_SetAnchorTo
func TestTree_SetAnchorTo(t *testing.T) {
	tree := grafo.NewTree("tree/1")
	parent := grafo.NewLeaf("parent/1")
	child1 := grafo.NewLeaf("child/1")
	child2 := grafo.NewLeaf("child/2")
	tree.AddChild(nil, parent)
	tree.AddChild(parent, child1)
	tree.AddChild(parent, child2)
	leaf := tree.SetAnchorTo(parent)
	if leaf == nil {
		t.Errorf("Tree:SetAnchorTo: anchor cannot be <nil>")
	}
	if leaf != parent {
		t.Errorf("Tree:SetAnchorTo: anchor mismatch: exp: %#v got: %#v\n", parent, leaf)
	}
}