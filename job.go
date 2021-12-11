package queue

// JobResult represents the result of a individual Job processed.
// Useful for testing a job as the queue processes asynchronously
// and ignores the return value.
type JobResult uint16

//go:generate go run golang.org/x/tools/cmd/stringer -type=JobResult
const (
	Fail JobResult = iota
	Success
	Skip
)

// Job is a single unit of work that can be enqueued for processing
type Job func() JobResult
