package gotask

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	pool := NewPool()

	taskNum := 3
	taskInterval := time.Duration(0)
	for i := range taskNum {
		fn := func(ctx context.Context) {
			t.Logf("Task %d is running...", i)
		}

		pool.AddTask(fn, taskInterval)
	}

	if len(pool.tasks) != taskNum {
		t.Errorf("Expected %d tasks, got %d.", taskNum, len(pool.tasks))
	}
}

func TestStartAndStop(t *testing.T) {
	// test normal case
	{
		pool := NewPool()
		// create channel with buffer
		ch1 := make(chan struct{}, 5)
		ch2 := make(chan struct{}, 5)
		// test one-time job
		t1 := &task{
			interval: 0,
			fn: func(ctx context.Context) {
				ch1 <- struct{}{}
			},
		}
		// test interval job
		t2 := &task{
			interval: 100 * time.Millisecond,
			fn: func(ctx context.Context) {
				ch2 <- struct{}{}
			},
		}

		pool.tasks = []*task{t1, t2}

		ctx1 := t.Context()
		pool.Start(ctx1)

		// Let it run for a bit
		time.Sleep(300 * time.Millisecond)
		// ch1 should only have one element as it's a one time job
		assert.Equal(t, 1, len(ch1))
		// ch2 should have elements over 2 as sleep 300ms and interval is 100ms
		assert.Greater(t, len(ch2), 2)
		pool.Stop()
		close(ch1)
		close(ch2)
	}

	// test context timeout case
	{
		pool := NewPool()
		ch1 := make(chan struct{}, 2)
		t1 := &task{
			interval: 100 * time.Millisecond,
			fn: func(ctx context.Context) {
				ch1 <- struct{}{}
			},
		}

		pool.tasks = []*task{t1}
		ctx1, cancel1 := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel1()
		pool.Start(ctx1)
		// Let it run for a bit
		time.Sleep(200 * time.Millisecond)
		assert.Equal(t, 1, len(ch1))
		pool.Stop()
		close(ch1)
	}
}
