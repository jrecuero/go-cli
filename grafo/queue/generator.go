package queue

import "github.com/jrecuero/go-cli/grafo"

// Generator represents ...
type Generator struct {
	*grafo.Vertex
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
