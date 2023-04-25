package main

import (
	"github.com/olivere/elastic/v7"
	"imooc.com/ccmouse/learngo/crawler_distributed/persist"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
)

func main() {

	serveRpc(":1234", "dating_profile")
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
