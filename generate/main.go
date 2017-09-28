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
  flookup = make(map[string]bool)
)

func foundPath(path string, f os.FileInfo, err error) error {
  if err != nil {
    panic(err)
  }

  rpath, err := filepath.Rel(*flagSrcDir, path)
  if err != nil {
    panic(err)
  }

  if !f.IsDir() && strings.HasSuffix(rpath, ".proto") {
    log.Printf("Found %s", rpath)
    // if path isn't in list, add it
    if _, ok := flookup[rpath]; !ok {
      flookup[rpath] = true
      files = append(files, rpath)
    }
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
    } else {
      panic(err)
    }
  }

  err := filepath.Walk(*flagSrcDir, foundPath)
  if err != nil {
    panic(err)
  }

  maxWorkers := 2
  log.Println("Launching...")
  parser.Run(maxWorkers, *flagSrcDir, &files)
}
