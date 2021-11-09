package main

import (
	"net/http"
	_ "net/http/pprof" //开启pprof
	"time"
)

func getMsg() string {
	time.Sleep(1 * time.Second)
	return "Hello"

}

//http://127.0.0.1:8080/debug/pprof/
func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	msg := getMsg()
	w.Write([]byte(msg))
}
