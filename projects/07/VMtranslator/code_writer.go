package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type codeWriter struct {
	input           []*os.File
	output          *os.File
	tmpSegmentCount int
}

func newCodeWriter(file *os.File) *codeWriter {
	input := make([]*os.File, 0)
	if isDir(file) {
		files, err := file.ReadDir(0)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			path := filepath.Join(file.Name(), f.Name())
			if f.Type().IsDir() {
				panic("can't translate directory")
			}
			inputFile, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			input = append(input, inputFile)
		}
	}
	output, err := os.Create(file.Name())
	if err != nil {
		panic(err)
	}

	return &codeWriter{
		tmpSegmentCount: 5,
		input:           input,
		output:          output,
	}
}

func (c *codeWriter) updateTmpSegmentCount() {
	if c.tmpSegmentCount == 12 {
		return
	}
	c.tmpSegmentCount++
}

func (c *codeWriter) setFileName(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	c.input = append(c.input, f)
}

func (c *codeWriter) writeArthmethic(command string) {
	unique := strconv.Itoa(int(time.Now().Unix()))
	trueLabel := "(TRUE_" + unique + ")"
	falseLabel := "(FALSE_" + unique + ")"
	nextLabel := "(NEXT_" + unique + ")"
	trueACoomand := "@TRUE_" + unique
	falseACoomand := "@FALSE_" + unique
	nextACoomand := "@NEXT_" + unique

	switch command {
	case "add":
		fmt.Fprintln(c.output,
			"@SP\n",
			"D=M-1\n",
			"D=D+M\n",
			"M=D",
		)
	case "sub":
		fmt.Fprintln(c.output,
			"@SP\n",
			"D=M-1\n",
			"D=D-M\n",
			"M=D",
		)
	case "ng":
		fmt.Fprintln(c.output,
			"@SP\n",
			"D=-D\n",
			"M=D",
		)
	case "gt":
		fmt.Fprintln(c.output,
			"@SP\n",
			"D=M-1\n",
			"D=D-M\n",
			trueACoomand+"\n",
			"@D;JEQ\n",
			falseACoomand+"\n",
			"D;JNE\n",
			trueLabel+"\n",
			"@SP\n",
			"M=-1\n",
			nextACoomand+"\n",
			"0;JMP\n",
			falseLabel+"\n",
			"@SP\n",
			"M=0\n",
			nextACoomand+"\n",
			"0;JMP\n",
			nextLabel,
		)
	case "lt":
		fmt.Fprintln(c.output,
			"@SP\n",
			"D=M-1\n",
			"D=D-M\n",
			trueACoomand+"\n",
			"D;JEQ\n",
			falseACoomand+"\n",
			"D;JNE\n",
			trueLabel+"\n",
			"@SP\n",
			"M=-1\n",
			nextACoomand+"\n",
			"0;JMP\n",
			falseLabel+"\n",
			"@SP\n",
			"M=0\n",
			nextACoomand+"\n",
			"0;JMP\n",
			nextLabel,
		)
	case "eq":
		fmt.Fprintln(c.output,
			"@SP\n",
			"D=M-1\n",
			"D=D-M\n",
			trueACoomand+"\n",
			"D;JEQ\n",
			falseACoomand+"\n",
			"D;JNEP\n",
			trueLabel+"\n",
			"@SP\n",
			"M=-1\n",
			nextACoomand+"\n",
			"0;JMP\n",
			falseLabel+"\n",
			"@SP\n",
			"M=0\n",
			nextACoomand+"\n",
			"0;JMP\n",
			nextLabel,
		)
	case "and":
		fmt.Fprintln(c.output,
			"@SP-1\n",
			"D=M\n",
			"@SP\n",
			"D=D&M\n",
			"@SP\n",
			"M=D",
		)
	case "or":
		fmt.Fprintln(c.output,
			"@SP-1\n",
			"D=M\n",
			"@SP\n",
			"D=D|M\n",
			"@SP\n",
			"M=D",
		)
	case "not":
		fmt.Fprintln(c.output,
			"@SP\n",
			"D=D!\n",
			"M=D",
		)
	}
}

func (c *codeWriter) writePushPop(command, segment string, index int) {
	strIndex := strconv.Itoa(index)

	switch command {
	case "push":
		switch segment {
		case "static":
			fmt.Fprintln(c.output,
				"@"+segment+"."+strIndex+"\n",
				"D=M\n",
				"@SP\n",
				"M=D",
			)
			return
		case "constant":
			fmt.Fprintln(c.output,
				"@"+strIndex+"\n",
				"D=A\n",
				"@SP\n",
				"M=D",
			)
			return
		case "pointer":
			target := "THIS"
			if index == 1 {
				target = "THAT"
			}
			fmt.Fprintln(c.output,
				"@"+target+"\n",
				"D=M",
				"@SP\n",
				"M=D\n",
			)
			return
		default:
			fmt.Fprintln(c.output,
				"@"+segment+"\n",
				"D=M+"+strIndex+"\n",
				"@SP\n",
				`M=D`,
			)
		}
	case "pop":
		switch segment {
		case "static":
			fmt.Fprintln(c.output,
				"@SP\n",
				"D=M\n",
				"@"+segment+"."+strIndex+"\n",
				"M=D",
			)
			c.updateTmpSegmentCount()
			return
		case "pointer":
			target := "THIS"
			if index == 1 {
				target = "THAT"
			}
			fmt.Fprintln(c.output,
				"@SP\n",
				"D=M\n",
				"@"+target+"\n",
				"M=D",
			)
			return
		default:
			fmt.Fprintln(c.output,
				"@SP\n",
				"D=M",
				"@"+segment+"\n",
				"A=A+"+strIndex+"\n",
				"M=D",
			)
		}
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

func (c *codeWriter) close() {
	for i := range c.input {
		c.input[i].Close()
	}
}
