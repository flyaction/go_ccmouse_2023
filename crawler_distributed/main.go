package main

import (
	"fmt"
	"log"

	worker "imooc.com/ccmouse/learngo/crawler_distributed/worker/client"

	"imooc.com/ccmouse/learngo/crawler/jiaou/parserjiaou"
	"imooc.com/ccmouse/learngo/crawler_distributed/config"
	itemsaver "imooc.com/ccmouse/learngo/crawler_distributed/persist/client"

	"imooc.com/ccmouse/learngo/crawler/scheduler"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

func main() {

	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
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
