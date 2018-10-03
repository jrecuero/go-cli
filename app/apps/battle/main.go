package main

import (
	"github.com/jrecuero/go-cli/app/apps/battle/commands"
	"github.com/jrecuero/go-cli/app/apps/battle/techs"
	"github.com/jrecuero/go-cli/app/code/battle"
	"github.com/jrecuero/go-cli/prompter"
	"github.com/jrecuero/go-cli/tools"
)

func main() {
	bt := battle.NewBattle()
	bt.AddTechBuilder(battle.NewTechniqueBuilder(techs.CreateTechKarate()))
	bt.AddTechBuilder(battle.NewTechniqueBuilder(techs.CreateTechBoxeo()))
	pr := prompter.NewPrompter()
	pr.Setup("battle", commands.SetupCommands(bt)...)
	tools.ToDisplay("\n")
	tools.ToDisplay("Battle System\n")
	tools.ToDisplay("-------------\n")
	pr.Run()
}
