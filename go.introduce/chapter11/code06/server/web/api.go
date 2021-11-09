package web

import (
	"encoding/json"
	"fmt"
)

//TOKEN 登录后获取真的token
const TOKEN = "REAL_JASDNFANFASLJI123LMFYd88"

//Login 登录验证
func Login(r *ReqData) ResData {
	res := ResData{}
	vjson := make(map[string]interface{})
	err := json.Unmarshal([]byte(r.data), &vjson)
	if err == nil {
		fmt.Println(vjson)
		//类型断言
		uname, upwd := "", ""
		if value, ok := vjson["uname"].(string); ok {
			uname = value
		}
		if value, ok := vjson["pwd"].(string); ok {
			upwd = value
		}

		if uname == "root" && upwd == "123" {
			res.Code = 1
			res.Msg = "ok"
			res.Data = TOKEN
		} else {
			res.Code = 0
			res.Msg = "fail"
			res.Data = ""
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println(res)
	return res
}

//UnLogin 未登录返回的消息
func UnLogin() ResData {
	res := ResData{}
	res.Code = 0
	res.Msg = "not login"
	res.Data = ""
	return res
}
