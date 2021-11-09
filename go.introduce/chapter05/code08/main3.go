package main

type BeforeSave func()

//Student 学生结构
type Student struct {
	XH     string
	Name   string
	Age    int
	Height float32
	Class  string
	//匿名字段(函数类型)
	BeforeSave
}

func (s Student) Save() int {
	if s.BeforeSave != nil {
		s.BeforeSave()
	}
	println("=======save========", s.Name)
	return 1
}
func main() {

	var stu = Student{}
	stu.BeforeSave = func() {
		println("=======before save========")

	}
	println(stu.Save())
}
