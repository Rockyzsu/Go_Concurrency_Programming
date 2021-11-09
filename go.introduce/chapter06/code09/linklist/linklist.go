//Go单链表实现
package linklist

import (
	"fmt"
)

type LinkList struct {
	head   *ListNode
	length uint
}

//CreateLinkList 初始化单链表
func CreateLinkList() *LinkList {
	return &LinkList{
		head:   &ListNode{0, nil},
		length: 0,
	}
}

//PrintLink 打印链表
func (this *LinkList) PrintLink() {
	cur := this.head.GetNext()
	for nil != cur {
		fmt.Printf("%v->", cur.GetValue())
		cur = cur.GetNext()
	}
	fmt.Println()
}
