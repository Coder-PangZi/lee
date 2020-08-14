package lexer

import (
	"github.com/davecgh/go-spew/spew"
	"strings"
	"testing"
)

func TestInt2Str(t *testing.T) {
	num := -10086
	t.Log(int2str(num))
}

var set = "0123456789"

func add() {
	for i := 0; i < 10000; i++ {
		_ = set[0]
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add()
	}
}

func TestFile(t *testing.T) {
	const SC = `abcdef asdw adw`
	f, err := NewFile(strings.NewReader(SC))
	if err != nil {
		t.Error(err)
	}
	err = f.next()
	if err != nil {
		t.Error(err)
	}
	spew.Dump(f.chr)
}
