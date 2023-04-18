package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"

func main() {
	re := regexp.MustCompile("[a-z0-9A-Z]+@[a-z0-9A-Z]+\\.[a-z0-9A-Z]+")

	math := re.FindString(text)

	fmt.Println(re, "\n", math)
}
