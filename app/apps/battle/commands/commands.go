package commands

import (
	"github.com/jrecuero/go-cli/app/code/battle"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// ActorCompleter represents the name completer
type ActorCompleter struct {
	*syntax.CompleterArgument
	bt *battle.Battle
}

// Query returns the query for any node completer.
func (nc *ActorCompleter) Query(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	defaultComplete, _ := nc.Complete(ctx, content, line, index)
	defaultHelp, _ := nc.Help(ctx, content, line, index)
	var data []*syntax.CompleteHelp
	for _, actor := range nc.bt.Actors {
		data = append(data, syntax.NewCompleteHelp(actor.GetName(), actor.GetDescription()))
	}
	data = append(data, syntax.NewCompleteHelp(defaultComplete, defaultHelp))

	//data := []*syntax.CompleteHelp{
	//    syntax.NewCompleteHelp("ME", "ME actor"),
	//    syntax.NewCompleteHelp("YOU", "Other actor"),
	//    syntax.NewCompleteHelp(defaultComplete, defaultHelp),
	//}
	return data, true
}

type selector struct{}

func (s *selector) SelectOrig(...interface{}) battle.IActor {
	return nil
}

func (s *selector) SelectTarget(...interface{}) battle.IActor {
	return nil
}

func (s *selector) SelectAmove(args ...interface{}) battle.IAmove {
	actor := args[0].(battle.IActor)
	mode := args[1].(battle.Amode)
	var amove battle.IAmove
	//tools.ToDisplay("%#v\n", actor.GetTechniques())
	//tools.ToDisplay("%#v\n", actor.GetStyles())
	//tools.ToDisplay("%#v\n", actor.GetStances())
	//tools.ToDisplay("%#v\n", actor.GetAmoves())
	if mode == battle.AmodeAttack {
		amove = actor.GetAmoveByName("Punch")
	} else if mode == battle.AmodeDefence {
		amove = actor.GetAmoveByName("Hammer")
	}
	//tools.ToDisplay("select-amove %s:%#v %#v\n", actor.GetName(), mode, amove)
	return amove
}

// SetupCommands is ...
func SetupCommands(bt *battle.Battle) []*syntax.Command {
	displayCommand := syntax.NewCommand(
		nil,
		"display",
		"Display ...",
		nil,
		nil)
	displayCommand.JustPrefix = true

	displayTechsCommand := syntax.NewCommand(
		displayCommand,
		"techs",
		"Display all available techniques.",
		nil,
		nil)
	displayTechsCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("Display all techs...\n")
		for _, tb := range bt.TechBuilders {
			tech := bt.CreateTechniqueByName(tb.Name)
			tools.ToDisplay("%s\n", tech)
		}
		return nil
	}

	displayActorCommand := syntax.NewCommand(
		displayCommand,
		"actor",
		"Display all available actors",
		nil,
		nil)
	displayActorCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("Display all actors...\n")
		for _, actor := range bt.Actors {
			tools.ToDisplay("-> %s\n", actor.GetName())
			for _, tech := range actor.GetTechniques() {
				tools.ToDisplay("-->%s\n", tech.GetName())
			}
		}
		return nil
	}

	createCommand := syntax.NewCommand(
		nil,
		"create",
		"Create ...",
		nil,
		nil)
	createCommand.JustPrefix = true

	createActorCommand := syntax.NewCommand(
		createCommand,
		"actor name desc",
		"Create a new actor",
		[]*syntax.Argument{
			syntax.NewArgument(
				"name",
				"Actor name",
				nil,
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"desc",
				"Actor description",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	createActorCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		name := params["name"].(string)
		desc := params["desc"].(string)
		bt.AddActor(battle.NewActor(name, desc))
		tools.ToDisplay("Create actor: %#v\n", name)
		return nil
	}

	addTechToActorCommand := syntax.NewCommand(
		nil,
		"add-tech-to-actor actor-name tech-name",
		"Add technique to actor",
		[]*syntax.Argument{
			syntax.NewArgument(
				"actor-name",
				"Actor name",
				&ActorCompleter{syntax.NewCompleterArgument("actor-name"), bt},
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"tech-name",
				"Technique name",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	addTechToActorCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		actorName := params["actor-name"].(string)
		techName := params["tech-name"].(string)
		tech := bt.CreateTechniqueByName(techName)
		actor := bt.GetActorByName(actorName)
		actor.AddTechnique(tech)
		tools.ToDisplay("Add technique %#v to actor %#v\n", techName, actorName)
		return nil
	}

	engageCommand := syntax.NewCommand(
		nil,
		"engage orig-name target-name",
		"Engage in combat",
		[]*syntax.Argument{
			syntax.NewArgument(
				"orig-name",
				"Actor originator name",
				nil,
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"target-name",
				"Actor target name",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	engageCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		origName := params["orig-name"].(string)
		targetName := params["target-name"].(string)
		tools.ToDisplay("Engage %#v vs %#v\n", origName, targetName)
		origActor := bt.GetActorByName(origName)
		targetActor := bt.GetActorByName(targetName)
		if origActor != nil && targetActor != nil {
			bt.Engage(origActor, targetActor)
		}
		return nil
	}

	bt.Selector = &selector{}

	cmds := []*syntax.Command{
		displayCommand,
		displayTechsCommand,
		displayActorCommand,
		createCommand,
		createActorCommand,
		addTechToActorCommand,
		engageCommand,
	}
	return cmds
}
