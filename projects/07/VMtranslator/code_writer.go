package main

import (
	"os"
	"strconv"
	"strings"
)

type codeWriter struct {
	input                []*os.File
	output               *os.File
	labelCounter         int
	returnAddressCounter int
	function             string
}

func newCodeWriter(file *os.File) *codeWriter {
	// 出力ファイル名の決定
	fileName := strings.Split(file.Name(), ".vm")[0] + ".asm"
	if isDir(file) {
		sPath := strings.Split(file.Name(), "/")
		fileName = file.Name() + "/" + sPath[len(sPath)-1] + ".asm"
	}

	// 出力ファイルの作成
	output, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return &codeWriter{
		output:               output,
		labelCounter:         0,
		returnAddressCounter: 0,
		function:             "",
	}
}

func (c *codeWriter) updateLabelCounter() {
	c.labelCounter++
}

func (c *codeWriter) setFileName(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	c.input = append(c.input, f)
	c.function = ""
}

func (c *codeWriter) writeArthmethic(command string) {
	unique := strconv.Itoa(c.labelCounter)
	c.updateLabelCounter()

	trueLabel := "(TRUE_" + unique + ")"
	falseLabel := "(FALSE_" + unique + ")"
	nextLabel := "(NEXT_" + unique + ")"
	trueACoomand := "@TRUE_" + unique
	falseACoomand := "@FALSE_" + unique
	nextACoomand := "@NEXT_" + unique

	switch command {
	case "add":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"M=M-1\n" +
				"@SP\n" +
				"A=M\n" +
				"A=A-1\n" +
				"D=M\n" +
				"@SP\n" +
				"A=M\n" +
				"D=M+D\n" +
				"A=A-1\n" +
				"M=D\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n",
		)
	case "sub":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"M=M-1\n" +
				"@SP\n" +
				"A=M\n" +
				"A=A-1\n" +
				"D=M\n" +
				"@SP\n" +
				"A=M\n" +
				"D=D-M\n" +
				"A=A-1\n" +
				"M=D\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n",
		)
	case "neg":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"A=M-1\n" +
				"D=M\n" +
				"D=-D\n" +
				"M=D\n",
		)
	case "gt":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"M=M-1\n" +
				"@SP\n" +
				"A=M-1\n" +
				"D=M\n" +
				"@SP\n" +
				"A=M\n" +
				"D=D-M\n" +
				trueACoomand + "\n" +
				"D;JGT\n" +
				falseACoomand + "\n" +
				"0;JMP\n" +
				trueLabel + "\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=-1\n" +
				nextACoomand + "\n" +
				"0;JMP\n" +
				falseLabel + "\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=0\n" +
				nextACoomand + "\n" +
				"0;JMP\n" +
				nextLabel + "\n",
		)
	case "lt":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"M=M-1\n" +
				"@SP\n" +
				"A=M-1\n" +
				"D=M\n" +
				"@SP\n" +
				"A=M\n" +
				"D=D-M\n" +
				trueACoomand + "\n" +
				"D;JLT\n" +
				falseACoomand + "\n" +
				"0;JMP\n" +
				trueLabel + "\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=-1\n" +
				nextACoomand + "\n" +
				"0;JMP\n" +
				falseLabel + "\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=0\n" +
				nextACoomand + "\n" +
				"0;JMP\n" +
				nextLabel + "\n",
		)
	case "eq":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"M=M-1\n" +
				"@SP\n" +
				"A=M-1\n" +
				"D=M\n" +
				"@SP\n" +
				"A=M\n" +
				"D=M-D\n" +
				trueACoomand + "\n" +
				"D;JEQ\n" +
				falseACoomand + "\n" +
				"D;JNE\n" +
				trueLabel + "\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=-1\n" +
				nextACoomand + "\n" +
				"0;JMP\n" +
				falseLabel + "\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=0\n" +
				nextACoomand + "\n" +
				"0;JMP\n" +
				nextLabel + "\n",
		)
	case "and":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"M=M-1\n" +
				"@SP\n" +
				"A=M-1\n" +
				"D=M\n" +
				"@SP\n" +
				"A=M\n" +
				"D=D&M\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=D\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n",
		)
	case "or":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"M=M-1\n" +
				"@SP\n" +
				"A=M-1\n" +
				"D=M\n" +
				"@SP\n" +
				"A=M\n" +
				"D=D|M\n" +
				"@SP\n" +
				"A=M-1\n" +
				"M=D\n" +
				"@SP\n" +
				"A=M\n" +
				"M=0\n",
		)
	case "not":
		c.output.WriteString(
			"\n// " + command + "\n" +
				"@SP\n" +
				"A=M-1\n" +
				"D=M\n" +
				"D=!D\n" +
				"M=D\n",
		)
	}
}

func (c *codeWriter) writePushPop(command, segment string, index int) {
	strIndex := strconv.Itoa(index)

	switch command {
	case "push":
		switch segment {
		case "static":
			c.output.WriteString(
				"\n// " + command + " " + segment + " " + strIndex + "\n" +
					"@" + segment + "." + strIndex + "\n" +
					"D=M\n" +
					"@SP\n" +
					"A=M\n" +
					"M=D\n" +
					"@SP\n" +
					"M=M+1\n",
			)
			return
		case "constant":
			c.output.WriteString(
				"\n// " + command + " " + segment + " " + strIndex + "\n" +
					"@" + strIndex + "\n" +
					"D=A\n" +
					"@SP\n" +
					"A=M\n" +
					"M=D\n" +
					"@SP\n" +
					"M=M+1\n",
			)
			return
		case "pointer":
			target := "THIS"
			if index == 1 {
				target = "THAT"
			}
			c.output.WriteString(
				"\n// " + command + " " + segment + " " + strIndex + "\n" +
					"@" + target + "\n" +
					"D=M\n" +
					"@SP\n" +
					"A=M\n" +
					"M=D\n" +
					"@SP\n" +
					"M=M+1\n",
			)
			return
		default:
			setSegmentAddr := "D=A"
			if segment == "local" || segment == "argument" || segment == "this" || segment == "that" {
				setSegmentAddr = "D=M"
			}
			c.output.WriteString(
				"\n// " + command + " " + segment + " " + strIndex + "\n" +
					"@" + c.getRegister(segment) + "\n" +
					setSegmentAddr + "\n" +
					"@" + strIndex + "\n" +
					"A=D+A\n" +
					"D=M\n" +
					"@SP\n" +
					"A=M\n" +
					"M=D\n" +
					"@SP\n" +
					"M=M+1\n",
			)
		}
	case "pop":
		switch segment {
		case "static":
			c.output.WriteString(
				"\n// " + command + " " + segment + " " + strIndex + "\n" +
					"@SP\n" +
					"M=M-1\n" +
					"@SP\n" +
					"A=M\n" +
					"D=M\n" +
					"@" + segment + "." + strIndex + "\n" +
					"M=D\n",
			)
			return
		case "pointer":
			target := "THIS"
			if index == 1 {
				target = "THAT"
			}
			c.output.WriteString(
				"\n// " + command + " " + segment + " " + strIndex + "\n" +
					"@SP\n" +
					"M=M-1\n" +
					"@SP\n" +
					"A=M\n" +
					"D=M\n" +
					"@" + target + "\n" +
					"M=D\n",
			)
			return
		default:
			setSegmentAddr := "D=A"
			if segment == "local" || segment == "argument" || segment == "this" || segment == "that" {
				setSegmentAddr = "D=M"
			}
			c.output.WriteString(
				"\n// " + command + " " + segment + " " + strIndex + "\n" +
					"@SP\n" +
					"M=M-1\n" +
					"@" + c.getRegister(segment) + "\n" +
					setSegmentAddr + "\n" +
					"@" + strIndex + "\n" +
					"D=D+A\n" +
					"@R13\n" +
					"M=D\n" +
					"@SP\n" +
					"A=M\n" +
					"D=M\n" +
					"@R13\n" +
					"A=M\n" +
					"M=D\n",
			)
		}
	}
}

func (c *codeWriter) writeLabel(label string) {
	if c.function != "" {
		label = c.function + "$" + label
	}
	c.output.WriteString(
		"\n// " + "label " + label + "\n" +
			"(" + label + ")" + "\n",
	)
}

func (c *codeWriter) writeGoto(label string) {
	if c.function != "" {
		label = c.function + "$" + label
	}
	c.output.WriteString(
		"\n// " + "goto " + label + "\n" +
			"@" + label + "\n" +
			"0;JMP\n",
	)
}

func (c *codeWriter) writeIf(label string) {
	if c.function != "" {
		label = c.function + "$" + label
	}
	c.output.WriteString(
		"\n// " + "if-goto " + label + "\n" +
			"@SP" + "\n" +
			"M=M-1" + "\n" +
			"@SP" + "\n" +
			"A=M" + "\n" +
			"D=M" + "\n" +
			// Dが0でなければJumpする
			"@" + label + "\n" +
			"D;JNE\n",
	)
}

func (c *codeWriter) writeCall(functionName string, numArgs int) {
	returnAddress := "RETURN" + strconv.Itoa(c.returnAddressCounter)
	c.returnAddressCounter++

	var stashedFunction string
	if c.function != "" {
		stashedFunction = c.function
		c.function = ""
	}

	c.output.WriteString(
		"\n// ------------ start call " + functionName + strconv.Itoa(numArgs) + "-------------- \n",
	)
	c.output.WriteString(
		"\n// " + "call " + functionName + strconv.Itoa(numArgs) + "\n",
	)
	c.output.WriteString(
		"// push return address\n" +
			"@" + returnAddress + "\n" +
			"D=A\n" +
			"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n",
	)
	c.output.WriteString(
		"\n// push local \n" +
			"@LCL\n" +
			"D=M\n" +
			"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n",
	)
	c.output.WriteString(
		"\n// push argument \n" +
			"@ARG\n" +
			"D=M\n" +
			"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n",
	)
	c.output.WriteString(
		"\n// push this \n" +
			"@THIS\n" +
			"D=M\n" +
			"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n",
	)
	c.output.WriteString(
		"\n// push that \n" +
			"@THAT\n" +
			"D=M\n" +
			"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n",
	)
	c.output.WriteString(
		"\n// ARG=SP-n-5\n" +
			"@SP\n" +
			"D=M\n" +
			"@5\n" +
			"D=D-A\n" +
			"@" + strconv.Itoa(numArgs) + "\n" +
			"D=D-A\n" +
			"@ARG\n" +
			"M=D\n",
	)
	c.output.WriteString(
		"// LCL=SP\n" +
			"@SP\n" +
			"D=M\n" +
			"@LCL\n" +
			"M=D\n",
	)
	c.writeGoto(functionName)
	c.writeLabel(returnAddress)
	c.output.WriteString(
		"\n// ------------ end call " + functionName + strconv.Itoa(numArgs) + "-------------- \n",
	)

	if stashedFunction != "" {
		c.function = stashedFunction
	}
}

func (c *codeWriter) writeReturn() {
	c.output.WriteString(
		"\n// " + "return\n",
	)
	c.output.WriteString(
		"\n// FRAME = LCL\n" +
			"@LCL\n" +
			"D=M\n" +
			"@R7\n" +
			"M=D\n" +
			"\n// RET = *(FRAME-5)\n" +
			"@5\n" +
			"D=A\n" +
			"@R7\n" +
			"A=M-D\n" +
			"D=M\n" +
			"@R14\n" +
			"M=D\n",
	)
	c.output.WriteString("\n// *ARG=pop()")
	c.writePushPop("pop", "argument", 0)
	c.output.WriteString(
		"\n// SP=ARG+1\n" +
			"@ARG\n" +
			"D=M+1\n" +
			"@SP\n" +
			"M=D\n" +
			"\n// THAT=*(FRAME-1)\n" +
			"@R7\n" +
			"D=M\n" +
			"@1\n" +
			"A=D-A\n" +
			"D=M\n" +
			"@THAT\n" +
			"M=D\n" +
			"\n// THIS=*(FRAME-2)\n" +
			"@R7\n" +
			"D=M\n" +
			"@2\n" +
			"A=D-A\n" +
			"D=M\n" +
			"@THIS\n" +
			"M=D\n" +
			"\n// ARG=*(FRAME-3)\n" +
			"@R7\n" +
			"D=M\n" +
			"@3\n" +
			"A=D-A\n" +
			"D=M\n" +
			"@ARG\n" +
			"M=D\n" +
			"\n// LCL=*(FRAME-4)\n" +
			"@R7\n" +
			"D=M\n" +
			"@4\n" +
			"A=D-A\n" +
			"D=M\n" +
			"@LCL\n" +
			"M=D\n",
	)
	c.output.WriteString(
		"\n// goto RET\n" +
			"@R14\n" +
			"A=M\n" +
			"0;JMP\n",
	)
}

func (c *codeWriter) writeFunction(functionName string, numLocals int) {
	c.function = functionName
	c.output.WriteString("\n// ------------- Declare Function -------------\n")
	c.output.WriteString(
		"\n// " + "function " + functionName + " " + strconv.Itoa(numLocals) + "\n" +
			"(" + functionName + ")" + "\n",
	)
	for i := 0; i < numLocals; i++ {
		c.output.WriteString("\t")
		c.writePushPop("push", "constant", 0)
	}
}

func (c *codeWriter) writeBootload() {
	c.output.WriteString("\n// ------------- Start Bootload -------------\n")
	c.output.WriteString(
		"@256" + "\n" +
			"D=A" + "\n" +
			"@SP" + "\n" +
			"M=D" + "\n",
	)
	c.writeCall("Sys.init", 0)
	c.output.WriteString("// ------------- End Bootload -------------\n")
}

func (c *codeWriter) getRegister(segment string) string {
	switch segment {
	case "local":
		return "LCL"
	case "argument":
		return "ARG"
	case "this":
		return "THIS"
	case "that":
		return "THAT"
	case "temp":
		return "R5"
	default:
		return segment
	}
}

func (c *codeWriter) close() {
	for i := range c.input {
		c.input[i].Close()
	}
}
