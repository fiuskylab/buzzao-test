package workerpool

import (
	"fmt"
	"sync"
)

// Pool manages Jobs and Workers
type Pool struct {
	Jobs        []*Job
	Workers     []*Worker
	Concurrency int

	collector chan *Job
	runAsync  chan bool
	wg        sync.WaitGroup
}

// NewPool return a Pool
// with given params
func NewPool(jobs []*Job, c int) *Pool {
	return &Pool{
		Jobs:        jobs,
		Concurrency: c,
		collector:   make(chan *Job, 1000),
	}
}

func (p *Pool) SetConcurrency(i int) error {
	if i <= 0 {
		return fmt.Errorf("Number of threads must be higher than 0, given: %d", i)
	}
	p.Concurrency = i
	return nil
}

func (p *Pool) Run() {
	for i := 0; i < p.Concurrency; i++ {
		w := NewWorker(p.collector)
		w.Run(&p.wg)
	}

	for _, j := range p.Jobs {
		p.collector <- j
	}

	close(p.collector)

	p.wg.Wait()
}

func (p *Pool) AddJob(j *Job) {
	p.collector <- j
}

func (p *Pool) RunAsync() {
	for i := 0; i < p.Concurrency; i++ {
		w := NewWorker(p.collector)
		p.Workers = append(p.Workers, w)

		go w.StartAsync()
	}

	for _, j := range p.Jobs {
		p.collector <- j
	}

	p.runAsync = make(chan bool)
	<-p.runAsync
}

func (p *Pool) Stop() {
	for _, w := range p.Workers {
		w.Stop()
	}

	p.runAsync <- true
}
