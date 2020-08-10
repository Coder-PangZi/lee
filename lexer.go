package lee

import (
	"bytes"
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

// 读取源码中的一个 token
func (l *lexer) getToken() (*token, error) {
	err := l.cs.skipHeadSpace()
	if err != nil {
		return nil, err
	}
	chr0, err := l.cs.peek()
	if err == io.EOF {
		return nil, err
	}
	switch chr0 {
	case '/':
		chr1, err := l.cs.peek(1)
		if err != nil {
			return nil, err
		}
		// 跳过注释代码
		if chr1 == '/' {
			err = l.skipComment()
			if err != nil {
				return nil, err
			}
			return l.getToken()
		}
		fallthrough
	case '+', '-', '*', '%', '^', '>', '<', '=':
		chr1, err := l.cs.peek(1)
		if err != nil && err != io.EOF {
			return nil, io.EOF
		}
		if chr1 == '=' {
			return l.getOperator(chr0, chr1)
		}
		return l.getOperator(chr0)

	case '(', ')', '{', '}':
		t := newBoundaryToken(l.cs.line, l.cs.row)
		t.buf.WriteByte(chr0)
		_ = l.cs.forward(1)
		return t, nil
	default:
		if chr0 >= '0' && chr0 <= '9' {
			return l.getNum()
		}
		if !letterAs.contains(chr0) {
			return nil, fmt.Errorf("unknow token: %c", chr0)
		}
		buf, err := l.getName()
		if err != nil {
			return nil, err
		}
		_, ok := Reserve[buf.String()]
		if ok {
			t := newReserveToken(l.cs.line, l.cs.row)
			t.buf.Write(buf.Bytes())
			return t, nil
		}
		t := newIdentifierToken(l.cs.line, l.cs.row)
		t.buf.Write(buf.Bytes())
		return t, nil
	}
}

// 获取注释块
func (l *lexer) skipComment() error {
	var err error
	var chrN uint8
	for {
		chrN, err = l.cs.peek()
		if chrN == '\n' || err == io.EOF {
			break
		}
		_ = l.cs.forward(1)
	}
	return nil
}

// 获取计算符号
func (l *lexer) getOperator(chr ...uint8) (*token, error) {
	t := newOperatorToken(l.cs.line, l.cs.row)
	t.row = l.cs.row
	t.buf.Write(chr)
	_ = l.cs.forward(uint32(len(chr)))
	return t, nil
}

// 获取数字
func (l *lexer) getNum() (*token, error) {
	var err error
	var chrN uint8
	t := newNumberToken(l.cs.line, l.cs.row)
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
		t.buf.WriteByte(chrN)
		_ = l.cs.forward(1)
	}
	return t, nil
}

// 获取数字
func (l *lexer) getName() (bytes.Buffer, error) {
	var err error
	var chrN uint8
	var buf bytes.Buffer
	for {
		chrN, err = l.cs.peek()
		if chrN == '\n' || err == io.EOF || !letterNumAS.contains(chrN) {
			break
		}
		buf.WriteByte(chrN)
		_ = l.cs.forward(1)
	}
	return buf, err
}
