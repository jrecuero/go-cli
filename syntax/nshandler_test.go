package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestNSHandler_NSHandler ensures the namespace handler struct works properly.
func TestNSHandler_NSHandler(t *testing.T) {
	h := syntax.NewNSHandler()

	if len(h.GetNameSpaces()) != 0 {
		t.Errorf("namespaces: len")
	}

	_, ok := h.FindNameSpace("test")
	if ok == nil {
		t.Errorf("found namespace")
	}

	var ns *syntax.NameSpace
	ns, ok = h.CreateNameSpace("test")
	if ok != nil {
		t.Errorf("create namespace: return")
	}
	if ns == nil {
		t.Errorf("create namespace: ns")
	}
	newNS, ok := h.FindNameSpace("test")
	if ok != nil {
		t.Errorf("find amespace: not found")
	}
	if newNS != ns {
		t.Errorf("find namespace: different")
	}

	delNS, ok := h.DeleteNameSpace("test")
	if ok != nil {
		t.Errorf("delete amespace: not found")
	}
	if delNS != ns {
		t.Errorf("delete namespace: different")
	}

	c := &syntax.Command{}
	_, ok = h.CreateNameSpace("test1")
	ok = h.RegisterCommandToNameSpace("test1", c)
	if ok != nil {
		t.Errorf("register command: failed")
	}
	ok = h.UnregisterCommandFromNameSpace("test1", c)
	if ok != nil {
		t.Errorf("unregister command: failed")
	}
}

// TestNSHandler_Activate ensures activating namespaces works fine.
func TestNSHandler_Activate(t *testing.T) {
	h := syntax.NewNSHandler()
	ns1, ok := h.CreateNameSpace("test1")

	ans, ok := h.ActivateNameSpace("test1")
	if ok != nil {
		t.Errorf("activate namespace: failed")
	}
	if ans != ns1 {
		t.Errorf("activate namespace: different")
	}
	if h.GetActive() == nil {
		t.Errorf("activate namespace: active nil")
	}
	if h.GetActive().Name != "test1" {
		t.Errorf("activate namespace: active name")
	}
	if h.GetActive().NS != ns1 {
		t.Errorf("activate namespace: active ns")
	}

	dns, ok := h.DeactivateNameSpace("test1")
	if ok != nil {
		t.Errorf("deactivate namespace: failed")
	}
	if dns != ns1 {
		t.Errorf("deactivate namespace: different")
	}
	if h.GetActive() != nil {
		t.Errorf("deactivate namespace: active nil")
	}
}

// TestNSHandler_Switch ensures switching namespaces works fine.
func TestNSHandler_Switch(t *testing.T) {
	h := syntax.NewNSHandler()
	ns1, ok := h.CreateNameSpace("test1")
	ns2, ok := h.CreateNameSpace("test2")
	_, ok = h.ActivateNameSpace("test1")
	sns1, sns2, ok := h.SwitchToNameSpace("test2")
	if ok != nil {
		t.Errorf("switch namespace: failed")
	}
	if sns1 != ns1 {
		t.Errorf("switch namespace: old different")
	}
	if sns2 != ns2 {
		t.Errorf("switch namespace: new different")
	}
	if h.GetActive() == nil {
		t.Errorf("switch namespace: active nil")
	}
	if h.GetActive().Name != "test2" {
		t.Errorf("switch namespace: active name")
	}
	if h.GetActive().NS != ns2 {
		t.Errorf("switch namespace: active ns")
	}
	if len(h.GetStack()) != 1 {
		t.Errorf("switch namespace: stack size")
	}

	bns2, bns1, ok := h.SwitchBackToNameSpace()
	if ok != nil {
		t.Errorf("switch-back namespace: failed")
	}
	if bns1 != ns1 {
		t.Errorf("switch-back namespace: new different")
	}
	if bns2 != ns2 {
		t.Errorf("switch-back namespace: old different")
	}
	if h.GetActive() == nil {
		t.Errorf("switch-back namespace: active nil")
	}
	if h.GetActive().Name != "test1" {
		t.Errorf("switch-back namespace: active name: %s <> %s", h.GetActive().Name, "test1")
	}
	if h.GetActive().NS != ns1 {
		t.Errorf("switch-back namespace: active ns")
	}
	if len(h.GetStack()) != 0 {
		t.Errorf("switch-back namespace: stack size")
	}
}
