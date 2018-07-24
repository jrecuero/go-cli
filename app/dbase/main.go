package main

import (
	"github.com/jrecuero/go-cli/app/dbase/commands"
	"github.com/jrecuero/go-cli/dbase"
	"github.com/jrecuero/go-cli/prompter"
	"github.com/jrecuero/go-cli/tools"
)

func runDbaseWork() {
	db := dbase.NewDB("database-test")
	tools.ToDisplay("Database %s has been created.\n", db.Name)
	tbname := "Person"
	//db.CreateTable(tbname, dbase.GetLayout(&dbase.Person{}))
	db.CreateTable(tbname, dbase.PersonLayout())
	var p = dbase.NewPerson("jose carlos", 51)
	key, ok := db.AddRow(tbname, p)
	if !ok {
		tools.ToDisplay("AddRow error for table %s data %v\n", tbname, p)
	}
	data, ok := db.GetRow(tbname, key)
	if !ok {
		tools.ToDisplay("GetRow error for table %s key %v\n", tbname, key)
	}
	tools.ToDisplay("Table: %s, Row nbr %d is %v %v %s\n", tbname, key, data, data.Get(), data.ToString())
	tb, _ := db.FindTable(tbname)
	tools.ToDisplay("Table: %s has this layout: %v\n", tb.Name, tb.Layout)
}

func runDbase() {
	pr := prompter.NewPrompter()
	pr.Setup("prompter", commands.SetupCommands()...)
	pr.Run()
}

func main() {
	runDbase()
}
