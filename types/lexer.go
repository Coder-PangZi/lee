package types

import (
	"github.com/davecgh/go-spew/spew"
	"io"
	"regexp"
)

type Lexer struct {
	reg   *regexp.Regexp
	queue *TokenQueue
}

func NewLexer(reg string) *Lexer {
	l := Lexer{reg: regexp.MustCompile(reg)}
	return &l
}

func (l *Lexer) ParseBytes(code io.Reader) {
	buf := make([]byte, 1024)
	_, _ = code.Read(buf)
	rst := l.reg.FindAll(buf, -1)
	l.queue = NewTokenQueue(len(rst))
	for _, t := range rst {
		l.queue.Push(NewToken(0, t))
	}
	spew.Dump(rst)
}
