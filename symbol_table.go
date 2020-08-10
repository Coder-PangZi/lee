package lee

import "unsafe"

type symbol struct {
	id   SymbolID
	kind TypeSymbol
	val  unsafe.Pointer
}

type symbolTable struct {
	p        *symbolTable
	id       SymbolTableID
	children map[SymbolTableID]*symbolTable
}
