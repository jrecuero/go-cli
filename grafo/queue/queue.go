package queue

import "github.com/jrecuero/go-cli/grafo"

// Job represents ...
type Job struct {
	Label   string
	Workout int
}

// NewJob is ...
func NewJob(label string, workout int) *Job {
	return &Job{
		Label:   label,
		Workout: workout,
	}
}

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

// ServerContent represents ...
type ServerContent struct {
	label   string
	job     *Job
	worker  int
	workout int
}

// GetLabel is ...
func (sc *ServerContent) GetLabel() string {
	return sc.label
}

// Serve is ...
func (sc *ServerContent) Serve(jobs *[]*Job) (interface{}, bool) {
	if sc.job == nil {
		if len(*jobs) != 0 {
			sc.job = (*jobs)[0]
			sc.workout = sc.job.Workout
			*jobs = (*jobs)[1:len(*jobs)]
		}
	}
	if sc.job != nil {
		sc.workout -= sc.worker
		if sc.workout <= 0 {
			sc.workout = 0
			job := sc.job
			sc.job = nil
			return job, true
		}
	}
	return nil, false
}

// NewServerContent is ...
func NewServerContent(label string, worker int) *ServerContent {
	return &ServerContent{
		label:  label,
		worker: worker,
	}
}

// Server represents ...
type Server struct {
	*grafo.Vertex
}

// NewServer is ...
func NewServer(label string, sc *ServerContent) *Server {
	server := &Server{
		grafo.NewVertex(label),
	}
	server.Content = sc
	return server
}

// System represents ..
type System struct {
	*grafo.Grafo
}

// AddQueue is ...
func (system *System) AddQueue(parent *grafo.Vertex, child *grafo.Vertex, limit int) error {
	if parent == nil {
		parent = system.GetRoot()
	}
	var edge grafo.IEdge = NewQueue(parent, child, limit)
	return system.AddEdge(parent, edge)
}

// NewSystem is ...
func NewSystem(label string) *System {
	return &System{
		grafo.NewGrafo(label),
	}
}

// ServerToVertex is ...
func ServerToVertex(server *Server) *grafo.Vertex {
	return server.Vertex
}

// ToServer is ...
func ToServer(vertex *grafo.Vertex) *Server {
	return &Server{
		vertex,
	}
}
