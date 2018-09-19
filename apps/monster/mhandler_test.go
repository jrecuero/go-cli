package monster_test

import (
	"testing"
	"time"

	"github.com/jrecuero/go-cli/apps/monster"
	"github.com/jrecuero/go-cli/tools"
)

func TestMonster_Run(t *testing.T) {
	mhandler := monster.NewMHandler()
	mhandler.AddActor(monster.NewActor("actor-1000", 1000))
	mhandler.AddActor(monster.NewActor("actor-490", 490))
	mhandler.AddActor(monster.NewActor("actor-710", 710))
	go func() {
		time.Sleep(2 * 1000 * time.Millisecond)
		mhandler.Stop()
	}()
	mhandler.Start()
	//time.Sleep(1 * 1000 * time.Millisecond)
	tools.ToDisplay("\n\n\nDisplay results\n")
	tools.ToDisplay("play %#v\n", mhandler.Next())
	if actor := mhandler.Next(); actor != nil {
		tools.ToDisplay("play %#v\n", actor)
		for i, actor := range mhandler.GetActors() {
			tools.ToDisplay("[%d] %#v\n", i, actor)
		}
	}
	tools.ToDisplay("\n\n\n\n")

	go func() {
		time.Sleep(2 * 1000 * time.Millisecond)
		mhandler.Stop()
	}()
	mhandler.Start()
	//time.Sleep(1 * 1000 * time.Millisecond)
	tools.ToDisplay("\n\n\nDisplay results\n")
	tools.ToDisplay("play %#v\n", mhandler.Next())
	if actor := mhandler.Next(); actor != nil {
		tools.ToDisplay("play %#v\n", actor)
		for i, actor := range mhandler.GetActors() {
			tools.ToDisplay("[%d] %#v\n", i, actor)
		}
	}
	tools.ToDisplay("\n\n\n\n")
}
