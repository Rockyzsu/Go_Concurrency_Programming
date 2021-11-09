package web

//ResData 返回数据格式
//属性必须大写，否则JSON序列化忽略
type ResData struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data string        `json:"data"`
	Tag  []interface{} `json:"tag"`
}
