package queue

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
