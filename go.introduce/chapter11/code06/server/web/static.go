package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

//RegisterRouter 路由处理程序注册
func RegisterRouter() {
	//静态资源路径解析
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
	http.Handle("/font/", http.StripPrefix("/font/", http.FileServer(http.Dir("font/"))))
	http.HandleFunc("/api", handleController)
	http.HandleFunc("/", handleIndex)
}
func handleIndex(wr http.ResponseWriter, req *http.Request) {
	//移除/index.html前面的/
	fileName := strings.TrimPrefix(req.URL.String(), "/")
	//fmt.Println(fileName)
	//index.gohtml -> index.html
	if strings.HasSuffix(fileName, ".gohtml") {
		fileName = strings.TrimSuffix(fileName, ".gohtml")
		fileName = fileName + ".html"
		fmt.Println(fileName)
		//读取html文件内容
		t, err := template.ParseFiles(fileName)
		if err == nil {
			//返回客户端
			t.Execute(wr, nil)
		} else {
			//文件不存在
			t, _ := template.ParseFiles("error.html")
			t.Execute(wr, nil)
		}

	} else {
		//返回错误页面
		t, _ := template.ParseFiles("error.html")
		t.Execute(wr, nil)
	}

}
