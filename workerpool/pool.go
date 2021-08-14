package workerpool

import "sync"

// Pool manages Jobs and Workers
type Pool struct {
	Jobs    []*Job
	Workers []*Worker

	concurrency int
	collector   chan *Job
	runAsync    chan bool
	wg          sync.WaitGroup
}

// NewPool return a Pool
// with given params
func NewPool(jobs []*Job, c int) *Pool {
	return &Pool{
		Jobs:        jobs,
		concurrency: c,
		collector:   make(chan *Job, 1000),
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
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
	for i := 0; i < p.concurrency; i++ {
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
