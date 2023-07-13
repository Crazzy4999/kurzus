package workers

import "sync"

type Worker interface {
	allocate(jobsNum int)
	worker(wg *sync.WaitGroup)
	createWorkerPool(workersNum int)
	Start(min int)
}
