package main

import (
	"fmt"
	"log"

	"imooc.com/ccmouse/learngo/crawler_distributed/config"

	"github.com/olivere/elastic/v7"
	"imooc.com/ccmouse/learngo/crawler_distributed/persist"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
)

func main() {

	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), fmt.Sprintf(":%s", config.ElasticIndex)))
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
