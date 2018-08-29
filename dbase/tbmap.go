package dbase

// TableMap represents ...
type TableMap struct {
	Map map[string]*Table
}

// NewTableMap is ...
func NewTableMap() *TableMap {
	return &TableMap{
		Map: make(map[string]*Table),
	}
}
