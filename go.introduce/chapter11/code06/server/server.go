package main

import (
	"fmt"
	"net/http"

	"go.introduce/chapter11/code06/server/web"
)

//http://localhost:8080/index.gohtml
func main() {
	web.RegisterRouter()
	fmt.Println("start Go Web Server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
