package main

import (
	"fmt"

	"github.com/jrecuero/go-cli/app/code/battle"
)

func main() {
	actor := battle.NewActor("me")
	actor.AddTechnique(battle.NewTechnique("hit"))
	actor.AddTechnique(battle.NewTechnique("punch"))
	actor.SetTechniqueByName("hit")
	fmt.Printf("%#v\n", actor)
	fmt.Printf("%#v\n", actor.TechniqueHandler)
	for _, tech := range actor.GetTechniques() {
		fmt.Printf("-> %#v\n", tech.GetName())
	}
	fmt.Printf("technique: %#v\n", actor.GetTechnique().GetName())
}
