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

func fanIn(chs ...chan string) chan string {
	c := make(chan string)

	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in // ch == in
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case n := <-c1:
				c <- n
			case n := <-c2:
				c <- n
			}
		}
	}()

	return c
}

func main() {
	m1 := msgGen("service1")
	m2 := msgGen("service2")

	m := fanIn(m1, m2)
	//m := fanInBySelect(m1, m2)

	for {
		fmt.Println(<-m)
	}
}
