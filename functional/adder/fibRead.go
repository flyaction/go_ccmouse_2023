package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci2() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

type intGent struct {
	gen           func() int
	currentReader io.Reader
}

func (g *intGent) Read(p []byte) (n int, err error) {

	err = io.EOF
	if g.currentReader != nil {
		n, err = g.currentReader.Read(p)
	}
	if err == io.EOF {
		next := g.gen()
		s := fmt.Sprintf("%d\n", next)
		g.currentReader = strings.NewReader(s)
		if n == 0 {
			n, err = g.currentReader.Read(p)
		}
	}

	return n, err
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	f := intGent{
		gen: fibonacci2(),
	}

	//printFileContents(f)
	for i := 0; i < 20; i++ {
		b := make([]byte, 1)
		n, err := f.Read(b)
		fmt.Printf("%d bytes read: %q. err = %v\n", n, b, err)
	}

}
