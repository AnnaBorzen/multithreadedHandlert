package workerPool

import (
	"fmt"
	"time"
)

type Result struct {
	JobID    int
	Data     interface{}
	Status   string
	Duration time.Duration
}

func (result *Result) Print() {
	fmt.Printf("Job ID: %d; Data: %v; Status: %s; Duration: %s\n", result.JobID, result.Data, result.Status, result.Duration)
}
