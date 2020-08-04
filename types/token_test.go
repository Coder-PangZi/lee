package types

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestLToken_Trim(t *testing.T) {
	tt := NewToken(1, []byte(" \t myName \n"))
	tt.trim()
	spew.Dump(tt)
}
