package api

type WorkerPool interface {
	// add a task to pool to excute
	Excute()

	// wait for all tasks to end
	Wait()

	// stop all running tasks, this cause Wait() to return immediately
	Cancel()
}
