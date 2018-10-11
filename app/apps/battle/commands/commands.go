package commands

import (
	"github.com/jrecuero/go-cli/app/apps/battle/casting"
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
	return data, true
}

// TechniqueCompleter represents the technique completer.
type TechniqueCompleter struct {
	*syntax.CompleterArgument
	bt         *battle.Battle
	actorLabel string
}

// Query returns the query for any node completer.
func (nc *TechniqueCompleter) Query(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	var data []*syntax.CompleteHelp
	if nc.actorLabel == "" {
		for _, tb := range nc.bt.TechBuilders {
			data = append(data, syntax.NewCompleteHelp(tb.Name, tb.Desc))
		}
	} else {
		if actorName, err := ctx.GetArgValueForArgLabelInMatched(nc.actorLabel); err == nil {
			if actor := nc.bt.GetActorByName(actorName.(string)); actor != nil {
				for _, tech := range actor.GetTechniques() {
					data = append(data, syntax.NewCompleteHelp(tech.GetName(), tech.GetDescription()))
				}

			}
		}
	}
	return data, true
}

// StyleCompleter represents the technique completer.
type StyleCompleter struct {
	*syntax.CompleterArgument
	bt         *battle.Battle
	actorLabel string
	techLabel  string
}

// Query returns the query for any node completer.
func (nc *StyleCompleter) Query(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	var data []*syntax.CompleteHelp
	if actorName, err := ctx.GetArgValueForArgLabelInMatched(nc.actorLabel); err == nil {
		if actor := nc.bt.GetActorByName(actorName.(string)); actor != nil {
			for _, tech := range actor.GetTechniques() {
				if techName, err := ctx.GetArgValueForArgLabelInMatched(nc.techLabel); err == nil {
					if tech.GetName() == techName {
						for _, style := range tech.GetStyles() {
							data = append(data, syntax.NewCompleteHelp(style.GetName(), style.GetDescription()))
						}
					}
				}
			}
		}
	}
	return data, true
}

// StanceCompleter represents the technique completer.
type StanceCompleter struct {
	*syntax.CompleterArgument
	bt         *battle.Battle
	actorLabel string
	techLabel  string
	styleLabel string
}

// Query returns the query for any node completer.
func (nc *StanceCompleter) Query(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	var data []*syntax.CompleteHelp
	if actorName, err := ctx.GetArgValueForArgLabelInMatched(nc.actorLabel); err == nil {
		if actor := nc.bt.GetActorByName(actorName.(string)); actor != nil {
			for _, tech := range actor.GetTechniques() {
				if techName, err := ctx.GetArgValueForArgLabelInMatched(nc.techLabel); err == nil {
					if tech.GetName() == techName {
						for _, style := range tech.GetStyles() {
							if styleName, err := ctx.GetArgValueForArgLabelInMatched(nc.styleLabel); err == nil {
								if style.GetName() == styleName {
									for _, stance := range style.GetStances() {
										data = append(data, syntax.NewCompleteHelp(stance.GetName(), stance.GetDescription()))
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return data, true
}

// AmoveCompleter represents the technique completer.
type AmoveCompleter struct {
	*syntax.CompleterArgument
	bt          *battle.Battle
	actorLabel  string
	techLabel   string
	styleLabel  string
	stanceLabel string
}

// Query returns the query for any node completer.
func (nc *AmoveCompleter) Query(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	var data []*syntax.CompleteHelp
	if actorName, err := ctx.GetArgValueForArgLabelInMatched(nc.actorLabel); err == nil {
		if actor := nc.bt.GetActorByName(actorName.(string)); actor != nil {
			for _, tech := range actor.GetTechniques() {
				if techName, err := ctx.GetArgValueForArgLabelInMatched(nc.techLabel); err == nil {
					if tech.GetName() == techName {
						for _, style := range tech.GetStyles() {
							if styleName, err := ctx.GetArgValueForArgLabelInMatched(nc.styleLabel); err == nil {
								if style.GetName() == styleName {
									for _, stance := range style.GetStances() {
										if stanceName, err := ctx.GetArgValueForArgLabelInMatched(nc.stanceLabel); err == nil {
											if stance.GetName() == stanceName {
												for _, amove := range stance.GetAmoves() {
													data = append(data, syntax.NewCompleteHelp(amove.GetName(), amove.GetDescription()))
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return data, true
}

// CastingCompleter represents ...
type CastingCompleter struct {
	*syntax.CompleterArgument
	bt *battle.Battle
}

// Query returns the query for any node completer.
func (cc *CastingCompleter) Query(ctx *syntax.Context, content syntax.IContent, line interface{}, index int) (interface{}, bool) {
	var data []*syntax.CompleteHelp
	for _, actor := range casting.GetActors() {
		data = append(data, syntax.NewCompleteHelp(actor.GetName(), actor.GetDescription()))
	}
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
		amove = actor.GetAmoveByName("UpperCut")
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
		"actor [actor-name]?",
		"Display all available actors",
		[]*syntax.Argument{
			syntax.NewArgument(
				"actor-name",
				"Actor name",
				&ActorCompleter{syntax.NewCompleterArgument("actor-name"), bt},
				"string",
				"",
				nil),
		},
		nil)
	displayActorCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		//params := arguments.(map[string]interface{})
		//actorName := params["actor-name"].(string)
		actorName := tools.GetStringFromArgs(arguments, "actor-name")
		if actorName == "" {
			tools.ToDisplay("Display all actors...\n")
			for _, actor := range bt.Actors {
				tools.ToDisplay("-> %s\n", actor.GetName())
				for _, tech := range actor.GetTechniques() {
					tools.ToDisplay("--> [TECH] %s\n", tech.GetDN())
				}
				for _, x := range actor.GetStyles() {
					tools.ToDisplay("--> [STYL] %s\n", x.GetDN())
				}
				for _, x := range actor.GetStances() {
					tools.ToDisplay("--> [STAN] %s\n", x.GetDN())
				}
				for _, x := range actor.GetAmoves() {
					tools.ToDisplay("--> [MOVE] %s\n", x.GetDN())
				}
			}
		} else {
			if actor := bt.GetActorByName(actorName); actor != nil {
				tools.ToDisplay("Display actor: %s\n", actor.GetName())
				tools.ToDisplay("-->%s\n", actor.GetDescription())
				tools.ToDisplay("-->%#v\n", actor.GetStats())
				if actor.GetTechnique() != nil {
					tools.ToDisplay("-->%#s\n", actor.GetTechnique().GetDN())
				}
				if actor.GetStyle() != nil {
					tools.ToDisplay("-->%#s\n", actor.GetStyle().GetDN())
				}
				if actor.GetStance() != nil {
					tools.ToDisplay("-->%#s\n", actor.GetStance().GetDN())
				}
				if actor.GetAmove() != nil {
					tools.ToDisplay("-->%#s\n", actor.GetAmove().GetDN())
				}
			}
		}
		return nil
	}

	displayCastingCommand := syntax.NewCommand(
		displayCommand,
		"casting",
		"Display actors in casting",
		nil,
		nil)
	displayCastingCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("Actors Casting:\n")
		for _, actor := range casting.GetActors() {
			tools.ToDisplay("-->%s\n", actor.GetName())
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
		bt.AddActor(battle.NewActor(name, desc, battle.NewStats()))
		tools.ToDisplay("Create actor: %#v\n", name)
		return nil
	}

	createCastingActorCommand := syntax.NewCommand(
		createCommand,
		"casting-actor casting-name name desc",
		"Create actor from the casting pool",
		[]*syntax.Argument{
			syntax.NewArgument(
				"casting-name",
				"Actor casting name",
				&CastingCompleter{syntax.NewCompleterArgument("casting-name"), bt},
				"string",
				"none",
				nil),
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
	createCastingActorCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		castingName := tools.GetStringFromArgs(arguments, "casting-name")
		actorName := tools.GetStringFromArgs(arguments, "name")
		actorDesc := tools.GetStringFromArgs(arguments, "desc")
		actor := casting.CreateCastingActor(castingName, actorName, actorDesc)
		bt.AddActor(actor)
		tools.ToDisplay("Create actor: %#v\n", actorName)
		return nil

	}

	updateCommand := syntax.NewCommand(
		nil,
		"update",
		"Update ...",
		nil,
		nil)
	updateCommand.JustPrefix = true

	updateActorCommand := syntax.NewCommand(
		updateCommand,
		"actor actor-name [str | agi | sta | pre | foc]+",
		"Update actor stat",
		[]*syntax.Argument{
			syntax.NewArgument(
				"actor-name",
				"Actor name",
				&ActorCompleter{syntax.NewCompleterArgument("actor-name"), bt},
				"string",
				"none",
				nil),
			syntax.NewArgument("str", "Actor strength", nil, "int", 0, nil),
			syntax.NewArgument("agi", "Actor agility", nil, "int", 0, nil),
			syntax.NewArgument("sta", "Actor stamina", nil, "int", 0, nil),
			syntax.NewArgument("pre", "Actor precision", nil, "int", 0, nil),
			syntax.NewArgument("foc", "Actor focus", nil, "int", 0, nil),
		},
		nil)
	updateActorCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		actorName := params["actor-name"].(string)
		if actor := bt.GetActorByName(actorName); actor != nil {
			if params["-str"] != nil {
				tools.ToDisplay("update actor:%s strength to %d\n", actorName, params["str"].(int))
				actor.GetStats().Str = battle.InterfaceToTStat(params["str"])
			}
			if params["-agi"] != nil {
				tools.ToDisplay("update actor:%s agility to %d\n", actorName, params["agi"].(int))
				actor.GetStats().Agi = battle.InterfaceToTStat(params["agi"])
			}
			if params["-sta"] != nil {
				tools.ToDisplay("update actor:%s stamina to %d\n", actorName, params["sta"].(int))
				actor.GetStats().Sta = battle.InterfaceToTStat(params["sta"])
			}
			if params["-pre"] != nil {
				tools.ToDisplay("update actor:%s precission to %d\n", actorName, params["pre"].(int))
				actor.GetStats().Pre = battle.InterfaceToTStat(params["pre"])
			}
			if params["-foc"] != nil {
				tools.ToDisplay("update actor:%s focus to %d\n", actorName, params["foc"].(int))
				actor.GetStats().Foc = battle.InterfaceToTStat(params["foc"])
			}
			tools.ToDisplay("actor stats: %#v\n", actor.GetStats())
		}
		return nil
	}

	selectCommand := syntax.NewCommand(
		nil,
		"select",
		"Select ...",
		nil,
		nil)
	selectCommand.JustPrefix = true

	selectAmoveCommand := syntax.NewCommand(
		selectCommand,
		"amove actor-name tech-name style-name stance-name amove-name",
		"Select amove from an actor",
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
				&TechniqueCompleter{syntax.NewCompleterArgument("tech-name"), bt, "actor-name"},
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"style-name",
				"Style name",
				&StyleCompleter{syntax.NewCompleterArgument("tech-name"), bt, "actor-name", "tech-name"},
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"stance-name",
				"stance name",
				&StanceCompleter{syntax.NewCompleterArgument("tech-name"), bt, "actor-name", "tech-name", "style-name"},
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"amove-name",
				"amove name",
				&AmoveCompleter{syntax.NewCompleterArgument("tech-name"), bt, "actor-name", "tech-name", "style-name", "stance-name"},
				"string",
				"none",
				nil),
		},
		nil)
	selectAmoveCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		actorName := tools.GetStringFromArgs(arguments, "actor-name")
		techName := tools.GetStringFromArgs(arguments, "tech-name")
		styleName := tools.GetStringFromArgs(arguments, "style-name")
		stanceName := tools.GetStringFromArgs(arguments, "stance-name")
		amoveName := tools.GetStringFromArgs(arguments, "amove-name")
		tools.ToDisplay("actor:%s select %s:%s:%s:%s\n", actorName, techName, styleName, stanceName, amoveName)
		if actor := bt.GetActorByName(actorName); actor != nil {
			if ok := actor.SetTechniqueByName(techName); !ok {
				return nil
			}
			if ok := actor.SetStyleByName(styleName); !ok {
				return nil
			}
			if ok := actor.SetStanceByName(stanceName); !ok {
				return nil
			}
			if ok := actor.SetAmoveByName(amoveName); !ok {
				return nil
			}
		}
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
				&TechniqueCompleter{syntax.NewCompleterArgument("tech-name"), bt, ""},
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
				&ActorCompleter{syntax.NewCompleterArgument("orig-name"), bt},
				"string",
				"none",
				nil),
			syntax.NewArgument(
				"target-name",
				"Actor target name",
				&ActorCompleter{syntax.NewCompleterArgument("target-name"), bt},
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
		displayCastingCommand,
		createCommand,
		createActorCommand,
		createCastingActorCommand,
		selectCommand,
		selectAmoveCommand,
		updateCommand,
		updateActorCommand,
		addTechToActorCommand,
		engageCommand,
	}
	return cmds
}
