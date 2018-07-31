package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jrecuero/go-cli/dbase"
)

// Dbase creates a database interface
type Dbase interface {
	add(entry string) int
	retrieve(key int) string
}

// DB is the database
type DB struct {
	data []string
}

func (db *DB) add(entry string) int {
	return 0
}

func (db *DB) retrieve(key int) string {
	return "none"
}

func add(itf Dbase, entry string) int {
	return itf.add(entry)
}

var mydbase Dbase

type shape interface {
	draw()
	named() string
}

type circle struct {
	x      int
	y      int
	radius int
}

type square struct {
	x    int
	y    int
	side int
}

type mockShape struct {
	shape
	mutex sync.Mutex
}

func newMockShape() shape {
	return &mockShape{}
}

func (m *mockShape) draw() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println("mockShape calling draw")
}

func (m *mockShape) named() string {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println("mockShape calling named")
	return "mockShape:named"
}

func draw(s shape) {
	s.draw()
}

func named(s shape) string {
	return s.named()
}

var _ shape = (*mockShape)(nil)

type context struct {
	s shape
}

type config func() shape

func initContext(configs ...config) *context {
	ctx := &context{}
	for _, c := range configs {
		ctx.s = c()
	}
	return ctx
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func runDataBase() {
	layout := dbase.NewTableLayout()
	layout.AddColumn(dbase.NewColumn("name", "string"), dbase.NewColumn("age", "int"))
	tb := dbase.NewTable("Person").SetLayout(layout)
	db := dbase.NewDataBase("WORK")
	db.AddTable(tb)
	tb.AddRow(dbase.NewRow("jc", "51"), dbase.NewRow("marce", "43"))
	fmt.Printf("Layout: %#v\n", layout)
	fmt.Printf("Table: %#v\n", tb)
	fmt.Printf("DataBase: %#v\n", db)
	fmt.Printf("Row: %#v\n", tb.GetRow(1))
	m := tb.GetRowAsByteArray(1)
	fmt.Printf("Row: %s\n", string(m))
	p := &person{}
	json.Unmarshal(m, &p)
	fmt.Printf("person: %#v\n", p)
	columns, _ := tb.SetLayoutFromStruct(&person{})
	for i, col := range columns {
		fmt.Printf("%d %#v\n", i, col)
	}
	_p := tb.GetRowAsStruct(1).(*person)
	fmt.Printf("p: %#v\n", _p)
	db.Save()
	newDB := dbase.NewDataBase("NEW-WORK")
	fmt.Printf("DataBase: %#v\n", newDB)
	newDB, _ = dbase.Load("WORK.db")
	fmt.Printf("DataBase: %#v\n", newDB)
	fmt.Printf("Table: %#v\n", newDB.Tables[0])
	fmt.Printf("Layout: %#v\n", newDB.Tables[0].Layout)
}

func main() {
	//var x = &DB{}
	//fmt.Println(add((Dbase)(x), "me"))

	//mocka := &mockShape{}
	//draw(mocka)
	//named(mocka)

	//contexto := initContext(newMockShape)
	//draw(contexto.s)
	//named((*contexto).s)

	runDataBase()
}
