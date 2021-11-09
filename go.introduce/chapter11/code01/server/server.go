package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			vars := r.URL.Query()
			key, ok := vars["key"]
			if ok {
				msg := "hello get " + key[0]
				w.Write([]byte(msg))
			} else {
				w.Write([]byte("hello world!"))
			}
		}
		if r.Method == "POST" {
			r.ParseForm()
			key := r.Form.Get("name")
			msg := "hello post " + key
			w.Write([]byte(msg))
		}

	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
