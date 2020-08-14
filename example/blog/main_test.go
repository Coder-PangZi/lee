package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func BenchmarkBS(b *testing.B) {

	users := map[string]struct {
		Name string
		age  int
	}{"lee": {Name: "lee", age: 18}}

	name := []byte{'l', 'e', 'e'}
	//_, _, _ = bufio.Reader.ReadLine()
	for i := 0; i < b.N; i++ {
		n := string(name)
		u := users[n]

		//n := ""
		//u := users[string(name)]

		_, _ = u, n
	}
}

func TestTypeCharge(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var newS []int64
	// 做法是利用 reflect 直接替换数据指针
	// 但是这个不保证在以后的版本中一直可用 ╮(╯▽╰)╭
	*(*reflect.SliceHeader)(unsafe.Pointer(&newS)) = *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("type:%T value:%v", newS, newS)
}

func AlgoFisherYates(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		k := rand.Intn(i + 1)
		print(k, "\t")
		s[i], s[k] = s[k], s[i]
	}
}

func BenchmarkFY(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AlgoFisherYates(s)
	}
}

func TestMMap(t *testing.T) {

}
