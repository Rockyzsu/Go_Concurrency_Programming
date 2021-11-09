package ydbase

//Student 学生结构
type Student struct {
	XH, Name string
	Age      int
	Height   float32
}

//学生表
var studentTable = []Student{}

func init() {
	//初始化表，容量100
	studentTable = make([]Student, 0, 100)
}
