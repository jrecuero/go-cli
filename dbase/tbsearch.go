package dbase

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
