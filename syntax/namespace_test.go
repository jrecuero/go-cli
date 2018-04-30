package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestNameSpace_NameSpace ensures the namepace structure works properly.
func TestNameSpace_NameSpace(t *testing.T) {
	ns := &syntax.NameSpace{
		Name: "test",
	}

	if ns.Name != "test" {
		t.Errorf("NameSpace Name failed: name")
	}
	if len(ns.GetCommands()) != 0 {
		t.Errorf("NameSpace Get Commands failed: len")
	}

	c := &syntax.Command{}
	ns.Add(c)
	if len(ns.GetCommands()) != 1 {
		t.Errorf("NameSpace Add Command failed: len")
	}
	if ns.GetCommands()[0] != c {
		t.Errorf("NameSpace Add Command failed: array")
	}

	index, ok := ns.Find(c)
	if ok != nil {
		t.Errorf("NameSpace Find Command failed: error")
	}
	if index != 0 {
		t.Errorf("NameSpace Find Command failed: index")
	}

	nc, ok := ns.DeleteForIndex(0)
	if ok != nil {
		t.Errorf("NameSpace Delete For Index failed: error")
	}
	if nc != c {
		t.Errorf("NameSpace Delete For Index failed: match")
	}
	if len(ns.GetCommands()) != 0 {
		t.Errorf("NameSpace Delete For Index failed: len")
	}

	ns.Add(c)
	ok = ns.DeleteForCommand(c)
	if ok != nil {
		t.Errorf("NameSpace Delete For Command failed: error")
	}
	if len(ns.GetCommands()) != 0 {
		t.Errorf("NameSpace Delete For Command failed: len")
	}
}
