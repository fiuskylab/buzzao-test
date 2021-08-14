package workerpool

import "sync"

// Worker is the responsible
// for processing the Jobs
type Worker struct {
	jobChan chan *Job
	stop    chan bool
}

// NewWorker return a Worker
// with given params
func NewWorker(ch chan *Job) *Worker {
	return &Worker{
		jobChan: ch,
	}
}

// Run exec each Job from jobChan
func (w *Worker) Run(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for j := range w.jobChan {
			exec(j)
		}
	}()
}

// StartAsync dispatch a worker into
// the background to exec jobs
// asynchronously
func (w *Worker) StartAsync() {
	for {
		select {
		case j := <-w.jobChan:
			exec(j)
		case <-w.stop:
			return
		}
	}
}

// Stop the Worker
func (w *Worker) Stop() {
	go func() {
		w.stop <- true
	}()
}
