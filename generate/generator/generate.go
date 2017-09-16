// Go generator for service definitions
package generator

import (
  "log"
  "sync"
  "time"

  "github.com/RobotStudio/choreo-svc/generate/work"
)

// Provide an interface to populate work queue
type Generator interface {
  GetDepsForTypes([]string) []string
  GetTypeNames() []string
}

// Task implements the Worker interface
func (m *Template) Task() {
  log.Println(m.name)
  time.Sleep(time.Second)
}

// main entrypoint
func Run(maxWorkers int, gen Generator) {
  // Create a work pool with 2 goroutines
  p := work.New(maxWorkers)

  var wg sync.WaitGroup
  wg.Add(gen.GetTypesCount())

  // Iterate over the slice of names
  for _, name := range gen.GetTypes {
    go func() {
      // Submit the task to be worked on. When RunTask returns we know it is
      // being handled.
      p.Run(&np)
      wg.Done()
    }()
  }

  wg.Wait()

  // Shutdown the work pool and wait for all existing work to be completed
  p.Shutdown()
}
