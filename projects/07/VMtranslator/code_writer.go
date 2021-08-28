package main

import (
	"fmt"
	"os"
)

type codeWriter struct {
	input  []*os.File
	output *os.File
}

func newCodeWriter(file *os.File) *codeWriter {
	return &codeWriter{}
}

func (c *codeWriter) setFileName(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	c.input = append(c.input, f)
}

func (c *codeWriter) writeArthmethic(command string) {
	switch command {
	case "add":
		fmt.Fprintln(c.output,
			`@SP
			D=M-1
			D=D+M
			M=D`,
		)
	case "sub":
		fmt.Fprintln(c.output,
			`@SP
			D=M-1
			D=D-M
			M=D`,
		)
	case "ng":
		fmt.Fprintln(c.output,
			`@SP
			D=-D
			M=D`,
		)
	case "gt":
		fmt.Fprintln(c.output,
			`@SP
			D=M-1
			D=D-M
			@TRUE
			D;JEQ
			@FALSE
			D;JNE
			(TRUE)
			@SP
			M=-1
			@NEXT
			0;JMP
			(FALSE)
			@SP
			M=0
			@NEXT
			0;JMP
			(NEXT)
			`,
		)
	case "lt":
		fmt.Fprintln(c.output,
			`@SP
			D=M-1
			D=D-M
			@TRUE
			D;JEQ
			@FALSE
			D;JNE
			(TRUE)
			@SP
			M=-1
			@NEXT
			0;JMP
			(FALSE)
			@SP
			M=0
			@NEXT
			0;JMP
			(NEXT)
			`,
		)
	case "eq":
		fmt.Fprintln(c.output,
			`@SP
			D=M-1
			D=D-M
			@TRUE
			D;JEQ
			@FALSE
			D;JNE
			(TRUE)
			@SP
			M=-1
			@NEXT
			0;JMP
			(FALSE)
			@SP
			M=0
			@NEXT
			0;JMP
			(NEXT)
			`,
		)
	case "and":
		fmt.Fprintln(c.output,
			`@SP-1
			D=M
			@SP
			D=D&M
			@SP
			M=D`,
		)
	case "or":
		fmt.Fprintln(c.output,
			`@SP-1
			D=M
			@SP
			D=D|M
			@SP
			M=D`,
		)
	case "not":
		fmt.Fprintln(c.output,
			`@SP
			D=D!
			M=D`,
		)
	}
}

func (c *codeWriter) writePushPop(command, segment string, index int) {
	switch command {
	case "push":
		fmt.Fprintln(c.output,
			`@SP
			D=D!
			M=D`,
		)
	case "pop":
		return
	case "label":
		return
	case "goto":
		return
	case "if-goto":
		return
	case "function":
		return
	case "return":
		return
	case "call":
		return
	}
}

func (c *codeWriter) transrate(segment string) (register string) {}

func (c *codeWriter) close() {
	for i := range c.input {
		c.input[i].Close()
	}
}
