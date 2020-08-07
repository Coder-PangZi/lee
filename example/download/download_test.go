package main

import (
	"github.com/davecgh/go-spew/spew"
	"os"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {
	println(1)
	spew.Dump(filepath.Abs("./"))
	println(2)
	spew.Dump(os.Stat("./aaa"))
}
