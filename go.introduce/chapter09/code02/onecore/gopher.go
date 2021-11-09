package onecore

import (
	"fmt"
	"time"
)

//Gopher 地鼠
type Gopher struct {
	Name       string
	Id         int
	CoffeeName string
}

var coffeeMachine = CoffeeMachine{}

//MakeCoffee 制作咖啡
func (this *Gopher) MakeCoffee(coffeeName string) {
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
	this.TakeCoffee()
	//(3)喝咖啡
	this.DrinkCoffee()
	//释放coffeeMachine锁
	coffeeMachine.Mlock.Unlock()

}
func (this *Gopher) TakeCoffee() {

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
func (this *Gopher) DrinkCoffee() {

	if this.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Drink Coffee", this.CoffeeName)
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee to Drink")
	}

}
