package main

import "go.introduce/chapter07/code03/entity"

func main() {
	name := "A6L"
	aoDi := entity.AoDi{}
	aoDi.Drive(name) //Drive AoDi A6L Car
	//值接收者实现的接口可以传递实体类型的地址或者实体类型
	mycar := entity.MyCar{&aoDi} //ok
	//mycar := entity.MyCar{aoDi}
	mycar.Drive(name) //Drive AoDi A6L Car
	name = "X6"
	//指针接收者实现的接口必须传递实体类型的地址
	mycar = entity.MyCar{&entity.BMW{}}
	mycar.Drive(name) //Drive BMW X6 Car

	mycar.IDrive.Drive(name) //Drive BMW X6 Car

}
