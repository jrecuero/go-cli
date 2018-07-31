package dbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
)

// Row represents ...
type Row struct {
	Data []interface{}
}

// NewRow is ...
func NewRow(cols ...interface{}) *Row {
	row := &Row{}
	for _, col := range cols {
		row.Data = append(row.Data, col)
	}
	return row
}

// Column represents ...
type Column struct {
	Name  string
	Label string
	CType string
	Width int
	Align string
	Key   bool
}

// SetLayout is ...
func (col *Column) SetLayout(label interface{}, width int, align string) *Column {
	var _label string
	if label == nil {
		_label = col.Name
	} else {
		_label = label.(string)
	}
	col.Label = _label
	col.Width = width
	col.Align = align
	return col
}

// NewColumn is ...
func NewColumn(name string, ctype string) *Column {
	return &Column{
		Name:  name,
		CType: ctype,
	}
}

// TableLayout represents ...
type TableLayout struct {
	Columns   []*Column
	ColumnMap map[string]*Column
}

// AddColumn is ...
func (tl *TableLayout) AddColumn(cols ...*Column) *TableLayout {
	for _, col := range cols {
		tl.Columns = append(tl.Columns, col)
		tl.ColumnMap[col.Name] = col
	}
	return tl
}

// GetColumn is ...
func (tl *TableLayout) GetColumn(colname string) *Column {
	return tl.ColumnMap[colname]
}

// NewTableLayout is ...
func NewTableLayout() *TableLayout {
	return &TableLayout{
		ColumnMap: make(map[string]*Column),
	}
}

// Table represents ...
type Table struct {
	Name        string
	Layout      *TableLayout
	Rows        []*Row
	ColumnIndex map[string]int
	Stut        interface{}
}

// SetLayout is ...
func (tb *Table) SetLayout(layout *TableLayout) *Table {
	tb.Layout = layout
	for index, col := range layout.Columns {
		tb.ColumnIndex[col.Name] = index
	}
	return tb
}

// AddRow is ...
func (tb *Table) AddRow(rows ...*Row) *Table {
	for _, row := range rows {
		tb.Rows = append(tb.Rows, row)
	}
	return tb
}

// GetColumnIndex is ...
func (tb *Table) GetColumnIndex(colname string) int {
	return tb.ColumnIndex[colname]
}

// GetColumnFromName is ...
func (tb *Table) GetColumnFromName(colname string) *Column {
	return tb.Layout.GetColumn(colname)
}

// GetColumnFromIndex is ...
func (tb *Table) GetColumnFromIndex(icol int) *Column {
	return tb.Layout.Columns[icol]
}

// GetRow is ...
func (tb *Table) GetRow(irow int) *Row {
	return tb.Rows[irow]
}

// GetRowAsByteArray is ...
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

// GetRowAsStruct is ...
func (tb *Table) GetRowAsStruct(irow int) interface{} {
	if tb.Stut != nil {
		rowbyte := tb.GetRowAsByteArray(irow)
		dummy := reflect.New(reflect.ValueOf(tb.Stut).Elem().Type()).Interface()
		json.Unmarshal(rowbyte, &dummy)
		return dummy
	}
	return nil
}

// SetLayoutFromStruct is ...
func (tb *Table) SetLayoutFromStruct(in interface{}) ([]*Column, error) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(in)
	// Get the type and kind of our "in" variable
	if t.Kind() == reflect.Ptr {
		t = reflect.Indirect(reflect.ValueOf(in)).Type()
	}
	tb.Layout = NewTableLayout()
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		//if field.PkgPath == ""
		col := NewColumn(field.Tag.Get("json"), field.Type.Name())
		col.Width, _ = strconv.Atoi(field.Tag.Get("width"))
		col.Align = field.Tag.Get("align")
		if col.Align == "" {
			col.Align = "LEFT"
		}
		col.Key = (field.Tag.Get("key") == "true")
		col.Label = field.Tag.Get("label")
		if col.Label == "" {
			col.Label = col.Name
		}
		tb.Layout.AddColumn(col)
	}
	tb.Stut = in
	return tb.Layout.Columns, nil
}

// NewTable is ...
func NewTable(name string) *Table {
	return &Table{
		Name:        name,
		ColumnIndex: make(map[string]int),
	}
}

// DataBase represents ...
type DataBase struct {
	Name     string
	Tables   []*Table
	TableMap map[string]*Table
}

// AddTable is ...
func (db *DataBase) AddTable(tables ...*Table) *DataBase {
	for _, tb := range tables {
		db.Tables = append(db.Tables, tb)
		db.TableMap[tb.Name] = tb
	}
	return db
}

// GetTable is ...
func (db *DataBase) GetTable(tbname string) *Table {
	return db.TableMap[tbname]
}

// Save is ...
func (db *DataBase) Save() error {
	data, err := json.Marshal(db)
	if err == nil {
		filename := fmt.Sprintf("%s.db", db.Name)
		ferr := ioutil.WriteFile(filename, data, 0666)
		if ferr != nil {
			panic(ferr)
		}
		return nil
	}
	return err
}

// NewDataBase is ...
func NewDataBase(name string) *DataBase {
	return &DataBase{
		Name:     name,
		TableMap: make(map[string]*Table),
	}
}

// Load is ...
func Load(filename string) (*DataBase, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var newDb *DataBase
	if err := json.Unmarshal(data, &newDb); err != nil {
		panic(err)
	}
	return newDb, nil
}
