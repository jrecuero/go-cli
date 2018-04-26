package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestContent_Content ensures the content structure works properly.
func TestContent_Content(t *testing.T) {
	c := syntax.NewContent("test", "test help", nil)

	if c.GetLabel() != "test" {
		t.Errorf("GetLabel <Content> failed")
	}
	if c.GetType() != "string" {
		t.Errorf("GetType <Content> failed")
	}
	if c.GetDefault() != "test" {
		t.Errorf("GetDefault <Content> failed")
	}
	if c.GetHelp() != "test help" {
		t.Errorf("GetHelp <Content> failed")
	}
	if c.GetCompleter() != nil {
		t.Errorf("GetCompleter <Content> failed")
	}
}

// TestContent_Joint ensures the content joint structure works properly.
func TestContent_Joint(t *testing.T) {
	c := syntax.NewContentJoint("test joint", "test joint help", nil)
	if c.GetLabel() != "test joint" {
		t.Errorf("GetLabel <ContentJoint> failed")
	}
	if c.GetHelp() != "test joint help" {
		t.Errorf("GetHelp <ContentJoint> failed")
	}
	if c.GetCompleter() != nil {
		t.Errorf("GetCompleter <ContentJoint> failed")
	}
}

// TestContent_CR ensures the content for cr structure works properly.
func TestContent_CR(t *testing.T) {
	c := syntax.GetCR()
	if c.GetLabel() != "<<<_CR_>>>" {
		t.Errorf("GetLabel <CR> failed")
	}
	if c.GetHelp() != "Carrier return" {
		t.Errorf("GetHelp <CR> failed")
	}
	if c.GetCompleter() == nil {
		t.Errorf("GetCompleter <CR> failed")
	}
}
