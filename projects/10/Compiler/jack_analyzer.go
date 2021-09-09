package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type jackAnalyzer struct {
	inputs             []*os.File
	outputFileNameBase string
}

func newJackAnalyzer(file *os.File) *jackAnalyzer {
	jackAnalyzer := new(jackAnalyzer)
	// ディレクトリであるか確認
	_, err := file.ReadDir(0)
	if err == nil {
		// ディレクトリのケース
		jackAnalyzer.outputFileNameBase = file.Name()

		file.Seek(0, 0)
		files, err := file.ReadDir(0)
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			path := filepath.Join(file.Name(), f.Name())

			// 拡張子チェック
			if jackAnalyzer.isJackFile(f.Name()) {
				log.Printf("%s is not jack file", f.Name())
				continue
			}

			input, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			// FIXME: cap
			jackAnalyzer.inputs = append(jackAnalyzer.inputs, input)
		}
	} else {
		// ファイルのケース
		jackAnalyzer.outputFileNameBase = strings.TrimRight(file.Name(), "")
		jackAnalyzer.inputs = append(jackAnalyzer.inputs, file)
	}
	return jackAnalyzer
}

func (j *jackAnalyzer) isJackFile(filePath string) bool {
	s := strings.Split(filePath, ".")
	return s[len(s)-1] != "jack"
}

func (j *jackAnalyzer) Close() {
	for _, input := range j.inputs {
		input.Close()
	}
}
