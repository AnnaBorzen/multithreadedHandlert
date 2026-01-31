package workerPool

import (
	"sync"
	"time"
)

type WorkerPool struct {
	jobChan    chan *Job
	resultChan chan *Result

	wg    *sync.WaitGroup
	count int
}

func NewWorkerPool(channel chan *Job, resultChan chan *Result, wg *sync.WaitGroup, count int) *WorkerPool {
	return &WorkerPool{
		jobChan:    channel,
		resultChan: resultChan,
		wg:         wg,
		count:      count,
	}
}

func worker(jobChan chan *Job, results chan *Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobChan {
		start := time.Now()
		status := "success"

		errFunc := job.Func(job.Data)
		if errFunc != nil {
			status = "failed"
		}

		duration := time.Since(start)

		result := &Result{
			JobID:    job.ID,
			Data:     job.Data,
			Status:   status,
			Duration: duration,
		}

		results <- result
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.count; i++ {
		wp.wg.Add(1)
		go worker(wp.jobChan, wp.resultChan, wp.wg)
	}
}

func (wp *WorkerPool) Stop() {
	wp.wg.Wait()
	close(wp.resultChan)
}
