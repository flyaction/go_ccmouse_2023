package client

import (
	"log"

	"imooc.com/ccmouse/learngo/crawler_distributed/config"

	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"

	"imooc.com/ccmouse/learngo/crawler/engine"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)

			if err != nil {
				log.Printf("Item Saver :error"+"saving item %v %v", item, err)
				continue
			}
		}
	}()

	return out, nil

}
