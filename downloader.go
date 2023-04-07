package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/infra"
)

func getRetriever() retriever {

	//return testing.Retriever{}
	return infra.Retriever{}
}

// ?: Something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {

	retriever := getRetriever()
	fmt.Println(retriever.Get("http://www.imooc.com"))
}
