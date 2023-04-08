package main

import "fmt"

func adder() func(int) int { //返回的是一个函数
	sum := 0                 //sum自由变量
	return func(v int) int { //v局部变量
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		base += v
		return base, adder2(base)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+......%d = %d \n", i, a(i)) // a的函数体存的是 sum值
	}

	fmt.Println()

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		s, _ := a2(i)
		fmt.Printf("0+1+......%d = %d \n", i, s)
	}
}
