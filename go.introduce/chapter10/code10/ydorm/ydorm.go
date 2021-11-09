package ydorm

import (
	"fmt"
)

//Save 新增实体对象
func Save(entity interface{}) (isOk bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOk = false
			fmt.Println("捕获异常:", err)
		}
	}()
	strSQL, p := genInsertSQL(entity)
	fmt.Println(strSQL)
	fmt.Println(p)
	//执行SQL
	isOk = true
	return
}

//Update 更新实体对象
func Update(entity interface{}) (isOk bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOk = false
			fmt.Println("捕获异常:", err)
		}
	}()
	strSQL, p := genUpdateSQL(entity)
	fmt.Println(strSQL)
	fmt.Println(p)
	//执行SQL
	isOk = true
	panic("模拟更新失败")
	return
}

//Delete 删除实体对象
func Delete(entity interface{}) (isOk bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOk = false
			fmt.Println("捕获异常:", err)
		}
	}()
	strSQL, p := genDeleteSQL(entity)
	fmt.Println(strSQL)
	fmt.Println(p)
	//执行SQL
	isOk = true
	return
}
