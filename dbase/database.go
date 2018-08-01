package dbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/jrecuero/go-cli/tools"
)

// Match represents any match entry in a filter. Any Match is composed for a
// kye, which will be usd to identify the data to search, an operand, a value
// to match and a link with further Matches instances.
type Match struct {
	Key     string
	Operand string
	Value   interface{}
	Link    string
}

// IsContinue returns if the match should allows to continue matching further
// Match instances.
func (m *Match) IsContinue(match bool) bool {
	var conti bool
	if m.Link == "AND" && match {
		conti = true
	} else if m.Link == "AND" && !match {
		conti = false
	} else if m.Link == "OR" {
		conti = true
	} else {
		panic(fmt.Sprintf("unknown link %#v", m.Link))
	}
	return conti
}

// IsMatch returns if there is a match and if the Match allows to continue
// matching further.
func (m *Match) IsMatch(value interface{}) (bool, bool) {
	var match bool
	switch m.Operand {
	case "EQUAL":
		match = value == m.Value
		break
	case "NOT EQUAL":
		match = value != m.Value
		break
	default:
		panic(fmt.Sprintf("unknown operand %#v", m.Operand))
	}
	return match, m.IsContinue(match)
}

// NewMatch returns a new Match instance.
func NewMatch(key string, operand string, value interface{}, link string) *Match {
	return &Match{
		Key:     key,
		Operand: operand,
		Value:   value,
		Link:    link,
	}
}

// FilterGetter represents any method that returns a value ot be used in a
// filter for the given key passed. The key is the value from the Match
// instance, and the return should feed Match.IsMatch call.
type FilterGetter func(key string) interface{}

// Filter represents a filter, which is composed by a list of possible matches.
type Filter struct {
	Matches []*Match
}

// Add is ...
func (f *Filter) Add(matches ...*Match) *Filter {
	for _, match := range matches {
		f.Matches = append(f.Matches, match)
	}
	return f
}

// IsMatch check if the filter matches for the given getter method. The getter
// will provide the value to feed for every March.IsMatch in the filter.
func (f *Filter) IsMatch(getter FilterGetter) bool {
	var gotMatch bool
	var conti bool
	var matchResult bool
	for _, match := range f.Matches {
		value := getter(match.Key)
		gotMatch, conti = match.IsMatch(value)
		if !gotMatch && !conti {
			return false
		}
		matchResult = matchResult || gotMatch
	}
	return matchResult
}

// NewFilter returns a new Filter instance.
func NewFilter(matches ...*Match) *Filter {
	f := &Filter{}
	return f.Add(matches...)
}

// Row represents data stored in any row. Data can be from any type so it is
// stored in a generic way as an array of interface{}.
type Row struct {
	Data []interface{}
}

// NewRow creates a new Row instance. Data from every column is provided in a
// variadic way.
func NewRow(cols ...interface{}) *Row {
	row := &Row{}
	for _, col := range cols {
		row.Data = append(row.Data, col)
	}
	return row
}

// Column represents information required to identify any column in the table.
// Some of the information is related with the data and some other information
// is related with the column layout to be displayed if required.
type Column struct {
	Name  string
	Label string
	CType string
	Width int
	Align string
	Key   bool
}

// SetLayout assgins values to the column layout attributes.
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

// NewColumn creates a new Column instance. Only data related attributes are
// being initialized at this stage.
func NewColumn(name string, ctype string) *Column {
	return &Column{
		Name:  name,
		CType: ctype,
	}
}

// TableLayout represents the layout for any table. Layout is basically a
// collection of columns.
type TableLayout struct {
	Columns   []*Column
	ColumnMap map[string]*Column
}

// AddColumn adds column information to the layout in a variadic way.
func (tl *TableLayout) AddColumn(cols ...*Column) *TableLayout {
	for _, col := range cols {
		tl.Columns = append(tl.Columns, col)
		tl.ColumnMap[col.Name] = col
	}
	return tl
}

// GetColumn retrieves the column instance for the given column name.
func (tl *TableLayout) GetColumn(colname string) *Column {
	return tl.ColumnMap[colname]
}

// NewTableLayout creates a new TableLayout instance.
func NewTableLayout() *TableLayout {
	return &TableLayout{
		ColumnMap: make(map[string]*Column),
	}
}

// Table represents table information and table data. Table should contain
// information about the layout and all rows containing the table data. The
// layout is required in order to be able to process row data properly.
type Table struct {
	Name        string
	Layout      *TableLayout
	Rows        []*Row
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
		tb.Rows = append(tb.Rows, row)
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
	return tb.Rows[irow]
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
		dummy := reflect.New(reflect.ValueOf(tb.Stut).Elem().Type()).Interface()
		json.Unmarshal(rowbyte, &dummy)
		return dummy
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
	tb.Rows[irow] = row
	return tb
}

// UpdateColByIndexInRow updates the given column (by index) at the given row
// with the given value.
func (tb *Table) UpdateColByIndexInRow(irow int, icol int, colvalue interface{}) *Table {
	row := tb.GetRow(irow)
	row.Data[icol] = colvalue
	tb.Rows[irow] = row
	return tb
}

// UpdateColByNameInRow updates the given column (by name) at the given row
// with the given value.
func (tb *Table) UpdateColByNameInRow(irow int, colname string, colvalue interface{}) *Table {
	row := tb.GetRow(irow)
	icol := tb.GetColumnIndex(colname)
	row.Data[icol] = colvalue
	tb.Rows[irow] = row
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

// DataBase represents the database information, contaiing all table data.
type DataBase struct {
	Name     string
	Tables   []*Table
	TableMap map[string]*Table
}

// AddTable adds tables to the database in a variadic way.
func (db *DataBase) AddTable(tables ...*Table) *DataBase {
	for _, tb := range tables {
		db.Tables = append(db.Tables, tb)
		db.TableMap[tb.Name] = tb
	}
	return db
}

// GetTable retrieves a  table instance for the given table name.
func (db *DataBase) GetTable(tbname string) *Table {
	return db.TableMap[tbname]
}

// Save saves all daabase information in a file.
// Database filename is set to <database-name>.db
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

// NewDataBase creates a new DataBase instance.
func NewDataBase(name string) *DataBase {
	return &DataBase{
		Name:     name,
		TableMap: make(map[string]*Table),
	}
}

// Load returns a new DataBase instance with the conntent from a file tha
// should contains a save database output.
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

// TableSearch represents a search for a table. Search is defined by a filter.
type TableSearch struct {
	table  *Table
	filter *Filter
}

// Search returns all rows in the table that match the search filter.
func (tbsearch *TableSearch) Search() []*Row {
	rows := []*Row{}
	for _, row := range tbsearch.table.Rows {
		if tbsearch.filter.IsMatch(func(key string) interface{} {
			icol := tbsearch.table.GetColumnIndex(key)
			return row.Data[icol]
		}) {
			rows = append(rows, row)
		}
	}
	return rows
}

// NewTableSearch returns a new TableSearch instance.
func NewTableSearch(table *Table, filter *Filter) *TableSearch {
	return &TableSearch{
		table:  table,
		filter: filter,
	}
}
