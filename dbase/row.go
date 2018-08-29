package dbase

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

// TableRow represents ...
type TableRow struct {
	*Row
}

// NewTableRow is ...
func NewTableRow(cols ...interface{}) *TableRow {
	return &TableRow{
		NewRow(cols...),
	}
}
