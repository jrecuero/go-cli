package dbase_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestTableLayout_NewTableLayout is ...
func TestTableLayout_NewTableLayout(t *testing.T) {
	tl := dbase.NewTableLayout()
	if len(tl.Columns) != 0 {
		t.Errorf("NewTableLayout: length mismatch: got: %d exp: 0", len(tl.Columns))
	}
	if len(tl.ColumnMap) != 0 {
		t.Errorf("NewTableLayout: length mismatch: got: %d exp: 0", len(tl.ColumnMap))
	}
}

// TestTableLayout_AddColumn is ...
func TestTableLayout_AddColumn(t *testing.T) {
	tl := dbase.NewTableLayout()
	tl.AddColumn(dbase.NewColumn("c1", "string"))
	if len(tl.Columns) != 1 {
		t.Errorf("AddColumn: length mismatch: got: %d exp: 1", len(tl.Columns))
	}
	if len(tl.ColumnMap) != 1 {
		t.Errorf("AddColumn: length mismatch: got: %d exp: 1", len(tl.ColumnMap))
	}
	if tl.Columns[0].Name != "c1" {
		t.Errorf("AddColumn: data mismatch: got: %#v exp: c1", tl.Columns[0].Name)
	}
	tl.AddColumn(dbase.NewColumn("c2", "int"))
	if len(tl.Columns) != 2 {
		t.Errorf("AddColumn: length mismatch: got: %d exp: 2", len(tl.Columns))
	}
	if len(tl.ColumnMap) != 2 {
		t.Errorf("AddColumn: length mismatch: got: %d exp: 2", len(tl.ColumnMap))
	}
	if tl.Columns[1].Name != "c2" {
		t.Errorf("AddColumn: data mismatch: got: %#v exp: c2", tl.Columns[0].Name)
	}
}

// TestTableLayout_GetColumn is ...
func TestTableLayout_GetColumn(t *testing.T) {
	tl := dbase.NewTableLayout()
	exp := dbase.NewColumn("c1", "string")
	tl.AddColumn(exp)
	tl.AddColumn(dbase.NewColumn("c2", "int"))
	got := tl.GetColumn("c1")
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("GetColumn output mismatch:\nexpected: %#v\ngot: %#v\n", exp, got)
	}
}
