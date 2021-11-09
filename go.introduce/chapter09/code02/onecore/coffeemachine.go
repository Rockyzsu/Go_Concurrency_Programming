package onecore

import "sync"

//CoffeeMachine 咖啡机（共享1个）
type CoffeeMachine struct {
	CoffeeName string
	Gopher                //获取使用权的地鼠
	Mlock      sync.Mutex //互斥锁
}
