package main

import (
	"bufio"
	"fmt"
	"os"

	"imooc.com/ccmouse/learngo/functional/fib"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {

	tryDefer()

	writeFile("def.txt")

}
