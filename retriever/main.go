package main

import (
	"fmt"
	"time"

	"imooc.com/ccmouse/learngo/retriever/mock"

	"imooc.com/ccmouse/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {

	return r.Get("http://xingdong365.com")
}

func main() {

	var r Retriever

	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))
	//type Assertion
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

}

// type switch
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch t := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", t.Contents)
	case *real.Retriever:
		fmt.Println("Useragent:", t.UserAgent, "Timeout:", t.Timeout)

	}
}
