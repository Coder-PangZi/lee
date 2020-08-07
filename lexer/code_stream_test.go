package lexer

import (
	"testing"
)

func TestRune(t *testing.T) {
	//st, ed := 0, 0x80
	//for tgt := st; tgt < 1; tgt++ {
	//	for chr := st; chr < ed; chr++ {
	//		t.Logf("%08b\t", chr)
	//	}
	//}
	store := uint32(0)

	for chr := 'a' - 1; chr < 'a'+33; chr++ {
		store |= 1 << uint32(chr&31)
		t.Logf("%c\t%032b\t%032b\t%032b", chr, uint32(chr&31), 1<<uint32(chr&31), store)
	}
}
