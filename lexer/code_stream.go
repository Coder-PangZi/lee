package lexer

import "io"

var (
	spaceCharAS,
	letterAs,
	numberAs,
	letterNumAS asciiSet
)

func init() {
	var ok bool
	spaceCharAS, ok = makeASCIISet(CharSpace)
	if !ok {
		panic(errAsicc)
	}
	letterAs, ok = makeASCIISet(CharLetter)
	if !ok {
		panic(errAsicc)
	}
	letterNumAS, ok = makeASCIISet(CharLetter + CharNum)
	if !ok {
		panic(errAsicc)
	}
	numberAs, ok = makeASCIISet(CharNum)
	if !ok {
		panic(errAsicc)
	}
}

type codeStream struct {
	code string
	cur  uint32
	line uint32
	row  uint32
}

func newCodeStream(code string) *codeStream {
	return &codeStream{
		code: code,
		cur:  0,
		line: 1,
		row:  1,
	}
}

func (cs *codeStream) peek(i ...uint32) (uint8, error) {
	var n uint32 = 0
	if len(i) > 0 {
		n = i[0]
	}
	if cs.cur+n < uint32(len(cs.code)) {
		return cs.code[cs.cur+n], nil
	}
	return 0, io.EOF
}

func (cs *codeStream) pop() (uint8, error) {
	defer func() {
	}()
	if cs.cur < uint32(len(cs.code)) {
		chr := cs.code[cs.cur]
		cs.cur++
		cs.row++
		if chr == '\n' {
			cs.line++
			cs.row = 1
		}
		return chr, nil
	}
	return 0, io.EOF
}

func (cs *codeStream) skipHeadSpace() error {
	var length = uint32(len(cs.code))
	for i := cs.cur; i < length; i++ {
		if !spaceCharAS.contains(cs.code[i]) {
			return nil
		}
		_ = cs.forward(1)
	}
	return io.EOF
}

func (cs *codeStream) forward(n uint32) error {
	if cs.cur+n >= uint32(len(cs.code)) {
		return io.EOF
	}
	for i := uint32(0); i < n; i++ {
		cs.cur++
		if '\n' == cs.code[i] {
			cs.line++
			cs.row = 1
		}
	}
	return nil
}
