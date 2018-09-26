package freeway_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/app/code/freeway"
)

func TestFreeway_Freeway(t *testing.T) {
	fway := freeway.NewFreeway()
	if fway == nil {
		t.Errorf("NewFreeway: null pointer\n")
	}
	if fway.GetLen() != 0 {
		t.Errorf("NewFreeway: section length mismatch: exp %d got: %d\n", 0, fway.GetLen())
	}
}

func TestFreeway_AddSection(t *testing.T) {
	fway := freeway.NewFreeway()
	sections := []*freeway.Section{
		freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil),
		freeway.NewSection(50, 1, freeway.Turn, nil, nil, nil),
		freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil),
		freeway.NewSection(50, 1, freeway.Turn, nil, nil, nil),
	}
	for i, sect := range sections {
		fway.AddSection(sect)
		if fway.GetLen() != i+1 {
			t.Errorf("AddSection: section length mismatch: exp %d got: %d\n", i+1, fway.GetLen())
		}
	}
	for i, got := range sections {
		exp := fway.GetSection(i)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("GetSection: section mismatch:\n\texp: %#v\n\tgot: %#v\n", exp, got)
		}
	}
	sectIndexLimit := fway.GetLen() - 1
	for i := 0; i < sectIndexLimit; i++ {
		isect, lap := fway.NextSectionIndex(i)
		if isect != i+1 {
			t.Errorf("NextSectionIndex: index mismatch: exp: %d got: %d\n", i+1, isect)
		}
		if lap {
			t.Errorf("NextSectionIndex: lap mismatch: exp: %v got: %v\n", false, lap)
		}
	}
	isect, lap := fway.NextSectionIndex(sectIndexLimit)
	if isect != 0 {
		t.Errorf("NextSectionIndex: index mismatch: exp: %d got: %d\n", 0, isect)
	}
	if !lap {
		t.Errorf("NextSectionIndex: lap mismatch: exp: %v got: %v\n", true, lap)
	}
}
