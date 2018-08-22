package queue

import (
	"github.com/jrecuero/go-cli/engine"
	"github.com/jrecuero/go-cli/grafo"
	"github.com/jrecuero/go-cli/tools"
)

// GenFollowUp represents ...
type GenFollowUp func(job *Job)

// Generator represents ...
type Generator struct {
	*grafo.Vertex
}

//GenEvent is ..
func (gen *Generator) GenEvent(attime engine.ETime, followUp GenFollowUp) *engine.Event {
	ev := engine.NewEvent("event/gen/0", attime)
	ev.SetCallback(func(params ...interface{}) error {
		//tools.ToDisplay("generator event callback\n")
		if job, ok := GetGeneratorContent(gen).Generate(); ok {
			followUp(job.(*Job))
			return nil
		}
		return tools.ERROR(nil, false, "Generator error")
	}, nil)
	return ev
}

// NewGenerator is ...
func NewGenerator(label string, gc *GeneratorContent) *Generator {
	gen := &Generator{
		grafo.NewVertex(label),
	}
	gen.Content = gc
	return gen
}

// GeneratorToVertex is ...
func GeneratorToVertex(gen *Generator) *grafo.Vertex {
	return gen.Vertex
}

// ToGenerator is ...
func ToGenerator(vertex *grafo.Vertex) *Generator {
	return &Generator{
		vertex,
	}
}
