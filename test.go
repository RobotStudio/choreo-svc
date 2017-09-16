package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

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

func main() {
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
  secondTok := false
  inMsgTypeDef := false

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {

    if pos := s.Position; secondTok {
		  fmt.Printf("%s: %s\n", pos, s.TokenText())
      secondTok = false
    } else if s.Column == 1 && s.TokenText() == "message" {
      secondTok = true
    } else if s.TokenText() == "{" {
      inMsgTypeDef = true
    } else if s.TokenText() == "}" {
      inMsgTypeDef = false
    } else if inMsgTypeDef && s.Column == 3 {
		  //fmt.Printf("%s: %s\n", pos, s.TokenText())
      if word := s.TokenText(); word != "repeated" {
        fmt.Printf("DEP: %s: %s\n", pos, s.TokenText())
      }
    }

	}

}
