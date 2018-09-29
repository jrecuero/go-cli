package commands

import (
	"github.com/jrecuero/go-cli/app/code/battle"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// SetupCommands is ...
func SetupCommands(bt *battle.Battle) []*syntax.Command {
	displayCommand := syntax.NewCommand(
		nil,
		"display",
		"Display ...",
		nil,
		nil)

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

	createActorCommand := syntax.NewCommand(
		createCommand,
		"actor name",
		"Create a new actor",
		[]*syntax.Argument{
			syntax.NewArgument(
				"name",
				"Actor name",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	createActorCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		name := params["name"].(string)
		bt.AddActor(battle.NewActor(name))
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
				nil,
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

	cmds := []*syntax.Command{
		displayCommand,
		displayTechsCommand,
		displayActorCommand,
		createCommand,
		createActorCommand,
		addTechToActorCommand,
	}
	return cmds
}
