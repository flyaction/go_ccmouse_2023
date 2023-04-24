package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {

	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	printFileContents(file)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
	)

	printFile("basic/loop/abc.txt")

	fmt.Println("=========================")

	s := `abc
kkk
fff
`
	printFileContents(strings.NewReader(s))

}
