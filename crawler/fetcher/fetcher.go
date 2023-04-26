package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"imooc.com/ccmouse/learngo/crawler/config"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {

	<-rateLimiter
	log.Printf("Fetching url %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code :%d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
