package main

import (
	"fmt"
)

func change(m map[string]string) {
	m["name"] = "Jack"
}
func main() {

	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	change(myMap)
	//map[key1:value1 key2:value2 name:Jack]
	fmt.Println(myMap)
	myMap2 := myMap
	delete(myMap2, "key1")
	//map[key2:value2 name:Jack]
	fmt.Println(myMap)
	//map[key2:value2 name:Jack]
	fmt.Println(myMap2)
	v, e := myMap2["name"] //判断键name是否存在
	if e {
		fmt.Println(v) //Jack
	}
}
