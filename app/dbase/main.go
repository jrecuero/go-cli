package main

import (
	"github.com/jrecuero/go-cli/app/dbase/commands"
	"github.com/jrecuero/go-cli/prompter"
	"github.com/jrecuero/go-cli/tools"
)

func runDbase() {
	pr := prompter.NewPrompter()
	pr.Setup("prompter", commands.SetupCommands()...)
	tools.ToDisplay("\n")
	tools.ToDisplay("DATABASE Command Line\n")
	tools.ToDisplay("---------------------\n")
	pr.Run()
}

func main() {
	runDbase()
}
