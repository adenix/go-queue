package queue

import "sync"

// workerQueue is a First-In-First-Out (FIFO) structure that allows enqueuing of Jobs to be processed asyncronusly
type workerQueue struct {
	ch chan Job
	wg *sync.WaitGroup
}

// Ensures workerQueue implements the Queue interface
var _ Queue = (*workerQueue)(nil)

// New constructs a queue with a specified capacity and concurrency to process Jobs
func NewWorkerQueue(capacity uint32, concurrency uint32) *workerQueue {
	q := &workerQueue{
		ch: make(chan Job, capacity),
		wg: &sync.WaitGroup{},
	}

	for i := uint32(0); i < concurrency; i++ {
		q.wg.Add(1)
		go q.worker()
	}

	return q
}

// worker pops items off the queue and executes the specified Action against them
func (q *workerQueue) worker() {
	defer q.wg.Done()

	for job := range q.ch {
		job()
	}
}

// Enqueue attempts to add a Job to the queue and returns true if successful otherwise false
func (q *workerQueue) Enqueue(job Job) bool {
	select {
	case q.ch <- job:
		return true
	default:
		return false
	}
}

// EnqueueBlock adds a Job to the queue blocking the thread until their is capacity
func (q *workerQueue) EnqueueBlock(job Job) {
	q.ch <- job
}

// Close disallows any futher Jobs from being added to the queue
func (q *workerQueue) Close() {
	close(q.ch)
}

// Wait blocks until all queued Jobs to finish being processed
func (q *workerQueue) Wait() {
	q.wg.Wait()
}

func (q *workerQueue) CloseAndWait() {
	q.Close()
	q.Wait()
}
