package lexer

import (
	"github.com/davecgh/go-spew/spew"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	l := Lexer{}

	const SC = `//abcdef asdw adw`
	f, err := NewFile(strings.NewReader(SC))
	if err != nil {
		t.Error(err)
	}
	l.Init(f, nil)
	spew.Dump(l.scanComment())
}
