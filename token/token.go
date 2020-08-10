package token

import (
	"bytes"
	"fmt"
)

type Type int

const (
	_ Type = iota
	type_beg
	TypeReserve // 保留字
	TypeBoundary
	TypeOperator
	TypeIdent
	TypeNumber
	TypeString
	type_end
)

var TypeNames = map[Type]string{
	TypeReserve:  "保留字",
	TypeBoundary: "定界符",
	TypeOperator: "操作符",
	TypeIdent:    "标识符",
	TypeNumber:   "数类型",
	TypeString:   "字符串",
}

var Reserve = map[string]struct{}{
	"var":       {},
	"true":      {},
	"false":     {},
	"null":      {},
	"undefined": {},
	"import":    {},
	"function":  {},
	"new":       {},
	"if":        {},
	"return":    {},
	"break":     {},
	"continue":  {},
	"while":     {},
	"switch":    {},
	"case":      {},
	"default":   {},
	"typeof":    {},
	"try":       {},
	"catch":     {},
	"finally":   {},
	"for":       {},
	"else":      {},
}

type token struct {
	buf  bytes.Buffer
	Type Type
	line uint32
	row  uint32
}

func newReserveToken(line uint32, row uint32) *token {
	return &token{Type: TypeReserve, line: line, row: row}
}

func newBoundaryToken(line uint32, row uint32) *token {
	return &token{Type: TypeBoundary, line: line, row: row}
}

func newOperatorToken(line uint32, row uint32) *token {
	return &token{Type: TypeOperator, line: line, row: row}
}

func newNumberToken(line uint32, row uint32) *token {
	return &token{Type: TypeNumber, line: line, row: row}
}

func newIdentifierToken(line uint32, row uint32) *token {
	return &token{Type: TypeIdent, line: line, row: row}
}

func (t *token) String() string {
	return fmt.Sprintf("[%s][%d:%d]\t%s", TypeNames[t.Type], t.line, t.row, t.buf.String())
}
