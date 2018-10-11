package techs

import "github.com/jrecuero/go-cli/app/code/battle"

func createAmoveKarateOneAggressivePunch(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "Punch"
	amovePunch := battle.NewAmove(amoveName, stance)
	return amovePunch
}

func createStanceKarateOneAggressive(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Aggressive"
	stanceHigh := battle.NewStance(stanceName, style)
	createAmoveKarateOneAggressivePunch(stanceHigh, args...)
	return stanceHigh
}

func createAmoveKarateOneProtectiveBlock(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "Block"
	amoveKick := battle.NewAmove(amoveName, stance)
	return amoveKick
}

func createStanceKarateOneProtective(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Protective"
	stanceLow := battle.NewStance(stanceName, style)
	createAmoveKarateOneProtectiveBlock(stanceLow, args...)
	return stanceLow
}

func createStyleKarateOne(tech battle.ITechnique, args ...interface{}) battle.IStyle {
	styleName := "One"
	styleEast := battle.NewStyle(styleName, tech)
	createStanceKarateOneAggressive(styleEast, args...)
	createStanceKarateOneProtective(styleEast, args...)
	return styleEast
}

func createAmoveKarateTwoRelaxTaunt(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "Taunt"
	amoveSoft := battle.NewAmove(amoveName, stance)
	return amoveSoft
}

func createStanceKarateTwoRelax(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Relax"
	stanceRelax := battle.NewStance(stanceName, style)
	createAmoveKarateTwoRelaxTaunt(stanceRelax, args...)
	return stanceRelax
}

func createStyleKarateTwo(tech battle.ITechnique, args ...interface{}) battle.IStyle {
	styleName := "Two"
	styleWest := battle.NewStyle(styleName, tech)
	createStanceKarateTwoRelax(styleWest, args...)
	return styleWest
}

// CreateTechKarate is ...
func CreateTechKarate(args ...interface{}) (string, string, battle.TechniqueBuilderCb) {
	techName := "Karate"
	techDesc := "Karate technique"
	return techName, techDesc, func(...interface{}) battle.ITechnique {
		karate := battle.NewTechnique(techName)
		karate.SetDescription(techDesc)
		createStyleKarateOne(karate, args...)
		createStyleKarateTwo(karate, args...)
		return karate
	}
}

func createAmoveBoxeoClassicBasicUpperCut(stance battle.IStance, args ...interface{}) battle.IAmove {
	amoveName := "UpperCut"
	amoveHammer := battle.NewAmove(amoveName, stance)
	return amoveHammer
}

func createStanceBoxeoClassicBasic(style battle.IStyle, args ...interface{}) battle.IStance {
	stanceName := "Basic"
	stanceTense := battle.NewStance(stanceName, style)
	createAmoveBoxeoClassicBasicUpperCut(stanceTense, args...)
	return stanceTense
}

func createStyleBoxeoClassic(tech battle.ITechnique, args ...interface{}) battle.IStyle {
	styleName := "Classic"
	styleWest := battle.NewStyle(styleName, tech)
	createStanceBoxeoClassicBasic(styleWest, args...)
	return styleWest
}

// CreateTechBoxeo is ...
func CreateTechBoxeo(args ...interface{}) (string, string, battle.TechniqueBuilderCb) {
	techName := "Boxeo"
	techDesc := "Boxeo technique"
	return techName, techDesc, func(...interface{}) battle.ITechnique {
		boxeo := battle.NewTechnique(techName)
		boxeo.SetDescription(techDesc)
		createStyleBoxeoClassic(boxeo, args...)
		return boxeo
	}
}

// CreateTechKickboxing is ...
func CreateTechKickboxing(args ...interface{}) (string, string, battle.TechniqueBuilderCb) {
	techName := "Kickboxing"
	techDesc := "Kickboxing technique"
	return techName, techDesc, func(...interface{}) battle.ITechnique {
		kb := battle.NewTechnique(techName)
		kb.SetDescription(techDesc)
		styleKbUFC := battle.NewStyle("UFC", kb)
		stanceKbUFCAggressive := battle.NewStance("Aggressive", styleKbUFC)
		stanceKbUFCPassive := battle.NewStance("Passive", styleKbUFC)
		battle.NewAmove("Super-Kick", stanceKbUFCAggressive).SetShortName("SKICK")
		battle.NewAmove("Super-Punch", stanceKbUFCAggressive).SetShortName("SPUNCH")
		battle.NewAmove("Guard", stanceKbUFCPassive)
		return kb
	}
}
