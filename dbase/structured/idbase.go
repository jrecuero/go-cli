package structured

// DbData represents the interface for any data in the database.
type DbData interface {
	Get() interface{}
	ToString() string
}

// Dbase represents the interface for database operations.
type Dbase interface {
	CreateTable(tbname string, layout []Layout) bool
	AddRow(tbname string, entry DbData) (int, bool)
	GetRow(tbname string, key int) (DbData, bool)
}
