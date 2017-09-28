// Parses protobuf definitions using work queue
package parser

import (
  "log"
  "sync"
  "strings"
  "io/ioutil"
  "text/scanner"
  "path/filepath"

  "github.com/RobotStudio/choreo-svc/generate/work"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

var mts = &msgTypes{}


// fileParser type for work queue
type fileParser struct {
  filename *string
  root string
}

// Parse the file
func (f *fileParser) parse() *[]msgType {
  dat, err := ioutil.ReadFile(filepath.Join(f.root, *f.filename))
  check(err)
  return parseFileContents(&dat, f.filename)
}

// Task implements the Worker interface
func (f *fileParser) Task() {
  log.Printf("Parsing: %s\n", *f.filename)
  mts.add(f.parse())
}


// msgTypes is a collection of all msgType entities
type msgTypes struct {
  nameDepLookup sync.Map
  mts []*msgType
}

// Parses the established filename
func (m *msgTypes) add(mts *[]msgType) {
  for _,mt := range *mts {
    m.mts = append(m.mts, &mt)
  }
}


// msgType provides support for generation
type msgType struct {
  name string
  path string
  deps []string
}


func parseFileContents(dat *[]byte, fname *string) *[]msgType {
  var mss []msgType
  var s scanner.Scanner
  s.Filename = *fname
  s.Init(strings.NewReader(string(*dat)))

  secTok := false
  inMsgDef := false
  var thisMsgType msgType

  for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
    if pos := s.Position; secTok {
      thisMsgType = msgType{
        name: s.TokenText(),
        path: *fname,
      }

      log.Printf("%s: %s\n", pos, thisMsgType.name)
      secTok = false
    } else if s.Column == 1 && s.TokenText() == "message" {
      secTok = true
    } else if s.TokenText() == "{" {
      inMsgDef = true
    } else if s.TokenText() == "}" {
      mss = append(mss, thisMsgType)
      thisMsgType = msgType{}
      inMsgDef = false
    } else if inMsgDef && s.Column == 3 && s.TokenText() != "repeated" {
      if m, ok := mts.nameDepLookup.Load(thisMsgType.name); ok {
        t := m.(sync.Map)
        t.Store(s.TokenText(), true)
      } else {
        t := sync.Map{}
        t.Store(s.TokenText(), true)
        mts.nameDepLookup.Store(thisMsgType.name, t)
      }
    }
  }

  return &mss
}

// main entrypoint
func Run(maxWorkers int, root string, files *[]string) *msgTypes {
  // Create a work pool with 2 goroutines
  q := work.New(maxWorkers)

  var wg sync.WaitGroup
  wg.Add(len(*files))

  log.Println(*files)

  // Iterate over the slice of names
  for i,_ := range *files {
    f := (*files)[i]
    go func() {
      // Submit the task to be worked on. When RunTask returns we know it is
      // being handled.
      q.Run(&fileParser{filename: &f, root: root})
      wg.Done()
    }()
  }

  wg.Wait()

  // Shutdown the work pool and wait for all existing work to be completed
  q.Shutdown()

  return mts
}
