package queue

import "fmt"

// JobResult represents the result of a individual Job processed.
// Useful for testing a job as the queue processes asynchronously
// and ignores the return value.
type JobResult uint16

// Ensures JobResult implements the fmt.Stringer interface
var _ fmt.Stringer = (*JobResult)(nil)

const (
	Fail JobResult = iota
	Success
	Skip
)

// String converts a JobResult into a human readable string
func (r JobResult) String() string {
	switch r {
	case Fail:
		return "fail"
	case Success:
		return "success"
	case Skip:
		return "skip"
	}
	return "unknown"
}

// Job is a single unit of work that can be enqueued for processing
type Job func() JobResult
