package main

import (
	"fmt"

	"github.com/jrecuero/go-cli/dbase"
)

func main() {
	db := dbase.NewDB("database-test")
	fmt.Printf("Database %s has been created.\n", db.Name)
	tbname := "Person"
	db.CreateTable(tbname)
	var p = dbase.NewPerson("jose carlos", 51)
	key, result := db.AddRow(tbname, p)
	if !result {
		fmt.Printf("AddRow error for table %s data %v\n", tbname, p)
	}
	data, result := db.GetRow(tbname, key)
	if !result {
		fmt.Printf("GetRow error for table %s key %v\n", tbname, key)
	}
	fmt.Printf("Table: %s, Row nbr %d is %v\n", tbname, key, data)
	tb, _ := db.FindTable(tbname)
	fmt.Printf("Table: %s has this layout: %v\n", tb.Name, tb.Layout)
}
