package queue

// Queue is a First-In-First-Out (FIFO) structure that allows enqueuing of Jobs to be processed asyncronusly
type Queue interface {

	// Enqueue attempts to add a Job to the queue and returns true if successful otherwise false
	Enqueue(job Job) bool

	// EnqueueBlock adds a Job to the queue blocking the thread until their is capacity
	EnqueueBlock(job Job)

	// Close disallows any futher Jobs from being added to the queue
	Close()

	// Wait blocks until all queued Jobs to finish being processed
	Wait()

	// CloseAndWait disallows any futher Jobs from being added to the queue and blocks until all queued Jobs to finish being processed
	CloseAndWait()
}
