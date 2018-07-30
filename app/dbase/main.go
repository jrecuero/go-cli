package main

import (
	"github.com/jrecuero/go-cli/app/dbase/commands"
	"github.com/jrecuero/go-cli/prompter"
)

func runDbase() {
	pr := prompter.NewPrompter()
	pr.Setup("prompter", commands.SetupCommands()...)
	pr.Run()
}

func main() {
	runDbase()
}
