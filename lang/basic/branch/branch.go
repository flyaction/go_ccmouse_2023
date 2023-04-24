package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	const fileName = "abc.txt"
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(50),
		grade(60),
		grade(70),
		grade(80),
		grade(90),
	)

}

func eval(a int, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "_":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("undefined :" + op)
	}

	return result
}

func grade(score int) string {
	var result string
	switch {
	case score < 0 || score > 100:
		panic("错误")
	case score < 60:
		result = "F"
	case score < 70:
		result = "D"
	case score < 80:
		result = "C"
	case score < 90:
		result = "B"
	case score <= 100:
		result = "A"
	default:
		result = "X"
	}
	return result
}
