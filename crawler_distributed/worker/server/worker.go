package main

import (
	"flag"
	"fmt"
	"log"

	"imooc.com/ccmouse/learngo/crawler_distributed/worker"

	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))

}
