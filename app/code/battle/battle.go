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
	TechBuilderCb TechniqueBuilderCb
}

// NewTechniqueBuilder is ...
func NewTechniqueBuilder(name string, cb TechniqueBuilderCb) *TechniqueBuilder {
	return &TechniqueBuilder{
		Name:          name,
		TechBuilderCb: cb,
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
	origAmove := b.SelectAmove(orig, AmodeAttack)
	targetAmove := b.SelectAmove(target, AmodeDefence)
	b.ExecuteEngage(orig, origAmove, target, targetAmove)
}

// SelectAmove is ...
func (b *Battle) SelectAmove(orig IActor, mode Amode) IAmove {
	if b.Selector != nil {
		return b.Selector.SelectAmove(orig, mode)
	}
	return nil
}

// ExecuteEngage is ...
func (b *Battle) ExecuteEngage(orig IActor, origAmove IAmove, target IActor, targetAmove IAmove) {
	tools.ToDisplay("Engage %s:%s vs %s:%s\n", orig.GetName(),
		origAmove.GetName(),
		target.GetName(),
		targetAmove.GetName())
}

// NewBattle is ...
func NewBattle() *Battle {
	return &Battle{}
}