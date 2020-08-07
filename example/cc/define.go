package main

type TokenType byte
type LexerStatus byte

type Token struct {
	Kind TokenType
	val  float64
	buf  [TOKEN_MAX_SIZE]byte
}

type Lexer struct {
	pos    int
	status LexerStatus
	token  *Token
}

func (l *Lexer) GetToken() *Token {

}
