package entity

//IDrive 接口
//只要实现了Drive(name string)方法即实现了IDrive 接口
type IDrive interface {
	Drive(name string)
}
