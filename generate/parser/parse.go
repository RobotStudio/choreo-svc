// Parses protobuf definitions using work queue
package parser

import (
  "log"
  "sync"

  "github.com/RobotStudio/choreo-svc/generate/work"
)

var mts = &msgTypes{
  nameDepLookup: make(map[string][]string),
}

// fileParser type for work queue
type fileParser struct {
  filename *string
}

// Task implements the Worker interface
func (f *fileParser) Task() {
  log.Printf("Parsing: %s\n", *f.filename)
  mts.add(f.filename)
}

// msgTypes is a collection of all msgType entities
type msgTypes struct {
  nameDepLookup map[string][]string
  mts []*msgType
}

// Parses the established filename
func (m *msgTypes) add(file *string) {
  mt := &msgType{}
  m.mts = append(m.mts, mt)
  m.nameDepLookup[mt.name] = mt.deps
}

// msgType provides support for generation
type msgType struct {
  name string
  path string
  deps []string
}


// main entrypoint
func Run(maxWorkers int, files *[]string) *msgTypes {
  // Create a work pool with 2 goroutines
  q := work.New(maxWorkers)

  var wg sync.WaitGroup
  wg.Add(len(*files))

  // Iterate over the slice of names
  for _, file := range *files {
    go func(f *string) {
      // Submit the task to be worked on. When RunTask returns we know it is
      // being handled.
      log.Println("Sending: ", *f)
      q.Run(&fileParser{filename: f})
      wg.Done()
    }(&file)
  }

  wg.Wait()

  // Shutdown the work pool and wait for all existing work to be completed
  q.Shutdown()

  return mts
}
