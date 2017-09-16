// Uses the work package to generate/execute work
package main

import (
  "log"
  "sync"
  "time"

  "github.com/RobotStudio/choreo-svc/generate/parser"
  //"github.com/RobotStudio/choreo-svc/generate/generator"
)

var files = []string{
  "../../../../github.com/RobotStudio/choreo-msg/src/primitive/bool.proto",
}

// main entrypoint
func main() {
  maxWorkers := 1
  mts := parser.Run(maxWorkers, files)
}
