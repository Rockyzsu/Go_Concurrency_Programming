package parallel

import "sync"

//CoffeeMachine 咖啡机
type CoffeeMachine struct {
	Name       string
	CoffeeName string
	Gopher                //获取使用权的地鼠
	Mlock      sync.Mutex //互斥锁
}
