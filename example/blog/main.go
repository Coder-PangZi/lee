package main

import (
	"fmt"
	"time"
)

func main() {
	sum()
}

func sum() {
	defer stat()()
	time.Sleep(time.Second / 2)
}

func stat() func() {
	st := time.Now()
	return func() {
		dt := time.Since(st).Nanoseconds()
		sec := dt / int64(time.Second)
		dt = dt % int64(time.Second)
		mill := dt / int64(time.Millisecond)
		dt = dt % int64(time.Millisecond)
		micro := dt / int64(time.Microsecond)
		nano := dt % int64(time.Microsecond)
		fmt.Printf("%02d:%02d:%02d:%02d", sec, mill, micro, nano)
	}
}
