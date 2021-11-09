package main

import (
	"fmt"
)

func main() {
	var myMap map[string]string
	//myMap["key1"] = "value1" //错误nil map
	fmt.Println(myMap) //map[]

	myMap = make(map[string]string)
	//新增
	myMap["key1"] = "value1"
	myMap["key2"] = "value2"
	//map[key1:value1 key2:value2]
	fmt.Println(myMap)
	//修改
	myMap["key1"] = "value2"
	//删除
	delete(myMap, "key2")
	fmt.Println(myMap) //map[key1:value2]
	//查
	fmt.Println(myMap["key1"]) //value2
	/////////////////
	//字面量创建
	myMap2 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	myMap2["key3"] = "value3"
	//map[key1:value1 key2:value2 key3:value3]
	fmt.Println(myMap2)
}
