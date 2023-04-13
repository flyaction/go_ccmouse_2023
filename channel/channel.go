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

func creatWorker(id int) chan<- int {

	c := make(chan int)
	go worker(id, c)
	return c

}
func chanDemo() {
	var channel [10]chan<- int
	for i := 0; i < 10; i++ {
		channel[i] = creatWorker(i)

	}

	for i := 0; i < 10; i++ {
		channel[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channel[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)

}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	bufferedChannel()
}
