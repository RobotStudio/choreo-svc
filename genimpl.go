// genimpl generates the choreo proto exchanges
// Adapted from https://github.com/josharian/impl (thanks)

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
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

const msgSrc = "./vendor/github.com/RobotStudio/choreo-msg/src"

var (
	flagSrcDir = flag.String("dir", "", "package source directory, useful for vendored code")
)

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

const stub = "func ({{.Recv}}) {{.Name}}" +
	"({{range .Params}}{{.Name}} {{.Type}}, {{end}})" +
	"({{range .Res}}{{.Name}} {{.Type}}, {{end}})" +
	"{\n" + "panic(\"not implemented\")" + "}\n\n"

var tmpl = template.Must(template.New("test").Parse(stub))

// genStubs prints nicely formatted method stubs
// for pbts using receiver expression recv.
// If recv is not a valid receiver expression,
// genStubs will panic.
func genStubs(pbts []string) []byte {
	var buf bytes.Buffer
	for _, pbt := range pbts {
		tmpl.Execute(&buf, meth)
	}

	pretty, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return pretty
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
