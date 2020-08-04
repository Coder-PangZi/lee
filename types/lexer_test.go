package types

import (
	"strings"
	"testing"
)

var code = ` // a
  mon1 = 3
mon2 = 2
c = a + b
while a > 1 {
	a --
}
`

func TestLexer_Parse(t *testing.T) {
	l := NewLexer(RegToken)
	l.ParseBytes(strings.NewReader(code))
}

func trimSpace(content string) {

}
