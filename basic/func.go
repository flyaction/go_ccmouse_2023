package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval2(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "_":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("undefined :" + op)
	}
}

func div(a, b int) (p, q int) {
	p = a / b
	q = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {

	p := reflect.ValueOf(op).Pointer()
	pName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d)\n", pName, a, b)
	return op(a, b)
}

func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {

	m, n := div(8, 3)

	fmt.Println(m, n)

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sumArgs(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	c, d := 7, 8
	c, d = swap2(c, d)
	fmt.Println(c, d)

}
