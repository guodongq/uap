package semaphore

type Semaphore struct {
	queue chan struct{}
}

func NewSemaphore(size int32) *Semaphore {
	return &Semaphore{
		queue: make(chan struct{}, size),
	}
}

func (w *Semaphore) Acquire() {
	w.queue <- struct{}{}
}

func (w *Semaphore) Release() {
	<-w.queue
}
