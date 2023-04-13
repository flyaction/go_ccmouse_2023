package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {

	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}

}
func chanDemo() {
	var channel [10]chan int
	for i := 0; i < 10; i++ {
		channel[i] = make(chan int)
		go worker(i, channel[i])
	}

	for i := 0; i < 10; i++ {
		channel[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channel[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}
