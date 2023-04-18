package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code", resp.StatusCode)
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	printCityList(all)

}

func printCityList(contents []byte) {

	re := regexp.MustCompile(`<a href=\"(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {

		for _, subMatch := range m {
			fmt.Printf("%s ", subMatch)
		}
		fmt.Println()

	}

}
