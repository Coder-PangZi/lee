package types

import (
	"sync"
	"sync/atomic"
)

type TokenQueue struct {
	tokens []*Token
	tail   int32
	cur    int32
	mux    sync.Mutex
}

func NewTokenQueue(len int) *TokenQueue {
	return &TokenQueue{tokens: make([]*Token, 0, len), tail: 0, cur: 0}
}

func (tq *TokenQueue) Push(t *Token) {
	tq.mux.Lock()
	defer tq.mux.Unlock()
	tq.tokens[tq.tail] = t
	atomic.AddInt32(&tq.tail, 1)
}

func (tq *TokenQueue) Pop() *Token {
	tq.mux.Lock()
	defer tq.mux.Unlock()
	t := tq.tokens[tq.tail]
	tq.tokens[tq.tail] = nil
	atomic.AddInt32(&tq.tail, -1)
	return t
}

func (tq *TokenQueue) Peak(i int32) *Token {
	tq.mux.Lock()
	defer tq.mux.Unlock()
	t := *tq.tokens[tq.tail+i]
	return &t
}
