package main

import (
	"fmt"
	"testing"
	"time"

	"imooc.com/ccmouse/learngo/crawler/config"
	rpcnames "imooc.com/ccmouse/learngo/crawler_distributed/config"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
	"imooc.com/ccmouse/learngo/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(
		host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Error(err)
	}

	// TODO: Use a fake fetcher to handle the url.
	// So we don't get data from zhenai.com
	req := worker.Request{
		Url: "http://www.7799520.com/user/6732595.html",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "哭泣de科拉",
		},
	}
	var result worker.ParseResult
	err = client.Call(
		rpcnames.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

	// TODO: Verify results
}
