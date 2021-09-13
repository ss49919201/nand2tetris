package main

import (
	"log"
	"os"
)

func main() {
	log.Println("start analyze")
	log.Printf("command line args: %s", os.Args[1:])

	// Open File
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	n := newJackAnalyzer(file)
	defer n.Close()
}
