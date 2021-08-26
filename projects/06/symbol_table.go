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
func newSymbolTable() *symbolTable {
	s := new(symbolTable)
	s.records = map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
	}

	return s
}

// テーブルに（symbol, address）を追加する
func (s *symbolTable) addEntry(symbol string, address int) {
	s.records[symbol] = address
}

// シンボルテーブルは与えられたシンボルを含むか否かを返す
func (s *symbolTable) contains(symbol string) bool {
	_, ok := s.records[symbol]
	return ok
}

// シンボルに対応する値を返す
func (s *symbolTable) getAddress(symbol string) int {
	v, ok := s.records[symbol]
	if !ok {
		panic("not found symbol")
	}
	return v
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
