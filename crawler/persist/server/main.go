package main

import (
	"github.com/olivere/elastic/v7"
	"imooc.com/ccmouse/learngo/crawler_distributed/persist"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
)

func main() {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	rpcsupport.ServeRpc(":1234", persist.ItemSaverService{
		Client: client,
		Index:  "dating_profile",
	})
}
