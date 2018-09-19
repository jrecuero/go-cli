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
		time.Sleep(5 * 1000 * time.Millisecond)
		mhandler.Stop()
		for i, actor := range mhandler.GetSchedule() {
			tools.ToDisplay("[%d] %#v\n", i, actor)
		}
	}()
	mhandler.Start()
}
