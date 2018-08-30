package dbase

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/jrecuero/go-cli/tools"
)

// Table represents table information and table data. Table should contain
// information about the layout and all rows containing the table data. The
// layout is required in order to be able to process row data properly.
type Table struct {
	Name        string
	Layout      *TableLayout
	Rows        []*TableRow
	ColumnIndex map[string]int
	Stut        interface{}
}

// SetLayout sets the given TableLayout to the table.
func (tb *Table) SetLayout(layout *TableLayout) *Table {
	tb.Layout = layout
	for index, col := range layout.Columns {
		tb.ColumnIndex[col.Name] = index
	}
	return tb
}

// AddRow adds rows to the table in a variadic way.
func (tb *Table) AddRow(rows ...*Row) *Table {
	for _, row := range rows {
		tb.Rows = append(tb.Rows, NewTableRowFromRow(row))
	}
	return tb
}

// GetColumnIndex retrieves the column position in the table for the given
// column name.
func (tb *Table) GetColumnIndex(colname string) int {
	return tb.ColumnIndex[colname]
}

// GetColumnFromName retrieves the column instance for the given column name.
func (tb *Table) GetColumnFromName(colname string) *Column {
	return tb.Layout.GetColumn(colname)
}

// GetColumnFromIndex retrieves the column instance placed in the given
// position,
func (tb *Table) GetColumnFromIndex(icol int) *Column {
	return tb.Layout.Columns[icol]
}

// GetRow retrieves the row instance placed in the given position.
func (tb *Table) GetRow(irow int) *Row {
	return tb.Rows[irow].Row
}

// GetRowAsByteArray retrieves the row information as a byte array. Byte array
// is formated in a JSON format, so it can be unmarshaled into a golang struct.
func (tb *Table) GetRowAsByteArray(irow int) []byte {
	row := tb.GetRow(irow)
	result := []byte("{")
	for i, field := range row.Data {
		if i != 0 {
			result = append(result, []byte(",")...)
		}
		column := tb.GetColumnFromIndex(i)
		entry := fmt.Sprintf("%#v:%#v", column.Name, field)
		result = append(result, []byte(entry)...)
	}
	result = append(result, []byte("}")...)
	return result
}

// GetRowAsStruct retrieves the row information as a golang structure. The
// structure is the one provided for the table layout.
func (tb *Table) GetRowAsStruct(irow int) interface{} {
	if tb.Stut != nil {
		rowbyte := tb.GetRowAsByteArray(irow)
		result := reflect.New(reflect.ValueOf(tb.Stut).Elem().Type()).Interface()
		json.Unmarshal(rowbyte, &result)
		return result
	}
	return nil
}

// SetLayoutFromStruct creates a table layout from a golang struct. Information
// is passed as field tags:
//
// json: column name. Use the the actual struct field name if not provided.
// label: column label. Used for displaying column header. Use column name if not
// present.
// type: column type. Use the field type.
// width: column width. Used for diplaying column width.
// align: column alignment. Used for displaying columnn alignment.
// kye: column key. Identifies if column is a key.
func (tb *Table) SetLayoutFromStruct(in interface{}) ([]*Column, error) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(in)
	// Get the type and kind of our "in" variable
	if t.Kind() == reflect.Ptr {
		t = reflect.Indirect(reflect.ValueOf(in)).Type()
	}
	layout := NewTableLayout()
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		//if field.PkgPath == ""
		col := NewColumn(tools.GetString(field.Tag.Get("json"), field.Name), field.Type.Name())
		col.Width, _ = strconv.Atoi(field.Tag.Get("width"))
		col.Align = tools.GetString(field.Tag.Get("align"), "LEFT")
		col.Key = (field.Tag.Get("key") == "true")
		col.Label = tools.GetString(field.Tag.Get("label"), col.Name)
		layout.AddColumn(col)
	}
	tb.Stut = in
	tb.SetLayout(layout)
	return tb.Layout.Columns, nil
}

// UpdateRow updates table with the given row at the given index.
func (tb *Table) UpdateRow(irow int, row *Row) *Table {
	tb.Rows[irow] = NewTableRowFromRow(row)
	return tb
}

// UpdateColByIndexInRow updates the given column (by index) at the given row
// with the given value.
func (tb *Table) UpdateColByIndexInRow(irow int, icol int, colvalue interface{}) *Table {
	tb.Rows[irow].Row.Data[icol] = colvalue
	return tb
}

// UpdateColByNameInRow updates the given column (by name) at the given row
// with the given value.
func (tb *Table) UpdateColByNameInRow(irow int, colname string, colvalue interface{}) *Table {
	icol := tb.GetColumnIndex(colname)
	tb.Rows[irow].Row.Data[icol] = colvalue
	return tb
}

// GetColByIndexInRow return the column value for the given row for the given
// column index.
func (tb *Table) GetColByIndexInRow(irow int, icol int) interface{} {
	row := tb.GetRow(irow)
	return row.Data[icol]
}

// GetColByNameInRow returns the column value for the given row for the given
// column name.
func (tb *Table) GetColByNameInRow(irow int, colname string) interface{} {
	row := tb.GetRow(irow)
	icol := tb.GetColumnIndex(colname)
	return row.Data[icol]
}

// DeleteRow deletes the row at the given index.
func (tb *Table) DeleteRow(irow int) *Table {
	rowLen := len(tb.Rows)
	tb.Rows = append(tb.Rows[0:irow], tb.Rows[irow+1:rowLen]...)
	return tb
}

// NewTable creates a new Table instance.
func NewTable(name string) *Table {
	return &Table{
		Name:        name,
		ColumnIndex: make(map[string]int),
	}
}
