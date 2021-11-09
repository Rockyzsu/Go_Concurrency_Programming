package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8080?key=Go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://127.0.0.1:8080",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Go"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func main() {
	httpGet()
	httpPost()
}
