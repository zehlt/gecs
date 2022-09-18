package gecs

import (
	"sync"
)

type Task func()

type WorkerPool struct {
	numGoroutine int
	sendWork     chan Task
	wg           sync.WaitGroup
}

func (w *WorkerPool) Start(num int) {
	w.numGoroutine = num
	w.sendWork = make(chan Task)
	w.wg = sync.WaitGroup{}

	for n := num; n > 0; n-- {
		w.wg.Add(1)

		go func(n int) {
			defer w.wg.Done()

			for work := range w.sendWork {
				work()
				// log.Printf("GOROUTINE: %d TASK DONE!", n)
			}

			// log.Printf("GOUTINE: %d STOPPED", n)
		}(n)
	}
}

func (w *WorkerPool) Send(t Task) {
	w.sendWork <- t
}

// func (w *WorkerPool) WaitDone() {
// 	// for som := range w.workDone {

// 	// }
// 	// close(w.sendWork)
// }

func (w *WorkerPool) Stop() {
	close(w.sendWork)
	w.wg.Wait()
}
