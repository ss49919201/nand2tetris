package main

import "os"

type jackTokenizer struct {
	inputs []*os.File
	output *os.File
}

func newJackTokenizer(inputs []*os.File, outputFileNameBase string) (*jackTokenizer, error) {
	jackTokenizer := new(jackTokenizer)

	jackTokenizer.inputs = inputs

	output, err := os.Create(outputFileNameBase + "T.xml")
	if err != nil {
		return nil, err
	}
	jackTokenizer.output = output

	return jackTokenizer, nil
}
