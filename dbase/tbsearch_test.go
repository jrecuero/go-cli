package dbase_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestTableSearch_NewTableSearch is ...
func TestTableSearch_NewTableSearch(t *testing.T) {
	if got := dbase.NewTableSearch(nil, nil); got == nil {
		t.Errorf("NewTableSearch: return mismatch: <nil")
	}
}

// TestTableSearch_Search is ...
func TestTableSearch_Search(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tb.SetLayoutFromStruct(&Person{})
	tb.AddRow(dbase.NewRow("john", 20))
	tb.AddRow(dbase.NewRow("ann", 25))
	tb.AddRow(dbase.NewRow("david", 30))
	tb.AddRow(dbase.NewRow("mary", 27))
	tb.AddRow(dbase.NewRow("frank", 35))
	tb.AddRow(dbase.NewRow("john", 50))
	if filter := dbase.NewFilter(dbase.NewMatch("Name", "EQUAL", "ann", "AND")); filter != nil {
		if tbs := dbase.NewTableSearch(tb, filter); tbs != nil {
			got := tbs.Search()
			if len(got) != 1 {
				t.Errorf("Search: result mistmatch: exp: 1 got:%d\n", len(got))
			}
			exp := []*dbase.Row{
				dbase.NewRow("ann", 25),
			}
			if !reflect.DeepEqual(exp[0], got[0]) {
				t.Errorf("Search output mismatch:exp:\t%#v\ngot:\t%#v\n", exp[0], got[0])
			}
		}
	}
	if filter := dbase.NewFilter(
		dbase.NewMatch("Name", "EQUAL", "ann", "OR"),
		dbase.NewMatch("Name", "EQUAL", "frank", "OR")); filter != nil {
		if tbs := dbase.NewTableSearch(tb, filter); tbs != nil {
			got := tbs.Search()
			if len(got) != 2 {
				t.Errorf("Search: result mistmatch: exp: 2 got:%d\n", len(got))
			}
			exp := []*dbase.Row{
				dbase.NewRow("ann", 25),
				dbase.NewRow("frank", 35),
			}
			if !reflect.DeepEqual(exp[0], got[0]) {
				t.Errorf("Search output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[0], got[0])
			}
			if !reflect.DeepEqual(exp[1], got[1]) {
				t.Errorf("Search output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[1], got[1])
			}
		}
	}
	if filter := dbase.NewFilter(
		dbase.NewMatch("Name", "EQUAL", "ann", "OR"),
		dbase.NewMatch("Age", "EQUAL", 20, "OR")); filter != nil {
		if tbs := dbase.NewTableSearch(tb, filter); tbs != nil {
			got := tbs.Search()
			if len(got) != 2 {
				t.Errorf("Search: result mistmatch: exp: 2 got:%d\n", len(got))
			}
			exp := []*dbase.Row{
				dbase.NewRow("john", 20),
				dbase.NewRow("ann", 25),
			}
			if !reflect.DeepEqual(exp[0], got[0]) {
				t.Errorf("Search output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[0], got[0])
			}
			if !reflect.DeepEqual(exp[1], got[1]) {
				t.Errorf("Search output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[1], got[1])
			}
		}
	}
	if filter := dbase.NewFilter(
		dbase.NewMatch("Name", "EQUAL", "john", "AND"),
		dbase.NewMatch("Age", "EQUAL", 50, "AND")); filter != nil {
		if tbs := dbase.NewTableSearch(tb, filter); tbs != nil {
			got := tbs.Search()
			if len(got) != 1 {
				t.Errorf("Search: result mistmatch: exp: 1 got:%d\n", len(got))
			}
			exp := []*dbase.Row{
				dbase.NewRow("john", 50),
			}
			if !reflect.DeepEqual(exp[0], got[0]) {
				t.Errorf("Search output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[0], got[0])
			}
		}
	}
}
