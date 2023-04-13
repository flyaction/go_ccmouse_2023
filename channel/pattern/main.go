package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s : message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				return
			}
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

func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {

	done := make(chan struct{})
	m1 := msgGen("service1", done)
	//m2 := msgGen("service2")

	//m := fanIn(m1, m2)
	//m := fanInBySelect(m1, m2)

	for i := 0; i < 5; i++ {
		fmt.Println(<-m1)
		//if m, ok := nonBlockingWait(m2); ok {
		if m, ok := timeoutWait(m1, 2*time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
			//fmt.Println("no message from service2")
		}
	}

	done <- struct{}{}
	time.Sleep(time.Second)
}
