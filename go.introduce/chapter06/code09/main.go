package main

import (
	"fmt"
	"io"

	"go.introduce/chapter06/code09/linklist"
)

var list *linklist.LinkList

func init() {
	list = linklist.CreateLinkList()
}
func main() {
	fmt.Print("$>")
	cmd := ""
	var data int = 0
	for {
		n, err := fmt.Scanln(&cmd, &data)
		if err == io.EOF {
			break
		}
		if n > 0 {
			if cmd == "init" {

				for i := 0; i < data; i++ {
					list.Insert(uint(i), linklist.CreateLinkNode(i+1))
				}
				list.PrintLink()
				fmt.Print("$>")
			} else if cmd == "add" {
				node := linklist.CreateLinkNode(data)
				list.Insert(0, node)
				list.PrintLink()
				fmt.Print("$>")
			} else if cmd == "remove" {
				//删除
				list.Delete(linklist.CreateLinkNode(data))
				list.PrintLink()

				fmt.Print("$>")
			} else if cmd == "get" {
				node, err := list.Find(uint(data))
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("node %d value:%v\n", data, node.GetValue())
				}

				fmt.Print("$>")
			} else if cmd == "exit" {
				break
			} else {
				fmt.Println("不支持的语法,支持init;add;remove;get;exit")
				fmt.Print("$>")
			}

		}

	}

}
