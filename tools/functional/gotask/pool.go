package gotask

import (
	"context"
	"sync"
	"time"
)

type taskFunc func(ctx context.Context)

type Pool struct {
	stopCh chan struct{}
	wg     sync.WaitGroup
	lock   sync.Mutex
	tasks  []*task
}

func NewPool() *Pool {
	return &Pool{
		stopCh: make(chan struct{}),
	}
}

func (p *Pool) AddTask(fn taskFunc, interval time.Duration) {
	t := &task{
		fn:       fn,
		interval: interval,
	}

	p.lock.Lock()
	defer p.lock.Unlock()
	p.tasks = append(p.tasks, t)
}

func (p *Pool) Start(ctx context.Context) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, task := range p.tasks {
		p.wg.Add(1)
		go p.doTask(ctx, task)
	}
}

func (p *Pool) doTask(ctx context.Context, task *task) {
	defer p.wg.Done()
	for {
		select {
		// wait for stop signal
		case <-ctx.Done():
			return
		case <-p.stopCh:
			return
		default:
			task.fn(ctx)
			// interval is 0 means it's a one time job, return directly
			if task.interval == 0 {
				return
			}
			time.Sleep(task.interval)
		}
	}
}

func (p *Pool) Stop() {
	close(p.stopCh)
	p.wg.Wait()
}

type task struct {
	fn       taskFunc
	interval time.Duration
}
