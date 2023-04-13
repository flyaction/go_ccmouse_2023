package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s : message %d", name, i)
			i++
		}
	}()

	return c
}

func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func main() {
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	m := fanIn(m1, m2)

	for {
		fmt.Println(<-m)
	}
}
