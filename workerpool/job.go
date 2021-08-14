package workerpool

import "github.com/fiuskylab/buzzao-test/data"

// Job is the process to be dispatched
type Job struct {
	Err  error
	Data *data.Data
	f    func(d *data.Data)
}

// NewJob return a Job pointer
// with given params
func NewJob(f func(d *data.Data), d *data.Data) *Job {
	return &Job{
		f:    f,
		Data: d,
	}
}

func exec(j *Job) {
	j.f(j.Data)
}
