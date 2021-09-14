package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type jackAnalyzer struct {
	inputs             []*os.File
	outputFileNameBase string
}

func newJackAnalyzer(file *os.File) (*jackAnalyzer, error) {
	jackAnalyzer := new(jackAnalyzer)
	// ディレクトリであるか確認
	_, err := file.ReadDir(0)
	if err == nil {
		// ディレクトリのケース
		jackAnalyzer.outputFileNameBase = file.Name()

		file.Seek(0, io.SeekStart)
		files, err := file.ReadDir(0)
		if err != nil {
			return nil, err
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
				return nil, err
			}
			// FIXME: cap
			jackAnalyzer.inputs = append(jackAnalyzer.inputs, input)
		}
	} else {
		// ファイルのケース
		jackAnalyzer.outputFileNameBase = strings.TrimRight(file.Name(), "")
		jackAnalyzer.inputs = append(jackAnalyzer.inputs, file)
	}
	return jackAnalyzer, nil
}

func (j *jackAnalyzer) isJackFile(filePath string) bool {
	s := strings.Split(filePath, ".")
	return s[len(s)-1] != "jack"
}

func (j *jackAnalyzer) close() {
	for _, input := range j.inputs {
		input.Close()
	}
}

func (j *jackAnalyzer) analyze() error {
	_, _ = newJackTokenizer(j.inputs, j.outputFileNameBase)
	return errors.New("not implement")
}
