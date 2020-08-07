package lexer

type asciiSet [8]uint32
type TokenType uint8

func (as *asciiSet) contains(c uint8) bool {
	return (as[c>>5] & (1 << uint(c&31))) != 0
}
