package main

import (
	"fmt"
)

// DbData represents the interface for any data in the database.
type DbData interface {
	Get() interface{}
}

// Person represents an example for a table layout.
type Person struct {
	Name string
	Age  int
}

// Layout represents the layout for a database table.
type Layout struct {
	ColName string
	ColType string
}

// NewPerson creates a new Person instance.
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// PersonLayout returns the layout for Person.
func PersonLayout() []Layout {
	return []Layout{
		{
			ColName: "name",
			ColType: "string",
		},
		{
			ColName: "age",
			ColType: "int",
		},
	}
}

// Get return Person data.
func (p *Person) Get() interface{} {
	return &Person{
		Name: p.Name,
		Age:  p.Age,
	}
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

func (db *DB) findTable(tbname string) (*Table, bool) {
	for k, tb := range db.Tables {
		if k == tbname {
			return tb, true
		}
	}
	return nil, false
}

// CreateTable creates a new table.
func (db *DB) CreateTable(tbname string) bool {
	if _, result := db.findTable(tbname); result {
		return false
	}
	tb := NewTable(tbname, PersonLayout())
	db.Tables[tbname] = tb
	return true
}

// AddRow adds a new row to the table.
func (db *DB) AddRow(tbname string, entry DbData) (int, bool) {
	if tb, result := db.findTable(tbname); result {
		index := len(tb.Data)
		tb.Data = append(tb.Data, entry)
		return index, true
	}
	return -1, false
}

// GetRow gets a row from the table.
func (db *DB) GetRow(tbname string, key int) (DbData, bool) {
	if tb, result := db.findTable(tbname); result {
		if key < len(tb.Data) {
			return tb.Data[key], true
		}
	}
	return nil, false
}

func main() {
	dbase := NewDB("database-test")
	fmt.Printf("Database %s has been created.\n", dbase.Name)
	tbname := "Person"
	dbase.CreateTable(tbname)
	var p = NewPerson("jose carlos", 51)
	key, result := dbase.AddRow(tbname, p)
	if !result {
		fmt.Printf("AddRow error for table %s data %v\n", tbname, p)
	}
	data, result := dbase.GetRow(tbname, key)
	if !result {
		fmt.Printf("GetRow error for table %s key %v\n", tbname, key)
	}
	fmt.Printf("Table: %s, Row nbr %d is %v\n", tbname, key, data)
	tb, _ := dbase.findTable(tbname)
	fmt.Printf("Table: %s has this layout: %v\n", tb.Name, tb.Layout)
}
