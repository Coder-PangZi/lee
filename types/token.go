package types

type Token struct {
	line  int32
	buf   []byte
	tType byte
}

func (b *Token) Type() byte {
	return b.tType
}

func (b *Token) SetType(tType byte) {
	b.tType = tType
}

func NewToken(line int32, buf []byte) *Token {
	lt := Token{line: line, buf: buf}
	lt.trim()
	return &lt
}

func (b *Token) String() string {
	return string(b.buf)
}

func (b *Token) Line() int32 {
	return b.line
}

func (b *Token) trim() {
	l := len(b.buf)
	start, end := 0, l
	flagStart, flagEnd := true, true
	for i := 0; i < l; i++ {
		if flagStart && (b.buf[i] != ' ' && b.buf[i] != '\t' && b.buf[i] != '\n') {
			start = i
			flagStart = false
		}
		curTail := l - i - 1
		if flagEnd && (b.buf[curTail] != ' ' && b.buf[curTail] != '\t' && b.buf[curTail] != '\n') {
			end = curTail
			flagEnd = false
		}
	}
	b.buf = b.buf[start : end+1]
}
