// token 包内的结构以用于记录词法分析结果
// 及其相关附加信息，比如如位置信息

package token

import "fmt"

// Token 表示词法分析的结果
type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMENT

	LiteralBeg
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	CHAR   // 'a'
	STRING // "abc"
	LiteralEnd

	OperatorBeg
	ADD        // +
	SUB        // -
	MUL        // *
	QUO        // /
	REM        // %
	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	LAND // &&
	LOR  // ||
	INC  // ++
	DEC  // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ // !=
	LEQ // <=
	GEQ // >=

	LPAREN    // (
	LBRACK    // [
	LBRACE    // {
	COMMA     // ,
	PERIOD    // .
	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	OperatorEnd

	KeywordBeg
	BREAK
	CONTINUE
	ELSE
	FOR

	FUNC
	IF

	RETURN
	VAR
	KeywordEnd
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	LAND: "&&",
	LOR:  "||",
	INC:  "++",
	DEC:  "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ: "!=",
	LEQ: "<=",
	GEQ: ">=",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	BREAK:    "break",
	CONTINUE: "continue",

	ELSE: "else",
	FOR:  "for",

	FUNC: "func",
	IF:   "if",

	RETURN: "return",
	VAR:    "var",
}

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = fmt.Sprintf("token(%d)", tok)
	}
	return s
}

type KeyWords map[string]Token

var keywords KeyWords

func (kw KeyWords) LookUp(ident string) (Token, bool) {
	tok, is := kw[ident]
	return tok, is
}

func init() {
	keywords = make(KeyWords)
	for i := KeywordBeg + 1; i < KeywordEnd; i++ {
		keywords[i.String()] = i
	}
}

func Type(ident string) Token {
	if tok, is := keywords.LookUp(ident); is {
		return tok
	}
	return IDENT
}

// IsLiteral 返回 token 类型是否为字面量类型。
// @return bool
//
func (tok Token) IsLiteral() bool { return LiteralBeg < tok && tok < LiteralEnd }

// IsOperator 返回 token 类型是否为操作符。
//
func (tok Token) IsOperator() bool { return OperatorBeg < tok && tok < OperatorEnd }

// IsKeyword 返回 token 类型是否为关键字
//
func (tok Token) IsKeyword() bool { return KeywordBeg < tok && tok < KeywordEnd }
