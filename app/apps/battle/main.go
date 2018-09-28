package main

import (
	"fmt"

	"github.com/jrecuero/go-cli/app/code/battle"
)

func testActor() {
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

func createKarate() (string, battle.TechniqueBuilderCb) {
	techName := "Karate"
	return techName, func(...interface{}) battle.ITechnique {
		fmt.Println("techName is: KARATE")
		karate := battle.NewTechnique(techName)
		karateStyleOne := battle.NewStyle("east", karate)
		battle.NewStance("high", karateStyleOne)
		battle.NewStance("low", karateStyleOne)
		karateStyleTwo := battle.NewStyle("west", karate)
		battle.NewStance("high", karateStyleTwo)
		return karate
	}
}

func createBoxeo() (string, battle.TechniqueBuilderCb) {
	techName := "Boxeo"
	return techName, func(...interface{}) battle.ITechnique {
		boxeo := battle.NewTechnique(techName)
		boxeoStyleOne := battle.NewStyle("west", boxeo)
		battle.NewStance("tense", boxeoStyleOne)
		return boxeo
	}
}

func main() {
	bt := battle.NewBattle()
	bt.AddTechBuilder(battle.NewTechniqueBuilder(createKarate()))
	bt.AddTechBuilder(battle.NewTechniqueBuilder(createBoxeo()))
	actor := battle.NewActor("me")
	actor.AddTechnique(bt.CreateTechniqueByName("Karate"))
	bt.AddActor(actor)
	fmt.Printf("%#v\n", actor)
	fmt.Printf("%#v\n", actor.TechniqueHandler)
	for _, tech := range actor.GetTechniques() {
		fmt.Printf("%s\n", tech)
	}
}
