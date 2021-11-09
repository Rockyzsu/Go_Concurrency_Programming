package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Msg string //首字母大写
}

// 返回静态html
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	data := Data{Msg: "Hello HTTP"}
	t.Execute(writer, data) //赋值
}
func main() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
