package parallel

import (
	"fmt"
	"sync"
	"time"
)

//Gopher 地鼠
type Gopher struct {
	Name       string
	Id         int
	CoffeeName string
}

//两台咖啡机
var coffeeMachineArr = [2]CoffeeMachine{}

func init() {
	var coffeeMachine = CoffeeMachine{Name: "CoffeeMachine1"}
	var coffeeMachine2 = CoffeeMachine{Name: "CoffeeMachine2"}
	coffeeMachineArr[0] = coffeeMachine
	coffeeMachineArr[1] = coffeeMachine2
}

//MakeCoffee 制作咖啡
func (this *Gopher) MakeCoffee(coffeeName string, wg *sync.WaitGroup) {
	//独有的咖啡机
	coffeeMachine := coffeeMachineArr[this.Id%2]
	//coffeeMachine加锁
	coffeeMachine.Mlock.Lock()
	//(1)研磨咖啡(没有才研磨)
	if coffeeMachine.CoffeeName == "" {
		coffeeMachine.CoffeeName = coffeeName
		coffeeMachine.Gopher = *this
		fmt.Println("Gopher", this.Id, "Make Coffee", coffeeMachine.CoffeeName)
		time.Sleep(10 * time.Second)
	}
	//(2)倒咖啡
	this.TakeCoffee(coffeeMachine)
	//(3)喝咖啡
	this.DrinkCoffee(coffeeMachine)
	//释放coffeeMachine锁
	coffeeMachine.Mlock.Unlock()
	wg.Done()
}
func (this *Gopher) TakeCoffee(coffeeMachine CoffeeMachine) {

	if coffeeMachine.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Take Coffee", coffeeMachine.CoffeeName)
		this.CoffeeName = coffeeMachine.CoffeeName
		time.Sleep(5 * time.Second)
		//倒完咖啡
		coffeeMachine.CoffeeName = ""
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee to Take")
		this.CoffeeName = ""
	}

}
func (this *Gopher) DrinkCoffee(coffeeMachine CoffeeMachine) {

	if this.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Drink Coffee", this.CoffeeName)
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee to Drink")
	}

}
