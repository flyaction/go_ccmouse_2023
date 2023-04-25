package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/crawler/jiaou/parserjiaou"
	"imooc.com/ccmouse/learngo/crawler_distributed/config"
	"imooc.com/ccmouse/learngo/crawler_distributed/persist/client"

	"imooc.com/ccmouse/learngo/crawler/scheduler"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

func main() {

	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
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
		Parser: engine.NewFuncParser(parserjiaou.ParseCityList, config.ParseCityList),
	})

}
