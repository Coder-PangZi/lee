package main

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/esdb/drbuffer"
	"sync"
	"sync/atomic"
	"testing"
)

func add(is []int64) int64 {
	t := int64(0)
	for _, n := range is {
		t += n
	}
	return t
}

func addExt(is []int64, cnt int) int64 {
	length := int64(len(is))
	chunkSize := length / int64(cnt)
	t := int64(0)
	wg := sync.WaitGroup{}
	wg.Add(cnt)
	for i := int64(0); i < int64(cnt); i++ {
		go func(i int64) {
			start := i * chunkSize
			end := start + chunkSize
			if end > length {
				end = length
			}
			sum := add(is[start:end])
			atomic.AddInt64(&t, sum)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return t
}

func gen(s []int64) error {
	l := int64(len(s))
	for i := int64(0); i < l; i++ {
		s[i] = i + 1
	}
	return nil
}

func TestAdd(t *testing.T) {
	is := make([]int64, 1000000, 1000000)
	_ = gen(is)
	t.Log(add(is))
	is = make([]int64, 1000000, 1000000)
	_ = gen(is)
	t.Log(addExt(is, 1))
}

func BenchmarkAdd(b *testing.B) {
	is := make([]int64, 1000000, 1000000)
	_ = gen(is)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addExt(is, 3)
	}
}

func TestError(t *testing.T) {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%s]", err1)
	err3 := fmt.Errorf("err3: [%s]", err2)
	fmt.Println(err3)
}

func TestFW(t *testing.T) {
	is := make([]int64, 1000000, 1000000)
	_ = gen(is)
	spew.Dump(FuncAdd(add).WarpAdd(is))
}

type FuncAdd func([]int64) int64

func (fa FuncAdd) WarpAdd(is []int64) int64 {
	return fa(is)
}

func TestRing(t *testing.T) {
	ring, err := drbuffer.Open("./", 1)
	if err == nil {
		t.Error(err.Error())
	}
	ring.PopOne()
}
