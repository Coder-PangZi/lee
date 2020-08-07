package lexer

import (
	"fmt"
	"io"
)

type lexer struct {
	tokens []*token
	cs     *codeStream
	cursor int32
	err    error
}

func NewLexer(text string) *lexer {
	return &lexer{
		tokens: []*token{},
		cs:     newCodeStream(text),
		cursor: 0,
	}
}

func (l *lexer) scan() {

}

func (l *lexer) getToken() (*token, error) {
	err := l.cs.skipHeadSpace()
	if err == io.EOF {
		return nil, err
	}
	chr0, err := l.cs.peek()
	if err == io.EOF {
		return nil, err
	}
	switch chr0 {
	case '/':
		chr1, err := l.cs.peek(1)
		if err == nil && chr1 == '/' {
			return l.getComment()
		}
		if err != nil && err != io.EOF {
			return nil, io.EOF
		}
		fallthrough
	case '+':
		fallthrough
	case '-':
		fallthrough
	case '*':
		fallthrough
	case '%':
		fallthrough
	case '^':
		return l.getCompute(chr0)
	case '$':
		_ = l.cs.forward(1)
		t := newFullToken(TokenTypeVar, l.cs.line, l.cs.row)
		t.val.WriteByte('$')
		return t, l.getName(t)
	case '>':
		fallthrough
	case '<':
		chr1, err := l.cs.peek(1)
		if err != nil {
			return nil, err
		}
		t := newFullToken(TokenTypeCmp, l.cs.line, l.cs.row)
		t.val.WriteByte(chr0)
		n := uint32(1)
		if chr1 == '=' {
			n++
			t.val.WriteByte('=')
		}
		_ = l.cs.forward(n)
		return t, nil
	case '=':
		chr1, err := l.cs.peek(1)
		if err != nil {
			return nil, err
		}
		if chr1 == '=' {
			t := newFullToken(TokenTypeCmp, l.cs.line, l.cs.row)
			t.val.Write([]byte{'=', '='})
			_ = l.cs.forward(2)
			return t, nil
		}
		_ = l.cs.forward(1)
		t := newFullToken(TokenTypeAssign, l.cs.line, l.cs.row)
		t.val.WriteByte('=')
		return t, nil
	case '(':
		fallthrough
	case ')':
		fallthrough
	case '{':
		fallthrough
	case '}':
		t := newFullToken(TokenTypeSymbol, l.cs.line, l.cs.row)
		t.val.WriteByte(chr0)
		_ = l.cs.forward(1)
		return t, nil
	default:
		if chr0 >= '0' && chr0 <= '9' {
			return l.getNum()
		} else if letterAs.contains(chr0) {
			t := newFullToken(TokenTypeVar, l.cs.line, l.cs.row)
			t.val.WriteByte(chr0)
			_ = l.cs.forward(1)
			return t, l.getName(t)
		}
		return nil, fmt.Errorf("unknow token: %08b", chr0)
	}
}

// 获取注释块
func (l *lexer) getComment() (*token, error) {
	var err error
	var chrN uint8
	t := newToken(l.cs.line)
	t.row = l.cs.line
	t.kind = TokenTypeComment
	for {
		chrN, err = l.cs.peek()
		if chrN == '\n' || err == io.EOF {
			break
		}
		t.val.WriteByte(chrN)
		_ = l.cs.forward(1)
	}
	return t, nil
}

// 获取计算符号
func (l *lexer) getCompute(chr uint8) (*token, error) {
	t := newToken(l.cs.line)
	t.kind = TokenTypeCompute
	t.row = l.cs.row
	t.val.WriteByte(chr)
	_ = l.cs.forward(1)
	return t, nil
}

// 获取注释块
func (l *lexer) getVar() (*token, error) {
	var err error
	var chrN uint8
	chrN, err = l.cs.peek(1)
	if err != nil {
		return nil, err
	}
	if !letterAs.contains(chrN) {
		return nil, fmt.Errorf("wrong type of var name $%c", chrN)
	}
	t := newToken(l.cs.line)
	t.kind = TokenTypeVar
	t.row = l.cs.row
	// 先将 $ 写入
	t.val.WriteByte('$')
	_ = l.cs.forward(1)

	for {
		chrN, err = l.cs.peek()
		if chrN == '\n' || err == io.EOF || !letterNumAS.contains(chrN) {
			break
		}
		t.val.WriteByte(chrN)
		_ = l.cs.forward(1)
	}
	return t, nil
}

// 获取数字
func (l *lexer) getNum() (*token, error) {
	var err error
	var chrN uint8
	t := newToken(l.cs.line)
	t.kind = TokenTypeConst
	t.row = l.cs.row
	pointFlag := true
	for {
		chrN, err = l.cs.peek()
		if chrN == '\n' || err == io.EOF || !numberAs.contains(chrN) {
			if chrN == '.' && pointFlag {
				pointFlag = false
			} else {
				break
			}
		}
		t.val.WriteByte(chrN)
		_ = l.cs.forward(1)
	}
	return t, nil
}

// 获取数字
func (l *lexer) getName(t *token) error {
	var err error
	var chrN uint8
	for {
		chrN, err = l.cs.peek()
		if chrN == '\n' || err == io.EOF || !letterNumAS.contains(chrN) {
			break
		}
		t.val.WriteByte(chrN)
		_ = l.cs.forward(1)
	}
	return nil
}
