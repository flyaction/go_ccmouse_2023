package main

import (
	"bufio"
	"errors"
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
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 06666)

	err = errors.New("this is a custom error")

	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		//fmt.Println("Error:", err)
		return
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
