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

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {

	return r.Get("http://xingdong365.com")
}

func post(poster Poster) {
	poster.Post("http://xingdong365.com", map[string]string{
		"name": "action",
		"age":  "18",
	})
}

func session(s RetrieverPoster) string {
	//poster.Get("http://xingdong365.com")
	s.Post("http://xingdong365.com", map[string]string{
		"contents": "ddddddddddd",
	})
	return s.Get("http://xingdong365.com")
}

func main() {

	var r Retriever

	r_mock := mock.Retriever{"this is a fake imooc.com"}
	inspect(&r_mock)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))
	//type Assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println(session(&r_mock))

}

// type switch
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch t := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", t.Contents)
	case *real.Retriever:
		fmt.Println("Useragent:", t.UserAgent, "Timeout:", t.Timeout)

	}
}
