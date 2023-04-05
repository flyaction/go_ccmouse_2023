package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/queue"
)

func main() {
	var q queue.Queue
	//q := queue.Queue{1}

	isEmpty := q.IsEmpty()

	fmt.Println(isEmpty)

	q.Push(222)

	fmt.Println(q)

	q.Pop()

	fmt.Println(q)

}
