package dbase_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestTable_NewTable is ...
func TestTable_NewTable(t *testing.T) {
	var tests = []struct {
		input *dbase.Table
		exp   *dbase.Table
	}{
		{
			input: &dbase.Table{
				Name: "test-table",
			},
			exp: &dbase.Table{
				Name:        "test-table",
				ColumnIndex: make(map[string]int),
			},
		},
	}
	for i, tt := range tests {
		out := dbase.NewTable(tt.input.Name)
		if !reflect.DeepEqual(tt.exp, out) {
			t.Errorf("%d. NewTable output mismatch:\ninput:\t%#v\nexp:\t%#v\ngot:\t%#v\n", i, tt.input, tt.exp, out)
		}
	}
}

// TestTable_SetLayout is ...
func TestTable_SetLayout(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tl := dbase.NewTableLayout()
	tl.AddColumn(dbase.NewColumn("c1", "int"), dbase.NewColumn("c2", "int"))
	tb.SetLayout(tl)
	if tb.Layout == nil {
		t.Errorf("SetLayout: layout mistmatch: layout\n")
	}
	if len(tb.ColumnIndex) != 2 {
		t.Errorf("SetLayout: column index mistmatch: len:  got: %d exp: 2\n", len(tb.ColumnIndex))
	}
	if tb.ColumnIndex["c1"] != 0 {
		t.Errorf("SetLayout: column index mistmatch: data: got %d exp 0\n", tb.ColumnIndex["c1"])
	}
	if tb.ColumnIndex["c2"] != 1 {
		t.Errorf("SetLayout: column index mistmatch: data: got %d exp 1\n", tb.ColumnIndex["c2"])
	}
}

// TestTable_AddRow is ...
func TestTable_AddRow(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tl := dbase.NewTableLayout()
	tl.AddColumn(dbase.NewColumn("c1", "string"), dbase.NewColumn("c2", "string"))
	tb.SetLayout(tl)
	if tb.AddRow(dbase.NewRow("a", "b"), dbase.NewRow("c", "d")) == nil {
		t.Errorf("AddRow: return mismatch: <nil>\n")
	}
	if len(tb.Rows) != 2 {
		t.Errorf("AddRow: len mismatch: got: %d exp: 2\n", len(tb.Rows))
	}
}

// TestTable_GetColumn is ...
func TestTable_GetColumn(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tl := dbase.NewTableLayout()
	c1 := dbase.NewColumn("c1", "string")
	c2 := dbase.NewColumn("c2", "string")
	tl.AddColumn(c1, c2)
	tb.SetLayout(tl)
	if result := tb.GetColumnIndex("c1"); result != 0 {
		t.Errorf("GetColumnIndex: result mistmatch: got: %d exp: 0\n", result)
	}
	if col := tb.GetColumnFromName("c2"); !reflect.DeepEqual(col, c2) {
		t.Errorf("GetColumnFromName: result mistmatch: got: %#v exp: %#v", col, c2)
	}
	if col := tb.GetColumnFromIndex(0); !reflect.DeepEqual(col, c1) {
		t.Errorf("GetColumnFromName: result mistmatch: got: %#v exp: %#v", col, c1)
	}
}

// TestTable_GetRow is ...
func TestTable_GetRow(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tl := dbase.NewTableLayout()
	tl.AddColumn(dbase.NewColumn("c1", "string"), dbase.NewColumn("c2", "string"))
	tb.SetLayout(tl)
	r1 := dbase.NewRow("a", "b")
	r2 := dbase.NewRow("c", "d")
	tb.AddRow(r1, r2)
	if got := tb.GetRow(0); !reflect.DeepEqual(r1, got) {
		t.Errorf("GetRow output mismatch:\nexpected: %#v\ngot: %#v\n", r1, got)
	}
	if got := tb.GetRow(1); !reflect.DeepEqual(r2, got) {
		t.Errorf("GetRow: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", r2, got)
	}
}

// TestTable_GetRowAsByteArray
func TestTable_GetRowAsByteArray(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tl := dbase.NewTableLayout()
	tl.AddColumn(dbase.NewColumn("c1", "string"), dbase.NewColumn("c2", "string"))
	tb.SetLayout(tl)
	r1 := dbase.NewRow("a", "b")
	r2 := dbase.NewRow("c", "d")
	tb.AddRow(r1, r2)
	exp := []byte(`{"c1":"a","c2":"b"}`)
	if got := tb.GetRowAsByteArray(0); !reflect.DeepEqual(exp, got) {
		t.Errorf("GetRowAsByteArray: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", string(exp), string(got))
	}
}

type Person struct {
	Name string
	Age  int
}

// TestTable_SetLayoutFromStruct is ...
func TestTable_SetLayoutFromStruct(t *testing.T) {
	tb := dbase.NewTable("test-table")
	cols, err := tb.SetLayoutFromStruct(&Person{})
	if err != nil {
		t.Errorf("SetLayoutFromStruct: layout mistmatch: return\n")
	}
	if len(cols) != 2 {
		t.Errorf("SetLayoutFromStruct: columns returnes: len:  got: %d exp: 2\n", len(cols))
	}
	exp := []*dbase.Column{
		&dbase.Column{
			Name:  "Name",
			CType: "string",
			Align: "LEFT",
			Key:   false,
			Label: "Name",
			Width: 0,
		},
		&dbase.Column{
			Name:  "Age",
			CType: "int",
			Align: "LEFT",
			Key:   false,
			Label: "Age",
			Width: 0,
		},
	}
	if !reflect.DeepEqual(exp[0], cols[0]) {
		t.Errorf("SetLayoutFromStuct: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[0], cols[0])
	}
	if !reflect.DeepEqual(exp[1], cols[1]) {
		t.Errorf("SetLayoutFromStuct: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[1], cols[1])
	}
	if tb.Layout == nil {
		t.Errorf("SetLayoutFromStruct: layout mistmatch: layout\n")
	}
	if len(tb.ColumnIndex) != 2 {
		t.Errorf("SetLayoutFromStruct: column index mistmatch: len:  got: %d exp: 2\n", len(tb.ColumnIndex))
	}
	if tb.ColumnIndex["Name"] != 0 {
		t.Errorf("SetLayoutFromStruct: column index mistmatch: data: got %d exp 0\n", tb.ColumnIndex["Name"])
	}
	if tb.ColumnIndex["Age"] != 1 {
		t.Errorf("SetLayoutFromStruct: column index mistmatch: data: got %d exp 1\n", tb.ColumnIndex["Age"])
	}
}

type PersonTag struct {
	Name string `json:"name" label:"TagName" width:"16" align:"RIGHT"`
	Age  int
}

// TestTable_SetLayoutFromStruct_Tag is ...
func TestTable_SetLayoutFromStruct_Tag(t *testing.T) {
	tb := dbase.NewTable("test-table")
	cols, err := tb.SetLayoutFromStruct(&PersonTag{})
	if err != nil {
		t.Errorf("SetLayoutFromStruct: layout mistmatch: return\n")
	}
	if len(cols) != 2 {
		t.Errorf("SetLayoutFromStruct: columns returnes: len:  got: %d exp: 2\n", len(cols))
	}
	exp := []*dbase.Column{
		&dbase.Column{
			Name:  "name",
			CType: "string",
			Align: "RIGHT",
			Key:   false,
			Label: "TagName",
			Width: 16,
		},
		&dbase.Column{
			Name:  "Age",
			CType: "int",
			Align: "LEFT",
			Key:   false,
			Label: "Age",
			Width: 0,
		},
	}
	if !reflect.DeepEqual(exp[0], cols[0]) {
		t.Errorf("SetLayoutFromStuct: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[0], cols[0])
	}
	if !reflect.DeepEqual(exp[1], cols[1]) {
		t.Errorf("SetLayoutFromStuct: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp[1], cols[1])
	}
	if tb.Layout == nil {
		t.Errorf("SetLayoutFromStruct: layout mistmatch: layout\n")
	}
	if len(tb.ColumnIndex) != 2 {
		t.Errorf("SetLayoutFromStruct: column index mistmatch: len:  got: %d exp: 2\n", len(tb.ColumnIndex))
	}
	if tb.ColumnIndex["name"] != 0 {
		t.Errorf("SetLayoutFromStruct: column index mistmatch: data: got %d exp 0\n", tb.ColumnIndex["name"])
	}
	if tb.ColumnIndex["Age"] != 1 {
		t.Errorf("SetLayoutFromStruct: column index mistmatch: data: got %d exp 1\n", tb.ColumnIndex["Age"])
	}
}

// TestTable_GetRowAsStruct is ...
func TestTable_GetRowAsStruct(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tb.SetLayoutFromStruct(&Person{})
	r := dbase.NewRow("josecarlos", 51)
	tb.AddRow(r)
	exp := &Person{
		Name: "josecarlos",
		Age:  51,
	}
	if got := tb.GetRowAsStruct(0); !reflect.DeepEqual(exp, got) {
		t.Errorf("SetRowAsStruct: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp, got)
	}
}

// TestTable_UpdateRow is ...
func TestTable_UpdateRow(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tb.SetLayoutFromStruct(&Person{})
	tb.AddRow(dbase.NewRow("josecarlos", 51))
	r := dbase.NewRow("jc", 51)
	tb.UpdateRow(0, r)
	exp := &Person{
		Name: "jc",
		Age:  51,
	}
	if got := tb.GetRowAsStruct(0); !reflect.DeepEqual(exp, got) {
		t.Errorf("SetRowAsStruct: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp, got)
	}
}

// TestTable_UpdateCol is ...
func TestTable_UpateCol(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tb.SetLayoutFromStruct(&Person{})
	tb.AddRow(dbase.NewRow("josecarlos", 51))
	if tb.UpdateColByIndexInRow(0, 1, 50) == nil {
		t.Errorf("UpdateColByIndex: return mismatch: <nil>\n")
	}
	exp := &Person{
		Name: "josecarlos",
		Age:  50,
	}
	if got := tb.GetRowAsStruct(0); !reflect.DeepEqual(exp, got) {
		t.Errorf("UpdateColByIndex: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp, got)
	}

	if tb.UpdateColByNameInRow(0, "Name", "JOSE CARLOS") == nil {
		t.Errorf("UpdateColByIndex: return mismatch: <nil>\n")
	}
	exp = &Person{
		Name: "JOSE CARLOS",
		Age:  50,
	}
	if got := tb.GetRowAsStruct(0); !reflect.DeepEqual(exp, got) {
		t.Errorf("UpdateColByName: output mismatch:\nexp:\t%#v\ngot:\t%#v\n", exp, got)
	}
}

// TestTable_GetCol is ...
func TestTable_GetCol(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tb.SetLayoutFromStruct(&Person{})
	tb.AddRow(dbase.NewRow("josecarlos", 51))
	if got := tb.GetColByIndexInRow(0, 1); got != 51 {
		t.Errorf("GetColByIndexInRow: return mismatch: exp: 51 got: %#v\n", got)
	}
	if got := tb.GetColByNameInRow(0, "Name"); got != "josecarlos" {
		t.Errorf("GetColByNameInRow: return mismatch: exp: josecarlos got: %#v\n", got)
	}
}

// TestTable_DeleteRow is ...
func TestTable_DeleteRow(t *testing.T) {
	tb := dbase.NewTable("test-table")
	tb.SetLayoutFromStruct(&Person{})
	tb.AddRow(dbase.NewRow("john", 20))
	tb.AddRow(dbase.NewRow("ann", 25))
	tb.AddRow(dbase.NewRow("david", 30))
	tb.AddRow(dbase.NewRow("mary", 27))
	tb.AddRow(dbase.NewRow("frank", 35))
	if tb.DeleteRow(2) == nil {
		t.Errorf("DeleteRow: return mismatch: <nil>\n")
	}
	if len(tb.Rows) != 4 {
		t.Errorf("DeleteRow: len mistmatch: exp:4 got: %d\n", len(tb.Rows))
	}
}
