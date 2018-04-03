package dbase

// DbData represents the interface for any data in the database.
type DbData interface {
	Get() interface{}
}

// Layout represents the layout for a database table.
type Layout struct {
	ColName string
	ColType string
}

// Dbase represents the interface for database operations.
type Dbase interface {
	CreateTable(tbname string) bool
	AddRow(tbname string, entry DbData) (int, bool)
	GetRow(tbname string, key int) (DbData, bool)
}

// Table represents the database table.
type Table struct {
	Name   string
	Layout []Layout
	Data   []DbData
}

// NewTable creates a new table.
func NewTable(name string, layout []Layout) *Table {
	return &Table{
		Name:   name,
		Layout: layout,
	}
}

// DB represents the database
type DB struct {
	Name   string
	Tables map[string]*Table
}

// NewDB creates a new database instance.
func NewDB(dbname string) *DB {
	return &DB{
		Name:   dbname,
		Tables: make(map[string]*Table),
	}
}

// FindTable returns the table for the given name.
func (db *DB) FindTable(tbname string) (*Table, bool) {
	for k, tb := range db.Tables {
		if k == tbname {
			return tb, true
		}
	}
	return nil, false
}

// CreateTable creates a new table.
func (db *DB) CreateTable(tbname string) bool {
	if _, result := db.FindTable(tbname); result {
		return false
	}
	tb := NewTable(tbname, PersonLayout())
	db.Tables[tbname] = tb
	return true
}

// AddRow adds a new row to the table.
func (db *DB) AddRow(tbname string, entry DbData) (int, bool) {
	if tb, result := db.FindTable(tbname); result {
		index := len(tb.Data)
		tb.Data = append(tb.Data, entry)
		return index, true
	}
	return -1, false
}

// GetRow gets a row from the table.
func (db *DB) GetRow(tbname string, key int) (DbData, bool) {
	if tb, result := db.FindTable(tbname); result {
		if key < len(tb.Data) {
			return tb.Data[key], true
		}
	}
	return nil, false
}
