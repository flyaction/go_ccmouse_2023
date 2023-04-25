package main

import (
	"fmt"
	"regexp"

	"imooc.com/ccmouse/learngo/crawler/jiaou/parserjiaou"

	"imooc.com/ccmouse/learngo/crawler/persist"

	"imooc.com/ccmouse/learngo/crawler/scheduler"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

func main() {

	//engine.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "https://www.7799520.com/jiaou",
	//	ParserFunc: parserjiaou.ParseCityList,
	//})

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		//panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}
	//e.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:    "https://www.7799520.com/jiaou",
		Parser: engine.NewFuncParser(parserjiaou.ParseCityList, "ParseCityList"),
	})

	//resp, err := http.Get("https://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error:status code", resp.StatusCode)
	//}
	//
	//all, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Printf("%s\n", all)
	//printCityList(all)

}

func printCityList(contents []byte) {

	re := regexp.MustCompile(`<a href=\"(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		fmt.Printf("City:%s,URL:%s\n", m[2], m[1])
	}

}
