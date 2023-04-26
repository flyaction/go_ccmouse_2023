package client

import (
	"net/rpc"

	"imooc.com/ccmouse/learngo/crawler/engine"
	"imooc.com/ccmouse/learngo/crawler_distributed/config"
	"imooc.com/ccmouse/learngo/crawler_distributed/worker"
)

//func CreateProcessor() (engine.Processor, error) {
//
//	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
//	if err != nil {
//		return nil, err
//	}
//
//	return func(req engine.Request) (engine.ParseResult, error) {
//
//		sReq := worker.SerializeRequest(req)
//
//		var sResult worker.ParseResult
//
//		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
//
//		if err != nil {
//			return engine.ParseResult{}, err
//		}
//		return worker.DeserializeResult(sResult), nil
//	}, nil
//}

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(req engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc,
			sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
