package novel

import "github.com/jrecuero/go-cli/tools"

// ActionNames represents any action in the app.
type ActionNames struct {
	Origins []string
	Actions []string
	Targets []string
	Flags   []string
}

// AddOrigin adds a new origin
func (an *ActionNames) AddOrigin(in string) *ActionNames {
	an.Origins = append(an.Origins, in)
	return an
}

// AddAction adds a new action.
func (an *ActionNames) AddAction(in string) *ActionNames {
	an.Actions = append(an.Actions, in)
	return an
}

// AddTarget adds a new target.
func (an *ActionNames) AddTarget(in string) *ActionNames {
	an.Targets = append(an.Targets, in)
	return an
}

// AddFlags adds a new flag.
func (an *ActionNames) AddFlags(in string) *ActionNames {
	an.Flags = append(an.Flags, in)
	return an
}

// ActionCallback is ...
type ActionCallback func(origins []*Actor, target []*Actor, flags []string) error

// ActionSequence is ...
type ActionSequence struct {
	Origins []*Actor
	Actions []ActionCallback
	Targets []*Actor
	Flags   []string
}

// RunAction is ...
func (as *ActionSequence) RunAction() error {
	//tools.ToDisplay("RunAction: %#v\n", as)
	for _, action := range as.Actions {
		if err := action(as.Origins, as.Targets, as.Flags); err != nil {
			return err
		}
	}
	return nil
}

// ActionHit is ...
func ActionHit(novel *Novel) ActionCallback {
	return func(origins []*Actor, targets []*Actor, flags []string) error {
		//tools.ToDisplay("Hit: origins: %#v targets: %#v\n", origins, targets)
		if len(origins) != 1 {
			return tools.ERROR(nil, false, "Too many origins (%d) for action 'Hit'\n", len(origins))
		}
		for _, target := range targets {
			origin := origins[0]
			damage := origin.Strength
			target.Life -= damage
			tools.ToDisplay("%#v hits with %d damage to %#v: %d life points\n", origin.Name, origin.Strength, target.Name, target.Life)
		}
		return nil
	}
}
