package battle

import "github.com/jrecuero/go-cli/tools"

// ISelector represents ...
type ISelector interface {
	SelectOrig(...interface{}) IActor
	SelectTarget(...interface{}) IActor
	SelectAmove(...interface{}) IAmove
}

// Stage represents ...
type Stage struct {
	Name string
}

// NewStage is ...
func NewStage(name string) *Stage {
	return &Stage{
		Name: name,
	}
}

// TechniqueBuilderCb represents ...
type TechniqueBuilderCb func(...interface{}) ITechnique

// TechniqueBuilder represents ...
type TechniqueBuilder struct {
	Name          string
	Desc          string
	TechBuilderCb TechniqueBuilderCb
}

// NewTechniqueBuilder is ...
func NewTechniqueBuilder(name string, desc string, cb TechniqueBuilderCb) *TechniqueBuilder {
	return &TechniqueBuilder{
		Name:          name,
		Desc:          desc,
		TechBuilderCb: cb,
	}
}

// Engagement represents ...
type Engagement struct {
	Origin      IActor
	Target      IActor
	OriginAmove IAmove
	TargetAmove IAmove
	Active      bool
}

// NewEngagement is ...
func NewEngagement() *Engagement {
	return &Engagement{
		Active: true,
	}
}

// Battle represents ...
type Battle struct {
	TechBuilders []*TechniqueBuilder
	Actors       []IActor
	stage        Stage
	Selector     ISelector
}

// GetTechBuilderByName is ...
func (b *Battle) GetTechBuilderByName(name string) *TechniqueBuilder {
	for _, tb := range b.TechBuilders {
		if tb.Name == name {
			return tb
		}
	}
	return nil
}

// AddTechBuilder is ...
func (b *Battle) AddTechBuilder(tb *TechniqueBuilder) *Battle {
	if b.GetTechBuilderByName(tb.Name) == nil {
		b.TechBuilders = append(b.TechBuilders, tb)
		return b
	}
	return nil
}

// CreateTechniqueByName is ...
func (b *Battle) CreateTechniqueByName(name string, args ...interface{}) ITechnique {
	if tb := b.GetTechBuilderByName(name); tb != nil {
		return tb.TechBuilderCb(args...)
	}
	return nil
}

// GetActorByName is ...
func (b *Battle) GetActorByName(name string) IActor {
	for _, actor := range b.Actors {
		if actor.GetName() == name {
			return actor
		}
	}
	return nil
}

// AddActor is ...
func (b *Battle) AddActor(actor IActor) *Battle {
	if b.GetActorByName(actor.GetName()) == nil {
		b.Actors = append(b.Actors, actor)
		return b
	}
	return nil
}

// Engage is ...
func (b *Battle) Engage(orig IActor, target IActor) {
	origAmove := orig.GetAmove()
	if origAmove == nil {
		origAmove = b.SelectAmove(orig, AmodeAttack)
	}
	targetAmove := target.GetAmove()
	if targetAmove == nil {
		targetAmove = b.SelectAmove(target, AmodeDefence)
	}
	b.ExecuteEngage(orig, origAmove, target, targetAmove)
}

// SelectAmove is ...
func (b *Battle) SelectAmove(orig IActor, mode Amode) IAmove {
	if b.Selector != nil {
		return b.Selector.SelectAmove(orig, mode)
	}
	return nil
}

// getEngageActorStr is ...
func (b *Battle) getEngageActorStr(actor IActor, amove IAmove) int {
	str := actor.GetStats().Str
	str = amove.GetTechnique().GetUpdateStats().UStr(str, actor)
	str = amove.GetStyle().GetUpdateStats().UStr(str, actor)
	str = amove.GetStance().GetUpdateStats().UStr(str, actor)
	str = amove.GetUpdateStats().UStr(str, actor)
	return int(str)
}

// getEngageActorSta is ...
func (b *Battle) getEngageActorSta(actor IActor, amove IAmove) int {
	sta := actor.GetStats().Sta
	sta = amove.GetTechnique().GetUpdateStats().USta(sta, actor)
	sta = amove.GetStyle().GetUpdateStats().USta(sta, actor)
	sta = amove.GetStance().GetUpdateStats().USta(sta, actor)
	sta = amove.GetUpdateStats().USta(sta, actor)
	return int(sta)
}

// ExecuteEngage is ...
func (b *Battle) ExecuteEngage(orig IActor, origAmove IAmove, target IActor, targetAmove IAmove) {
	str := b.getEngageActorStr(orig, origAmove)
	sta := b.getEngageActorSta(target, targetAmove)
	tools.ToDisplay("Engage ATK: %s:%s:%d\n", orig.GetName(), origAmove.GetName(), str)
	tools.ToDisplay("Engage DEF: %s:%s:%d\n", target.GetName(), targetAmove.GetName(), sta)
}

// NewBattle is ...
func NewBattle() *Battle {
	return &Battle{}
}
