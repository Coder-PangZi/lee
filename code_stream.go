package lee

import "io"

var (
	spaceCharAS,
	letterAs,
	numberAs,
	letterNumAS asciiSet
)

const (
	CharsSpace  = " \r\n\t"
	CharsLetter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsNum    = "0123456789"
)

func init() {
	var ok bool
	spaceCharAS, ok = makeASCIISet(CharsSpace)
	if !ok {
		panic(errAscii)
	}
	letterAs, ok = makeASCIISet(CharsLetter)
	if !ok {
		panic(errAscii)
	}
	letterNumAS, ok = makeASCIISet(CharsLetter + CharsNum)
	if !ok {
		panic(errAscii)
	}
	numberAs, ok = makeASCIISet(CharsNum)
	if !ok {
		panic(errAscii)
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

// 读取指定偏移量位置的字符串
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

// 跳过平开头的空白字符
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

// 指针前进
func (cs *codeStream) forward(u ...uint32) error {
	var n uint32 = 1
	if len(u) > 0 {
		n = u[0]
	}
	if cs.cur+n >= uint32(len(cs.code)) {
		return io.EOF
	}
	for i := uint32(0); i < n; i++ {
		if cs.code[cs.cur] == '\n' {
			cs.line++
			cs.row = 1
		}
		cs.row++
		cs.cur++
	}
	return nil
}
