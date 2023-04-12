package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 33
//var bb = 44
//var ss = "def"

var (
	aa = 33
	bb = 44
	ss = "def2"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)

}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "action"
	fmt.Println(a, b, s)

}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "Singsong"
	b = 5
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "test"
	fmt.Println(a, b, c, s)
}

//欧拉公式
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) //小数点后三位

}

//强制类型转换
func triangle() {

	var a, b int = 3, 4
	fmt.Println(callTriangle(a, b))

}

func callTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const fileName = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b)) //a,b不定义类型，不用转
	fmt.Println(fileName, c)
}

func enums() {
	const (
		java = iota
		_
		php
		python
		c
		js
	)
	fmt.Println(java, php, python, c, js) // 0,2,3,4,5

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb) //1 1024 1048576 1073741824 1099511627776

}

func main() {
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)

	euler()

	triangle()

	consts()

	enums()
}
