package main

import (
	"fmt"
	"os"
)

// symbolTableモジュール
type symbolTable struct {
	records map[string]int
}

// シンボルテーブルの初期化
// func newSymbolTable() *symbolTable {}

// テーブルに（symbol, address）を追加する
// func (s *symbolTable) addEntry(symbol string, address int) {}

// シンボルテーブルは与えられたシンボルを含むか否かを返す
// func (s *symbolTable) contains(symbol string) bool {}

// シンボルに対応する値を返す
// func (s *symbolTable) getAddress(symbol string) int {}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
