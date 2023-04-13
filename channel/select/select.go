package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(5 * time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func creatWorker(id int) chan<- int {

	c := make(chan int)
	go worker(id, c)
	return c

}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func main() {

	var c1, c2 = generator(), generator()
	var worker = creatWorker(0)

	//haseValue := false
	var values []int

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		}
	}

}
