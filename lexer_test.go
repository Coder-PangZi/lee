package lee

import (
	"github.com/davecgh/go-spew/spew"
	"io"
	"testing"
)

const Code = `
// comment
// test
var a = 3
func sum(a)
{
	return a + 3
}
`

func TestNewLexer(t *testing.T) {
	l := NewLexer(Code)
	for {
		token, err := l.getToken()
		if err == io.EOF {
			println("词法分析结束")
			break
		}
		if err != nil {
			println("出错", l.cs.line, ":", l.cs.row, " ", err.Error())
			spew.Dump(l.cs)
			break
		}
		println(token.String())
	}
}
