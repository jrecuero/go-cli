package techs

import "github.com/jrecuero/go-cli/app/code/battle"

func createStanceKarateEastHigh(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "High"
	stanceHigh := battle.NewStance(stanceName, style)
	return stanceHigh
}

func createStanceKarateEastLow(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Low"
	stanceLow := battle.NewStance(stanceName, style)
	return stanceLow
}

func createStyleKarateEast(tech battle.ITechnique, args ...interface{}) battle.IStyle {
	styleName := "East"
	styleEast := battle.NewStyle(styleName, tech)
	createStanceKarateEastHigh(styleEast, args...)
	createStanceKarateEastLow(styleEast, args...)
	return styleEast
}

func createStanceRelax(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Relax"
	stanceRelax := battle.NewStance(stanceName, style)
	return stanceRelax
}

func createStyleKarateWest(tech battle.ITechnique, args ...interface{}) battle.IStyle {
	styleName := "West"
	styleWest := battle.NewStyle(styleName, tech)
	createStanceRelax(styleWest, args...)
	return styleWest
}

// CreateTechKarate is ...
func CreateTechKarate(args ...interface{}) (string, battle.TechniqueBuilderCb) {
	techName := "Karate"
	return techName, func(...interface{}) battle.ITechnique {
		karate := battle.NewTechnique(techName)
		createStyleKarateEast(karate, args...)
		createStyleKarateWest(karate, args...)
		return karate
	}
}

func createStanceBoxeoWestTense(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Tense"
	stanceTense := battle.NewStance(stanceName, style)
	return stanceTense
}

func createStyleBoxeoWest(tech battle.ITechnique, args ...interface{}) battle.IStyle {
	styleName := "West"
	styleWest := battle.NewStyle(styleName, tech)
	createStanceBoxeoWestTense(styleWest, args...)
	return styleWest
}

// CreateTechBoxeo is ...
func CreateTechBoxeo(args ...interface{}) (string, battle.TechniqueBuilderCb) {
	techName := "Boxeo"
	return techName, func(...interface{}) battle.ITechnique {
		boxeo := battle.NewTechnique(techName)
		createStyleBoxeoWest(boxeo, args...)
		return boxeo
	}
}
