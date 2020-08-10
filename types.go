package lee

type (
	asciiSet      [8]uint32
	SymbolTableID uint32
	SymbolID      uint32
	TypeSymbol    uint32
	TypeOp        uint32
)

func (as *asciiSet) contains(c uint8) bool {
	return (as[c>>5] & (1 << uint(c&31))) != 0
}
