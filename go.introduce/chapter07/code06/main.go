package main

import (
	"fmt"

	"go.introduce/chapter07/code06/entity"
)

func main() {
	animal := entity.Factory("duck")
	animal.SingGua()
	fmt.Printf("animal is %s ", animal.Type())
	animal = entity.Factory("goose")
	fmt.Println()
	animal.SingGua()
	fmt.Printf("animal is %s ", animal.Type())
}
