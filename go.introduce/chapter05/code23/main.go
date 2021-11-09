package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"go.introduce/chapter05/code23/ydbase"
)

func main() {

	fmt.Print("#>")
	uname := ""
	upwd := ""
	for {
		n, err := fmt.Scanln(&uname, &upwd)
		if err == io.EOF {
			break
		}
		if n > 0 {
			if !ydbase.Login(uname, upwd) {
				fmt.Println("登录失败")
				fmt.Print("#>")
			} else {
				fmt.Println("欢迎使用YdBase 1.0")
				fmt.Print("$>")
				break
			}

		}

	}
	cmd := ""
	exp := ""
	for {
		n, err := fmt.Scanln(&cmd, &exp)
		//fmt.Println(cmd, exp)
		if err == io.EOF {
			break
		}
		if n > 0 {
			if cmd == "init" {
				len1, _ := strconv.Atoi(exp)
				for i := 1; i <= len1; i++ {
					stu := ydbase.Student{
						XH:     "09" + strconv.Itoa(i),
						Name:   "Name" + strconv.Itoa(i),
						Age:    22 + i,
						Height: 1.7,
					}
					//新增
					ydbase.Insert(stu)

				}
				fmt.Printf("操作成功,共%d行\n", ydbase.Count())
				fmt.Print("$>")
			} else if cmd == "insert" {
				fields := strings.Split(exp, ",")
				leng := len(fields)
				if leng != 4 {
					fmt.Println("insert需要4个参数")
					fmt.Print("$>")
				}
				age, _ := strconv.Atoi(fields[2])
				//string to float64
				height64, _ := strconv.ParseFloat(fields[3], 32)
				//float64 to float32
				height := float32(height64)
				stu := ydbase.Student{
					XH:     fields[0],
					Name:   fields[1],
					Age:    age,
					Height: height,
				}
				//新增
				ydbase.Insert(stu)

				fmt.Printf("操作成功,共%d行\n", ydbase.Count())
				fmt.Print("$>")
			} else if cmd == "delete" {
				//删除
				ydbase.Delete(exp)
				fmt.Printf("操作成功,共%d行\n", ydbase.Count())
				fmt.Print("$>")
			} else if cmd == "select" {
				//"age<=24"
				exp := strings.Replace(exp, "%", " like ", 1)
				stu := ydbase.Query(exp)
				fmt.Printf("操作成功,匹配%d行\n", len(stu))
				if len(stu) > 0 {
					//fmt.Println(stu)
					ydbase.Print(stu)
				}

				fmt.Print("$>")
			} else if cmd == "update" {
				fields := strings.Split(exp, ",")
				leng := len(fields)
				if leng != 4 {
					fmt.Println("insert需要4个参数")
					fmt.Print("$>")
				}
				stu := ydbase.Query("xh==" + fields[0])
				if len(stu) == 1 {
					stu[0].Name = fields[1]
					age, _ := strconv.Atoi(fields[2])
					//string to float64
					height64, _ := strconv.ParseFloat(fields[3], 32)
					//float64 to float32
					height := float32(height64)
					stu[0].Age = age
					stu[0].Height = height
					//修改
					ydbase.Modify(stu[0])
					fmt.Printf("操作成功,影响1行\n")
				}

				fmt.Print("$>")
			} else if cmd == "count" {
				fmt.Printf("操作成功,共%d行\n", ydbase.Count())
				fmt.Print("$>")
			} else if cmd == "exit" {
				break
			} else {
				fmt.Println("不支持的语法,支持init;delete;update;select")
				fmt.Print("$>")
			}

		}

	}
	fmt.Print("#>")
	// root 123
	//init 10
	//select age<=25
	//count
	//delete 091
	//insert 001,jack,33,1.81
	//select name==jack
	//update 001,jack2,32,1.79
	//select name%Nam -> name like '%Nam%'
}
