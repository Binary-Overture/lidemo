package main

import (
	"firstproject/memo"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)
	return io.ReadAll(response.Body)
}

func incomingURLs() <-chan string {
	urls := make(chan string)

	// 在后台获取URL列表并将其发送到通道
	go func() {
		urls <- "https://www.baidu.com"
		//urls <- "https://www.ctfer.vip/index"
		urls <- "https://www.bilibili.com"
		// ...更多URL
		close(urls)
	}()

	return urls
}

func main() {
	m := memo.New(httpGetBody)
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			resp, err := m.Get(url)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(resp.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
