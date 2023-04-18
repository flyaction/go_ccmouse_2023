package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com
email2 is abc@qq.com
email3 is ccc@126.com
end
`

func main() {
	re := regexp.MustCompile("[a-z0-9A-Z]+@[a-z0-9A-Z.]+\\.[a-z0-9A-Z]+")

	math := re.FindAllString(text, -1)

	fmt.Println(re, "\n", math)
}
