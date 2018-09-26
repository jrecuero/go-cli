package freeway_test

import (
	"testing"

	"github.com/jrecuero/go-cli/app/freeway"
)

func TestSection_Section(t *testing.T) {
	section := freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil)
	if section == nil {
		t.Errorf("NewSection: null pointer\n")
	}
	if section.GetLen() != 100 {
		t.Errorf("NewSection: lenght  mismatch: exp: %d got: %d\n", 100, section.GetLen())
	}
	if section.GetWidth() != 1 {
		t.Errorf("NewSection: width  mismatch: exp: %d got: %d\n", 1, section.GetWidth())
	}
	if section.GetSpec() != freeway.Straight {
		t.Errorf("NewSection: spec  mismatch: exp: %s got: %s\n", freeway.Straight, section.GetSpec())
	}
	if section.Traversing() != 1 {
		t.Errorf("NewSection: traversing mismatch: exp: %d got: %d\n", 1, section.Traversing())
	}
	if section.Entering(10) != 1 {
		t.Errorf("NewSection: entering mismatch: exp: %d got: %d\n", 1, section.Entering(10))
	}
	if section.Exiting(10) != 1 {
		t.Errorf("NewSection: exiting mismatch: exp: %d got: %d\n", 1, section.Exiting(10))
	}
	if section.String() != "length/width: 100/1 spec: Straight\n" {
		t.Errorf("NewSection: string mismatch: %#v\n", section.String())
	}
}

func TestSection_Traversing(t *testing.T) {
	section := freeway.NewSection(100, 1, freeway.Straight,
		func(freeway.Spec) int { return 101 },
		func(int) int { return 102 },
		func(int) int { return 103 })
	if section.Traversing() != 101 {
		t.Errorf("NewSection: traversing mismatch: exp: %d got: %d\n", 101, section.Traversing())
	}
	if section.Entering(10) != 102 {
		t.Errorf("NewSection: entering mismatch: exp: %d got: %d\n", 102, section.Entering(10))
	}
	if section.Exiting(10) != 103 {
		t.Errorf("NewSection: exiting mismatch: exp: %d got: %d\n", 103, section.Exiting(10))
	}
}

func TestQSection_QSection(t *testing.T) {
	section := freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil)
	qs := freeway.NewQSection(section)
	if qs.GetSection().(*freeway.Section) != section {
		t.Errorf("NewQSection: section mismatch: exp: %v got: %v\n", section, qs.GetSection())
	}
}
