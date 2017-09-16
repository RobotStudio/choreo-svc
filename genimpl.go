// +build ignore
// genimpl generates the choreo proto exchanges
// Adapted from https://github.com/josharian/impl (thanks)

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
  "io/ioutil"
	"os"
	"path/filepath"
	"text/template"
  "strings"
  "text/scanner"
)

const usage = `genimpl [-dir directory]

genimpl generates method stubs for recv to implement iface.

Examples:

./genimpl
./genimpl -dir $GOPATH/src/github.com/RobotStudio/choreo-msg/msg
`

const (
  msgSrc = "./vendor/github.com/RobotStudio/choreo-msg/src"
  template = "./tmpl/grpc.proto"
)

var (
	flagSrcDir = flag.String("dir", "", "package source directory, useful for vendored code")
)

type Templates struct {
  filename string
  path string
  Proto
}

type Proto struct {
  Imports []string
  Services []Service
}

type Service struct {
  Rpcs []Rpc
}

type Rpc struct {
  Call string
  Params []Param
  Res []Param
}

type Param struct {
  Type string
}

// pbtypes returns the set of methods required to implement iface.
func pbtypes(path string) ([]string, error) {
	// Parse the package and find the interface declaration.
	var pbts []string
  var s scanner.Scanner

  d, err := os.Open(path)
  fatalIfErr(err)
  defer d.Close()
  fi, err := d.Readdir(-1)
  fatalIfErr(err)
  for _, f := range fi {
    if f.Mode().IsRegular() {
      s.Init(strings.NewReader(f.Name()))
      secTok := false
      for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
        if pos := s.Position; secTok {
		      pbts = append(pbts, s.TokenText())
          secTok = false
        } else if s.Column == 1 {
          if s.TokenText() == "message" {
            secTok = true
          }
        }
      }
    }
	}
	return pbts, nil
}


var tmpl = template.Must(template.New("test").ParseFiles(tmplate))

// genStubs prints nicely formatted method stubs
// for pbts using receiver expression recv.
// If recv is not a valid receiver expression,
// genStubs will panic.
func genStubs(pbts []string) {
	var buf bytes.Buffer
	for _, pbt := range pbts {
    err := tmpl.ExecuteTemplate(&buf, pbt)
    fatalIfErr(err)
	}
}

func main() {
	flag.Parse()

	if *flagSrcDir == "" {
		if dir, err := os.Getwd(); err == nil {
      dir = filepath.Join(dir, msgSrc)
			*flagSrcDir = dir
		}
	}

	pbts, err := pbtypes(*flagSrcDir)
	if err != nil {
		fatal(err)
	}

	src := genStubs(pbts)
	fmt.Print(string(src))
}

func fatalIfErr(err interface{}) {
  if err != nil {
    fatal(err)
  }
}

func fatal(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
