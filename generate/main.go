// Uses the work package to generate/execute work
package main

import (
  "log"
  "github.com/RobotStudio/choreo-svc/generate/parser"
  //"github.com/RobotStudio/choreo-svc/generate/generator"
)

var files = []string{
  "../../../../github.com/RobotStudio/choreo-msg/src/primitive/bool.proto",
}

// main entrypoint
func main() {
  maxWorkers := 2
  log.Println("Launching...")
  mts := parser.Run(maxWorkers, &files)
  log.Println(mts)
}
