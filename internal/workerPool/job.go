package workerPool

type Job struct {
	ID   int
	Data interface{}
	Func func(interface{}) error
}
