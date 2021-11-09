package entity

//Student 用Tag进行字段信息描述
type Student struct {
	TableName string "t_student"
	Id        string `field:"id" iskey:"1"`
	Name      string `field:"cname"`
	Age       int    `field:"age"`
	Sex       string
}
