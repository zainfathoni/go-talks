package main

import "time"

// START Concurrency OMIT

func HeavyComputation(ch chan int32) {
	// long, serious math stuff
	time.Sleep(2 * time.Second)

	ch <- 12345
}

func main() {
	ch := make(chan int32)

	println("Start computing!")
	go HeavyComputation(ch)

	println("Meanwhile, do other stuffs...")

	result := <-ch
	println(result)
}

// END Concurrency OMIT
