package casting

import (
	"github.com/jrecuero/go-cli/app/apps/battle/techs"
	"github.com/jrecuero/go-cli/app/code/battle"
)

// newHero is ...
func newHero(name string, desc string) battle.IActor {
	if desc == "" {
		desc = "This is the hero"
	}
	stats := battle.NewStats()
	stats.Set(battle.StatLix, 100)
	stats.Set(battle.StatStr, 15)
	stats.Set(battle.StatAgi, 20)
	stats.Set(battle.StatSta, 10)
	stats.Set(battle.StatPre, 25)
	stats.Set(battle.StatFoc, 50)
	actor := battle.NewActor(name, desc, stats)
	_, _, techBuilder := techs.CreateTechKickboxing()
	actor.AddTechnique(techBuilder())
	actor.SetStyle(actor.GetStyleByName("UFC"))
	actor.SetStance(actor.GetStanceByName("Aggressive"))
	actor.SetAmove(actor.GetAmoveByName("Super-Kick"))
	actor.SetDefaultAsTechnique()
	return actor
}

// newGansta is ...
func newGansta(name string) battle.IActor {
	return battle.NewActor(name, "This is a gansta!", battle.NewStats())
}

// newBoss is ...
func newBoss(name string) battle.IActor {
	return battle.NewActor(name, "This is a boss!", battle.NewStats())
}

// GetActors is ...
func GetActors() []*battle.ActorInfo {
	return []*battle.ActorInfo{
		battle.NewActorInfo("Hero", "This is the hero"),
		battle.NewActorInfo("Gansta", "This is a gansta!"),
		battle.NewActorInfo("Boss:Joker", "This is Joker Boss"),
		battle.NewActorInfo("Boss:Octopus", "This is Octopus Boss"),
	}
}

// CreateCastingActor is ...
func CreateCastingActor(castingName string, name string, desc string) battle.IActor {
	switch castingName {
	case "Hero":
		return newHero(name, desc)
	case "Gansta":
		return newGansta(name)
	case "Boss:Joker":
		return newBoss("Joker")
	case "Boss:Ocotopus":
		return newBoss("Octopus")
	default:
		return nil
	}
}
