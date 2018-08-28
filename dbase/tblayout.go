package dbase

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
