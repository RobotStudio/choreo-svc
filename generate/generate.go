// Go generator for service definitions
package generate

import (
  "log"
  "sync"
  "time"

  "github.com/RobotStudio/choreo-svc/generate/work"
)

var names = []string{
  "steve",
  "bob",
  "mary",
  "therese",
  "jason",
}

// namePrinter provides special support for printing names
type namePrinter struct {
  name string
}

// Task implements the Worker interface
func (m *namePrinter) Task() {
  log.Println(m.name)
  time.Sleep(time.Second)
}

// main entrypoint
func main() {
  // Create a work pool with 2 goroutines
  p := work.New(2)

  var wg sync.WaitGroup
  wg.Add(100 * len(names))

  for i := 0; i < 100; i++ {
    // Iterate over the slice of names
    for _, name := range names {
      // Create a namePrinter and provide the specific name
      np := namePrinter{
        name: name,
      }

      go func() {
        // Submit the task to be worked on. When RunTask returns we know it is
        // being handled.
        p.Run(&np)
        wg.Done()
      }()
    }
  }

  wg.Wait()

  // Shutdown the work pool and wait for all existing work to be completed
  p.Shutdown()
}
