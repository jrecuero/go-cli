package queue

import (
	"github.com/jrecuero/go-cli/grafo"
)

// Queue represents ...
type Queue struct {
	*grafo.Edge
	Jobs  []*Job
	limit int
}

// Check is ...
func (queue *Queue) Check(params ...interface{}) (interface{}, bool) {
	var topass = []interface{}{queue}
	topass = append(topass, params...)
	return queue.Clearance(queue.GetParent(), queue.GetChild(), topass...)
}

// NewQueue is ...
func NewQueue(parent *grafo.Vertex, child *grafo.Vertex, limit int) *Queue {
	return &Queue{
		Edge: grafo.NewEdge(parent,
			child,
			func(parent *grafo.Vertex, child *grafo.Vertex, params ...interface{}) (interface{}, bool) {
				queue := params[0].(*Queue)
				job := params[1].(*Job)
				if len(queue.Jobs) > limit {
					return nil, false
				}
				queue.Jobs = append(queue.Jobs, job)
				return queue.Jobs, true
			}),
		limit: limit,
	}
}
