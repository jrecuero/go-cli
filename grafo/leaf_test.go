package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
)

// TestLeaf_NewLeaf is ...
func TestLeaf_NewLeaf(t *testing.T) {
	if leaf := grafo.NewLeaf("leaf/0"); leaf == nil {
		t.Errorf("NewLeaf: leaf can not be <nil>")
	}
}

// TestLeaf_AddBranch is ...
func TestLeaf_AddBranch(t *testing.T) {
	parentLeaf := grafo.NewLeaf("root/1")
	childLeaf := grafo.NewLeaf("child/1")
	branch := grafo.StaticBranch(parentLeaf, childLeaf)
	if err := parentLeaf.AddBranch(branch); err != nil {
		t.Errorf("Leaf:AddBranch: return code error: %#v\n", err)
	}
	if len(parentLeaf.Branches) != 1 {
		t.Errorf("Leaf:AddBranch: branches length mismatch: exp %d got: %d\n", 1, len(parentLeaf.Branches))
	}
}
