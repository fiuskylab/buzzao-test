package workerpool

import "github.com/fiuskylab/buzzao-test/data"

// Job is the process to be dispatched
type Job struct {
	Err  error
	Data *data.Data
	f    func(d *data.Data) error
}

// NewJob return a Job pointer
// with given params
func NewJob(f func(d *data.Data) error, d *data.Data) *Job {
	return &Job{
		f:    f,
		Data: d,
	}
}

func exec(j *Job) error {
	return j.f(j.Data)
}
