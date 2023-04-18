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
	re := regexp.MustCompile("([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\\.([a-zA-Z0-9]+)")

	//math := re.FindAllString(text, -1)
	//fmt.Println(re, "\n", math)

	math := re.FindAllStringSubmatch(text, -1)
	for _, m := range math {
		fmt.Println(m)
	}

}
