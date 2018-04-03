package main

import (
	"fmt"
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

func main() {
	var x = &DB{}
	fmt.Println(add((Dbase)(x), "me"))
}
