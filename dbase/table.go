package dbase

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
	columnMap map[string]*Column
}

// AddColumn is ...
func (tl *TableLayout) AddColumn(cols ...*Column) *TableLayout {
	for _, col := range cols {
		tl.Columns = append(tl.Columns, col)
		tl.columnMap[col.Name] = col
	}
	return tl
}

// GetColumn is ...
func (tl *TableLayout) GetColumn(colname string) *Column {
	return tl.columnMap[colname]
}

// NewTableLayout is ...
func NewTableLayout() *TableLayout {
	return &TableLayout{
		columnMap: make(map[string]*Column),
	}
}

// Table represents ...
type Table struct {
	Name        string
	layout      *TableLayout
	Rows        []*Row
	columnIndex map[string]int
}

// SetLayout is ...
func (tb *Table) SetLayout(layout *TableLayout) *Table {
	tb.layout = layout
	for index, col := range layout.Columns {
		tb.columnIndex[col.Name] = index
	}
	return tb
}

// GetLayout is ...
func (tb *Table) GetLayout() *TableLayout {
	return tb.layout
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
	return tb.columnIndex[colname]
}

// GetColumnFromName is ...
func (tb *Table) GetColumnFromName(colname string) *Column {
	return tb.layout.GetColumn(colname)
}

// GetColumnFromIndex is ...
func (tb *Table) GetColumnFromIndex(icol int) *Column {
	return tb.layout.Columns[icol]
}

// GetRow is ...
func (tb *Table) GetRow(irow int) *Row {
	return tb.Rows[irow]
}

// NewTable is ...
func NewTable(name string) *Table {
	return &Table{
		Name:        name,
		columnIndex: make(map[string]int),
	}
}

// DataBase represents ...
type DataBase struct {
	Name     string
	Tables   []*Table
	tableMap map[string]*Table
}

// AddTable is ...
func (db *DataBase) AddTable(tables ...*Table) *DataBase {
	for _, tb := range tables {
		db.Tables = append(db.Tables, tb)
		db.tableMap[tb.Name] = tb
	}
	return db
}

// GetTable ...
func (db *DataBase) GetTable(tbname string) *Table {
	return db.tableMap[tbname]
}

// NewDataBase is ...
func NewDataBase(name string) *DataBase {
	return &DataBase{
		Name:     name,
		tableMap: make(map[string]*Table),
	}
}
