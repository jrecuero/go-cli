package grafo

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
	*Branch
	Jobs  []*Job
	limit int
}

// Check is ...
func (queue *Queue) Check(params ...interface{}) (interface{}, bool) {
	var topass = []interface{}{queue}
	topass = append(topass, params...)
	return queue.clearance(queue.GetParent(), queue.GetChild(), topass...)
}

// NewQueue is ...
func NewQueue(parent *Leaf, child *Leaf, limit int) *Queue {
	return &Queue{
		Branch: NewBranch(parent,
			child,
			func(parent *Leaf, child *Leaf, params ...interface{}) (interface{}, bool) {
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
	label string
}

// GetLabel is ...
func (sc *ServerContent) GetLabel() string {
	return sc.label
}

// Serve is ...
func (sc *ServerContent) Serve(jobs []*Job) (interface{}, bool) {
	return nil, true
}

// Server represents ...
type Server struct {
	*Leaf
}

// NewServer is ...
func NewServer(label string, sc *ServerContent) *Server {
	server := &Server{
		NewLeaf(label),
	}
	server.Content = sc
	return server
}

// System represents ..
type System struct {
	*Tree
}

// AddQueue is ...
func (system *System) AddQueue(parent *Leaf, child *Leaf, limit int) error {
	if parent == nil {
		parent = system.GetRoot()
	}
	var branch IBranch = NewQueue(parent, child, limit)
	return system.AddBranch(parent, branch)
}

// NewSystem is ...
func NewSystem(label string) *System {
	return &System{
		NewTree(label),
	}
}

// ServerToLeaf is ...
func ServerToLeaf(server *Server) *Leaf {
	return server.Leaf
}

// ToServer is ...
func ToServer(leaf *Leaf) *Server {
	return &Server{
		leaf,
	}
}
