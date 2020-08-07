package parser

import "unsafe"

type symbol struct {
	id   SymbolID
	kind SymbolType
	val  unsafe.Pointer
}

type symbolTable struct {
	p        *symbolTable
	id       SymbolTableID
	children map[SymbolTableID]*symbolTable
}
