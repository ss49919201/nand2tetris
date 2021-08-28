package main

import "os"

type codeWriter struct {
	file *os.File
}

func newCodeWriter() *codeWriter {
	return &codeWriter{}
}

func (c *codeWriter) setFileName(fileName string) {

}

func (c *codeWriter) writeArthmethic(command string) {

}

func (c *codeWriter) writePushPop(command, segment string, index int) {

}

func (c *codeWriter) close() {}
