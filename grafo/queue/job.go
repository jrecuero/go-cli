package queue

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
