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

// TableRow represents the full data related with a row stored in a table. It
// should include required information to be used in any transaction.
type TableRow struct {
	*Row
	updated bool
	added   bool
	deleted bool
	shadow  *Row
}

// CleanPrivate cleans all private attributes used for transaction.
func (tbrow *TableRow) CleanPrivate() {
	tbrow.updated = false
	tbrow.added = false
	tbrow.deleted = false
	tbrow.shadow = nil
}

// SetShadow assigns a new value to the shadow row associated with the row for
// a transacton.
func (tbrow *TableRow) SetShadow(row *Row) {
	tbrow.shadow = row
}

// UpdateWith sets the row as being updated in a trannsaction, with the given
// new row value.
func (tbrow *TableRow) UpdateWith(row *Row) {
	tbrow.updated = true
	tbrow.SetShadow(row)
}

// AddWith sets the row as being added in a transaction with the given new row.
func (tbrow *TableRow) AddWith(row *Row) {
	tbrow.added = true
	tbrow.SetShadow(row)
}

// IsUpdated returns if the row has been updated in a transaction.
func (tbrow *TableRow) IsUpdated() bool {
	return tbrow.updated
}

// IsAdded returns if the row has been added in a transaction.
func (tbrow *TableRow) IsAdded() bool {
	return tbrow.added
}

// IsDeleted returns if the row has been deleted in a transaction.
func (tbrow *TableRow) IsDeleted() bool {
	return tbrow.deleted
}

// IsAnyChange returns if the row has been updated, added or deleted in a
// transaction..
func (tbrow *TableRow) IsAnyChange() bool {
	return tbrow.IsUpdated() || tbrow.IsAdded() || tbrow.IsDeleted()
}

// DeleteWith sets the row as being deleted in a transaction.
func (tbrow *TableRow) DeleteWith() {
	tbrow.deleted = true
	tbrow.SetShadow(nil)
}

// NewTableRow creates a new TableRow instance based on column values.
func NewTableRow(cols ...interface{}) *TableRow {
	return &TableRow{
		Row: NewRow(cols...),
	}
}

// NewTableRowFromRow creates a new TableRow instance based on the given Row
// instance..
func NewTableRowFromRow(row *Row) *TableRow {
	return &TableRow{
		Row: row,
	}
}
