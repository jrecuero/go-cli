package dbase_test

import (
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestFilter_NewFilter is ...
func TestFilter_NewFilter(t *testing.T) {
	if got := dbase.NewFilter(); true {
		if got == nil {
			t.Errorf("NewFilter: return mismatch: <nil>")
		}
		if len(got.Matches) != 0 {
			t.Errorf("NewFilter: len mismatch: exp: 0 got: %d\n", len(got.Matches))
		}
	}
}

// TestFilter_Add is ...
func TestFilter_Add(t *testing.T) {
	if got := dbase.NewFilter(); got != nil {
		if got.Add(dbase.NewMatch("a", "EQUAL", "b", "AND")) == nil {
			t.Errorf("Add: return mismatch: <nil>\n")
		}
		if len(got.Matches) != 1 {
			t.Errorf("Add: len matches mismatch: exp: 1 got: %d\n", len(got.Matches))
		}
	}
}

// TestFilter_IsMatch is ...
func TestFilter_IsMatch(t *testing.T) {
	if filter := dbase.NewFilter(dbase.NewMatch("name", "EQUAL", "home", "AND")); filter != nil {
		if !filter.IsMatch(func(key string) interface{} {
			return "home"
		}) {
			t.Errorf("IsMatch: return mistmatch: exp: true got:false\n")
		}
	}
	if filter := dbase.NewFilter(
		dbase.NewMatch("name", "EQUAL", "home", "AND"),
		dbase.NewMatch("age", "EQUAL", 30, "AND")); filter != nil {
		if !filter.IsMatch(func(key string) interface{} {
			switch key {
			case "name":
				return "home"
			case "age":
				return 30
			}
			return nil
		}) {
			t.Errorf("IsMatch: return mistmatch: exp: true got:false\n")
		}
	}
}
