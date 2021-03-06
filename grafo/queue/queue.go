package queue

import (
	"github.com/jrecuero/go-cli/engine"
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

// QueueEvent is ...
func (queue *Queue) QueueEvent(job *Job, followUp func(*[]*Job)) *engine.Event {
	ev := engine.NewEvent("event/gen-to-que/0", 2)
	ev.SetCallback(func(params ...interface{}) error {
		//tools.ToDisplay("queue event callback\n")
		queue.Check(job)
		followUp(&queue.Jobs)
		return nil
	}, nil)
	return ev
}

// NewQueue is ...
func NewQueue(parent grafo.IVertex, child grafo.IVertex, limit int) *Queue {
	return &Queue{
		Edge: grafo.NewEdge(parent,
			child,
			func(parent grafo.IVertex, child grafo.IVertex, params ...interface{}) (interface{}, bool) {
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
