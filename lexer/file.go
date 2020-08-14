package lexer

import (
	"errors"
	"io"
	"io/ioutil"
	"sync"
	"unicode/utf8"
)

var EOF = errors.New("EOF")

type File struct {
	file   io.Reader
	src    []byte
	mutex  sync.Mutex
	name   string
	line   int
	column int
	offset int
	lines  []int
	size   int
	chr    rune
	Position
}

func NewFile(reader io.Reader) (f *File, err error) {
	f = &File{}
	f.file = reader
	f.src, err = ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	f.mutex = sync.Mutex{}
	f.size = len(f.src)
	f.Position = Position{
		Offset: 0,
		Line:   1,
		Column: 1,
	}
	return f, nil
}

// next 获取下一个 unicode 字符，放入 l.chr 中
// offset 向前移动对应的距离
// s.chr < 0 表示文件读取完成
func (f *File) next() error {
	if f.offset >= len(f.src) {
		f.offset = len(f.src)
		if f.chr == '\n' {
			f.AddLine(f.offset)
		}
		f.chr = -1
		return EOF
	}
	if f.chr == '\n' {
		f.AddLine(f.offset)
	}
	r, w := rune(f.src[f.offset]), 1
	switch {
	case r == 0:
		return errors.New(f.Position.String())
	case r >= utf8.RuneSelf:
		r, w = utf8.DecodeRune(f.src[f.offset:])
		if r == utf8.RuneError && w == 1 {
			return errors.New("illegal UTF-8 encoding")
		} else if r == bom && f.offset > 0 {
			return errors.New("illegal byte order mark")
		}
	}
	f.offset += w
	f.chr = r
	return nil
}

// Position 用于记录每个 Token 的位置信息
type Position struct {
	Filename string // 文件名
	Offset   int    // 位置，从 0 开始
	Line     int    // 行号，从 1 开始
	Column   int    // 列号，从 1 开始
}

func (pos Position) String() string {
	s := pos.Filename
	if s != "" {
		s += ":"
	}
	s += int2str(pos.Line)
	if pos.Column != 0 {
		s += int2str(pos.Column)
	}
	if s == "" {
		s = "-"
	}
	return s
}

func (f *File) peek() byte {
	if f.offset < len(f.src) {
		return f.src[f.offset]
	}
	return 0
}

func (f *File) read(start, end int) []byte {
	if end < len(f.src) && end > 0 {
		return f.src[start:end]
	}
	if end == 0 {
		return f.src[start:]
	}
	return nil
}

// 数字转字符串
func int2str(src int) string {
	neg := false
	if src < 0 {
		src = -src
		neg = true
	}
	buf := [64 + 1]byte{}
	i := len(buf) - 1
	for ; src > 0; i-- {
		q := src / 10
		buf[i] = byte(src - q*10 + '0')
		src = q
	}
	if neg {
		buf[i] = '-'
		i--
	}
	return string(buf[i+1:])
}

func (f *File) AddLine(offset int) {
	f.mutex.Lock()
	if i := len(f.lines); (i == 0 || f.lines[i-1] < offset) && offset < f.size {
		f.lines = append(f.lines, offset)
		f.Line++
		f.Column = 0
	}
	f.mutex.Unlock()
}

func (f *File) Size() int {
	return f.size
}
