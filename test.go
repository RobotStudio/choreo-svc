package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `syntax = "proto3";
package choreo;

option go_package = "github.com/RobotStudio/choreo-msg/msg/primitive";

message BoolArray {
  repeated Bool data = 1;
}

message Bool {
  bool data = 1;
}
`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
  secondTok := false
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
    if pos := s.Position; secondTok {
		  fmt.Printf("%s: %s\n", pos, s.TokenText())
      secondTok = false
    } else if s.Column == 1 {
      if s.TokenText() == "message" {
        secondTok = true
      }
    }
	}

}

