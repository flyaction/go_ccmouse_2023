package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {

	lastOccurred := make(map[rune]int)

	start, maxLength := 0, 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength

}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcdabcde"))
	fmt.Println(lengthOfNonRepeatingSubStr("我爱我爱慕课"))
}
