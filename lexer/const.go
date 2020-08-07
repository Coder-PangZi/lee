package lexer

import "fmt"

var (
	keyWord = []string{
		"$", "true", "false", "null", "undefined",
		"import", "function", "new", "if", "return",
		"break", "continue", "while", "switch", "case",
		"default", "typeof", "try", "catch", "finally",
		"throw", "for", "else",
	}
)

const (
	CharSpace                      = " \r\n\t"
	CharLetter                     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharNum                        = "0123456789"
	SpaceChar                      = " \r\n\t"
	RuneSelf                       = 0x80
	TokenLineSizeDefault           = 20
	TokenTypeComment     TokenType = 1
	TokenTypeVar         TokenType = 2
	TokenTypeConst       TokenType = 2
	TokenTypeOp          TokenType = 3
	TokenTypeCmp         TokenType = 4
	TokenTypeCompute     TokenType = 5
	TokenTypeBracket     TokenType = 6
	TokenTypeAssign      TokenType = 7
	TokenTypeSymbol      TokenType = 8
)

var errAsicc = fmt.Errorf("ascii only")

//func (tt *TokenType) String() string{
//	return
//}
