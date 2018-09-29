package techs

import "github.com/jrecuero/go-cli/app/code/battle"

func createHighPunch(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "Punch"
	amovePunch := battle.NewAmove(amoveName, stance)
	return amovePunch
}

func createStanceKarateEastHigh(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "High"
	stanceHigh := battle.NewStance(stanceName, style)
	createHighPunch(stanceHigh, args...)
	return stanceHigh
}

func createLowKick(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "Kick"
	amoveKick := battle.NewAmove(amoveName, stance)
	return amoveKick
}

func createStanceKarateEastLow(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Low"
	stanceLow := battle.NewStance(stanceName, style)
	createLowKick(stanceLow, args...)
	return stanceLow
}

func createStyleKarateEast(tech battle.ITechnique, args ...interface{}) battle.IStyle {
	styleName := "East"
	styleEast := battle.NewStyle(styleName, tech)
	createStanceKarateEastHigh(styleEast, args...)
	createStanceKarateEastLow(styleEast, args...)
	return styleEast
}

func createSoft(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "Soft"
	amoveSoft := battle.NewAmove(amoveName, stance)
	return amoveSoft
}

func createStanceRelax(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Relax"
	stanceRelax := battle.NewStance(stanceName, style)
	createSoft(stanceRelax, args...)
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

func createHammer(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "Hammer"
	amoveHammer := battle.NewAmove(amoveName, stance)
	return amoveHammer
}

func createStanceBoxeoWestTense(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Tense"
	stanceTense := battle.NewStance(stanceName, style)
	createHammer(stanceTense, args...)
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
