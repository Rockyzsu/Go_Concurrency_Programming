package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 返回动态内容
//http://localhost:8080/api
func handleController(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("======api==========")
	if req.Method == "POST" {
		req.ParseForm()
		api := req.Form.Get("api")
		fmt.Println(api)
		data := req.Form.Get("data")
		fmt.Println(data)
		token := req.Form.Get("token")
		if api == "login" {
			reqdata := &ReqData{api: api, data: data, token: token}
			res := Login(reqdata)
			//[]byte
			strRes, err := json.Marshal(&res)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(strRes))
			wr.Write(strRes)
		} else {

			if token != TOKEN {
				//未登录
				res := UnLogin()
				strRes, _ := json.Marshal(&res)
				wr.Write(strRes)
			}
		}
	} else {
		wr.Write([]byte("only support post"))
	}

}
