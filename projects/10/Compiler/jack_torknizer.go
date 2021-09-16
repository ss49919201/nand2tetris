package main

import (
	"bufio"
	"os"
)

type jackTokenizer struct {
	input   *os.File
	inputs  []*os.File
	output  *os.File
	token   string
	scanner *bufio.Scanner
}

func newJackTokenizer(inputs []*os.File, outputFileNameBase string) (*jackTokenizer, error) {
	jackTokenizer := new(jackTokenizer)

	jackTokenizer.inputs = inputs
	jackTokenizer.input = inputs[0]
	jackTokenizer.scanner = bufio.NewScanner(inputs[0])

	output, err := os.Create(outputFileNameBase + "T.xml")
	if err != nil {
		return nil, err
	}
	jackTokenizer.output = output

	return jackTokenizer, nil
}

// 現在のファイルにテキストは存在するか
func (j *jackTokenizer) hasMoreText() bool {
	return false
}

// 行から1文字ずつ取り出し何らかのトークンに一致するか
func (j *jackTokenizer) hasMoreTokens() bool {
	return false
}
