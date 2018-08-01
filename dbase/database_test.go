package dbase_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/dbase"
)

// TestRow_NewRow is ...
func TestRow_NewRow(t *testing.T) {
	var row *dbase.Row
	row = dbase.NewRow("a")
	if len(row.Data) != 1 {
		t.Errorf("NewRow: length mismatch: got: %d exp: 1", len(row.Data))
	}
	if row.Data[0] != "a" {
		t.Errorf("NewRow: data mismatch: got: %#v exp: \"a\"", row.Data[0])
	}

	row = dbase.NewRow(1, 2, 3, 4, 5)
	if len(row.Data) != 5 {
		t.Errorf("NewRow: length mismatch: got: %d exp: 5", len(row.Data))
	}
	for i := 0; i < 5; i++ {
		if row.Data[i] != i+1 {
			t.Errorf("NewRow: data mismatch: got: %#v exp: %d", row.Data[i], i+1)
		}
	}
}

// TestColumn_NewColumn is ...
func TestColumn_NewColumn(t *testing.T) {
	var tests = []struct {
		input *dbase.Column
		exp   *dbase.Column
	}{
		{
			input: &dbase.Column{
				Name:  "test",
				CType: "string",
			},
			exp: &dbase.Column{
				Name:  "test",
				CType: "string",
				Label: "",
				Align: "",
				Width: 0,
				Key:   false,
			},
		},
	}
	for i, tt := range tests {
		out := dbase.NewColumn(tt.input.Name, tt.input.CType)
		if !reflect.DeepEqual(tt.exp, out) {
			t.Errorf("%d. NewColumn output mismatch:\ninput:\t%#v\nexp:\t%#v\ngot:\t%#v\n", i, tt.input, tt.exp, out)
		}
	}
}

// TestColumn_SetLayout is ...
func TestColumn_SetLayout(t *testing.T) {
	var tests = []struct {
		input *dbase.Column
		exp   *dbase.Column
	}{
		{
			input: &dbase.Column{
				Name:  "test",
				CType: "string",
				Label: "TEST",
				Align: "RIGHT",
				Width: 10,
			},
			exp: &dbase.Column{
				Name:  "test",
				CType: "string",
				Label: "TEST",
				Align: "RIGHT",
				Width: 10,
				Key:   false,
			},
		},
	}
	for i, tt := range tests {
		out := dbase.NewColumn(tt.input.Name, tt.input.CType)
		out = out.SetLayout(tt.input.Label, tt.input.Width, tt.input.Align)
		if !reflect.DeepEqual(tt.exp, out) {
			t.Errorf("%d. SetLayout  output mismatch:\ninput:\t%#v\nexp:\t%#v\ngot:\t%#v\n", i, tt.input, tt.exp, out)
		}
	}
}

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
