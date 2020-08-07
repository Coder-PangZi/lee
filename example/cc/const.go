package main

const (
	TOKEN_TYPE_BAD TokenType = 1
	TOKEN_TYPE_NUM TokenType = 2
	TOKEN_TYPE_ADD TokenType = 3
	TOKEN_TYPE_SUB TokenType = 4
	TOKEN_TYPE_MUL TokenType = 5
	TOKEN_TYPE_DIV TokenType = 6
	TOKEN_TYPE_EOF TokenType = 7

	TOKEN_MAX_SIZE = 100
)

const (
	LEXER_STATUS_INIT LexerStatus = 1
	LEXER_STATUS_INT  LexerStatus = 1
	LEXER_STATUS_DOT  LexerStatus = 1
	LEXER_STATUS_FRAC LexerStatus = 1
)
