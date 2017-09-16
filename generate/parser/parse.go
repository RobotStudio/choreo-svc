// Parses protobuf definitions using work queue
package generate

import (
  "log"
  "sync"
  "strings"

  "github.com/RobotStudio/choreo-svc/generate/work"
)

// msgTypes is a collection of all msgType entities
type msgTypes struct {
  nameDepLookup map[string][]string
  mts []msgType
}

// msgType provides support for generation
type msgType struct {
  name string
  path string
  deps []string
}

// Task implements the Worker interface
func (m *msgType) Task() {
  log.Println("PARSE: %s", m.filename)
  m.parse()
}

// Parses the established filename
func (m *msgType) parse() {
}

// main entrypoint
func Run(maxWorkers int, files []strings) {
  var mts msgTypes
  mts.mts = make(map[string][]string)

  // Create a work pool with 2 goroutines
  q := work.New(maxWorkers int)

  var wg sync.WaitGroup
  wg.Add(len(mts.mts))

  // Iterate over the slice of names
  for _, mt := range mts.mts {
    go func() {
      // Submit the task to be worked on. When RunTask returns we know it is
      // being handled.
      q.Run(&mt)
      mts.nameDepLookup[mt.name] = mt.deps
      wg.Done()
    }()
  }

  wg.Wait()

  // Shutdown the work pool and wait for all existing work to be completed
  q.Shutdown()

  return mts
}
