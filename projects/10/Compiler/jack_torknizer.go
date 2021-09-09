package main

import "os"

type jackTokenizer struct {
	inputs []*os.File
	output *os.File
}

func newJackTokenizer(inputs []*os.File, outputFileNameBase string) *jackTokenizer {
	jackTokenizer := new(jackTokenizer)

	jackTokenizer.inputs = inputs

	output, err := os.Create(outputFileNameBase + "T.xml")
	if err != nil {
		panic(err)
	}
	jackTokenizer.output = output

	return jackTokenizer
}
