package main

func main() {
	ch := make(chan int, 1)

	ch <- 1
	close(ch)
	dat1, ok1 := <-ch
	dat2, ok2 := <-ch
	println(dat1, ok1)
	println(dat2, ok2)
}
