package dbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
