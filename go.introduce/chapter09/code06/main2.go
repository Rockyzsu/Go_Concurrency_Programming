package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet(url string, ch chan string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bs, _ := ioutil.ReadAll(res.Body)
	ch <- fmt.Sprintf("[%s]:%d", url, len(bs))
}
func main() {
	urls := []string{
		"http://www.jd.com",
		"http://www.baidu.com",
	}
	bch := make(chan string, len(urls))
	for _, url := range urls {
		go httpGet(url, bch)
	}
	for range urls {
		msg := <-bch
		fmt.Println(msg)
	}
}
