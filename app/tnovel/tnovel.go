package main

import (
	"github.com/jrecuero/go-cli/app/tnovel/novel"
	"github.com/jrecuero/go-cli/prompter"
	"github.com/jrecuero/go-cli/tools"
)

func run() {
	novel.TheNovel.Run()
	pr := prompter.NewPrompter()
	pr.Setup("prompter", novel.SetupCommands()...)
	tools.ToDisplay("\n")
	tools.ToDisplay("Novel Command Line\n")
	tools.ToDisplay("------------------\n")
	pr.Run()
}

func main() {
	run()
}
