package lexer

import "bytes"

type token struct {
	val  bytes.Buffer
	kind TokenType
	line uint32
	row  uint32
}

func newFullToken(kind TokenType, line uint32, row uint32) *token {
	return &token{kind: kind, line: line, row: row}
}

func newToken(line uint32) *token {
	return &token{line: line}
}
