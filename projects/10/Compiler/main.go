package main

import (
	"fmt"
	"log"
	"os"
)

func exitAnalyze(err error) {
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	log.Println(err)
	fmt.Println("fail analyze😢")
	os.Exit(1)
}

func main() {
	fmt.Println("start analyze😶")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")

	log.Printf("command line args: %s", os.Args[1:])

	// Open File
	file, err := os.Open(os.Args[1])
	if err != nil {
		exitAnalyze(err)
	}
	n, err := newJackAnalyzer(file)
	if err != nil {
		exitAnalyze(err)
	}
	defer n.close()

	if err := n.analyze(); err != nil {
		exitAnalyze(err)
	}

	fmt.Println("complete analyze😋")
}
