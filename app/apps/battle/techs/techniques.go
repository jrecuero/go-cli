package techs

import "github.com/jrecuero/go-cli/app/code/battle"

// CreateTechKarate is ...
func CreateTechKarate(args ...interface{}) (string, string, battle.TechniqueBuilderCb) {
	techName := "Karate"
	techDesc := "Karate technique"
	return techName, techDesc, func(...interface{}) battle.ITechnique {
		karate := battle.NewTechnique(techName)
		karate.SetDescription(techDesc)
		// Styles
		styleKarateOne := battle.NewStyle("One", karate)
		styleKarateTwo := battle.NewStyle("Two", karate)
		// Stances
		stanceKarateOneAggressive := battle.NewStance("Aggressive", styleKarateOne)
		stanceKarateOneProtective := battle.NewStance("Protective", styleKarateOne)
		stanceKarateTwoRelax := battle.NewStance("Relax", styleKarateTwo)
		// Amoves
		battle.NewAmove("Punch", stanceKarateOneAggressive).GetUpdateStats().Set(battle.StatStr, func(stat battle.TStat, actor battle.IActor, args ...interface{}) battle.TStat {
			return stat * 1.2
		})
		battle.NewAmove("Block", stanceKarateOneProtective)
		battle.NewAmove("Taunt", stanceKarateTwoRelax)
		return karate
	}
}

// CreateTechBoxeo is ...
func CreateTechBoxeo(args ...interface{}) (string, string, battle.TechniqueBuilderCb) {
	techName := "Boxeo"
	techDesc := "Boxeo technique"
	return techName, techDesc, func(...interface{}) battle.ITechnique {
		boxeo := battle.NewTechnique(techName)
		boxeo.SetDescription(techDesc)
		// Styles
		styleBoxeoClassic := battle.NewStyle("Classic", boxeo)
		// Stances
		stanceBoxeoClassicBasic := battle.NewStance("Basic", styleBoxeoClassic)
		stanceBoxeoClassicBasic.GetUpdateStats().Set(battle.StatSta, func(stat battle.TStat, actor battle.IActor, args ...interface{}) battle.TStat {
			return stat * 1.25
		})
		// Amoves
		battle.NewAmove("Block", stanceBoxeoClassicBasic).GetUpdateStats().Set(battle.StatSta, func(stat battle.TStat, actor battle.IActor, args ...interface{}) battle.TStat {
			return stat * 1.5
		})
		battle.NewAmove("UpperCut", stanceBoxeoClassicBasic)
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
		// Styles
		styleKbUFC := battle.NewStyle("UFC", kb)
		// Stances
		stanceKbUFCAggressive := battle.NewStance("Aggressive", styleKbUFC)
		stanceKbUFCPassive := battle.NewStance("Passive", styleKbUFC)
		// Amoves
		battle.NewAmove("Super-Kick", stanceKbUFCAggressive).SetShortName("SKICK")
		battle.NewAmove("Super-Punch", stanceKbUFCAggressive).SetShortName("SPUNCH")
		battle.NewAmove("Guard", stanceKbUFCPassive)
		return kb
	}
}
