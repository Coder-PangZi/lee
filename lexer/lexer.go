package lexer

const bom = 0xFEFF

type ErrHandle func(err error)

type Lexer struct {
	*File
	ErrCnt    int
	errHandle ErrHandle
}

func (l *Lexer) error(err error) {
	if l.errHandle != nil {
		l.errHandle(err)
	}
	l.ErrCnt++
}

// Init 会初始化一个词法解析器
func (l *Lexer) Init(file *File, err ErrHandle) {
	l.File = file
	l.errHandle = err
	_ = l.next()
	if l.chr == bom {
		_ = l.next() // 忽略 bom
	}
}

// scanComment 读取注释
func (l *Lexer) scanComment() string {
	// initial '/' already consumed; s.ch == '/' || s.ch == '*'
	// 第一个 '/' 的位置
	offStart := l.offset - 1
	// 单行注释
	if l.chr == '/' {
		l.next()
		for l.chr != '\n' && l.chr >= 0 {
			l.next()
		}
	} else {
		// 多行注释
		//l.next()
		for l.chr >= 0 {
			ch := l.chr
			//l.next()
			if ch == '*' && l.chr == '/' {
				//l.next()
				break
			}
		}
	}

	lit := l.src[offStart:l.offset]

	// windows 平台下换行符可能是 '\r\n'
	// 注释的换行处 '\r' 需要移除
	//lit = stripCR(lit, lit[1] == '*')
	return string(lit)
}

func stripCR(b []byte, comment bool) []byte {
	c := make([]byte, len(b))
	i := 0
	for j, ch := range b {
		if ch != '\r' || comment && i > len("/*") && c[i-1] == '*' && j+1 < len(b) && b[j+1] == '/' {
			c[i] = ch
			i++
		}
	}
	return c[:i]
}
