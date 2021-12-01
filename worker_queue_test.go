package queue

import (
	"fmt"
	"sync"
	"testing"
)

type spy struct {
	sync.Mutex
	called uint16
}

func newMockJob() (Job, *spy) {
	s := &spy{}
	return func() JobResult {
		s.Lock()
		s.called++
		s.Unlock()
		return Success
	}, s
}

func TestQueue(t *testing.T) {
	cases := []struct {
		capacity       uint32
		concurrency    uint32
		enqueue        int
		calledExpected uint16
	}{
		{
			capacity:       2,
			concurrency:    0,
			enqueue:        2,
			calledExpected: 0,
		},
		{
			capacity:       2,
			concurrency:    1,
			enqueue:        2,
			calledExpected: 2,
		},
		{
			capacity:       2,
			concurrency:    2,
			enqueue:        500,
			calledExpected: 500,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("EnqueueBlock-%d", i), func(t *testing.T) {
			j, s := newMockJob()
			q := NewWorkerQueue(c.capacity, c.concurrency)

			wg := new(sync.WaitGroup)
			for i := 0; i < c.enqueue; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					q.EnqueueBlock(j)
				}()
			}

			wg.Wait()
			q.CloseAndWait()

			s.Lock()
			if s.called != c.calledExpected {
				t.Errorf("called should be %d but is %d", c.calledExpected, s.called)
			}
			s.Unlock()
		})
	}
}

func TestEnqueue(t *testing.T) {
	j, s := newMockJob()
	q := NewWorkerQueue(2, 0)

	failed := 0
	for i := 0; i < 4; i++ {
		if ok := q.Enqueue(j); !ok {
			failed++
		}
	}

	if failed != 2 {
		t.Errorf("expected 2 queue failures, got %d", failed)
	}
	if s.called != 0 {
		t.Errorf("called should be 0 but is %d", s.called)
	}

	<-q.(*workerQueue).ch
	if ok := q.Enqueue(j); !ok {
		failed++
	}

	if failed != 2 {
		t.Errorf("expected 2 queue failures, got %d", failed)
	}
}

func TestEnqueueBlock(t *testing.T) {
	j, s := newMockJob()
	q := NewWorkerQueue(2, 0)

	outterWg := new(sync.WaitGroup)
	innerWg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	innerComplete := 0
	for i := 0; i < 4; i++ {
		outterWg.Add(1)
		go func() {
			defer outterWg.Done()
			innerWg.Add(1)
			go func() {
				defer innerWg.Done()
				q.EnqueueBlock(j)
				mu.Lock()
				innerComplete++
				mu.Unlock()
			}()
		}()
	}

	outterWg.Wait()
	innerWg.Done()
	innerWg.Done()
	innerWg.Wait()

	s.Lock()
	if s.called != 0 {
		t.Errorf("called should be 0 but is %d", s.called)
	}
	s.Unlock()

	mu.Lock()
	if innerComplete != 2 {
		t.Errorf("inner complete should be 2 but is %d", innerComplete)
	}
	mu.Unlock()
}
