// Work manages a pool of go routines to perform work
package work

import "sync"

// Worker interface is implemented by types that want to use the work pool
type Worker interface {
  Task()
}

// Pool provides a pool of goroutines that can execute any Worker tasks
type Pool struct {
  tasks chan Worker
  wg sync.WaitGroup
}

// New creates a new work pool
func New(maxWorkers int) *Pool {
  p := Pool{
    tasks: make(chan Worker),
  }

  p.wg.Add(maxWorkers)
  for i := 0; i < maxWorkers; i++ {
    go func() {
      for w := range p.tasks {
        w.Task()
      }
      p.wg.Done()
    }()
  }

  return &p
}

// Run submits work to the pool
func (p *Pool) Run(w Worker) {
  p.tasks <- w
}

// Shutdown waits for all the goroutines to shutdown
func (p *Pool) Shutdown() {
  close(p.tasks)
  p.wg.Wait()
}
