# Go Queue

A simple importable queue library for Golang.

## Example Usage

```go
package main

import (
    "fmt"

    "github.com/adenix/go-queue"
)

func main() {
    q := queue.NewWorkerQueue(10, 10)

    for i := 0; i < 100; i++ {
        q.EnqueueBlock(newJob(fmt.Sprintf("Austin %d", i)))
    }
}

func newJob(name string) queue.Job {
    return func() {
        fmt.Printf("Hello, %s!\n", name)
    }
}
```
