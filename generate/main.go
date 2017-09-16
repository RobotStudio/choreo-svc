// Uses the work package to generate/execute work
package main

import (
  "os"
  "log"
  "flag"
  "strings"
  "path/filepath"
  "github.com/RobotStudio/choreo-svc/generate/parser"
  //"github.com/RobotStudio/choreo-svc/generate/generator"
)

var (
  flagSrcDir = flag.String("dir", "",
      "package source directory, useful for vendored coded code")

  root = filepath.Join(
    "vendor",
    "github.com",
    "RobotStudio",
    "choreo-msg",
    "src",
  )

  files []string
)

func foundPath(path string, f os.FileInfo, err error) error {
  rpath, err := filepath.Rel(*flagSrcDir, path)
  if !f.IsDir() && strings.HasSuffix(rpath, ".proto") {
    log.Printf("Found %s", rpath)
    files = append(files, rpath)
  }
  return nil
}

// main entrypoint
func main() {
  flag.Parse()
  if *flagSrcDir == "" {
    if dir, err := os.Getwd(); err == nil {
      dir = filepath.Join(dir, root)
      *flagSrcDir = dir
    }
  }

  err := filepath.Walk(*flagSrcDir, foundPath)
  if err != nil {
    panic(err)
  }

  maxWorkers := 2
  log.Println("Launching...")
  mts := parser.Run(maxWorkers, *flagSrcDir, &files)
  log.Println(mts)
}
