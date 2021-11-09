package linklist

import "errors"

func (this *LinkList) Delete(node *ListNode) error {
	if nil == node {
		return errors.New("节点为空")
	}
	pre := this.head
	cur := this.head.GetNext() //头指向的节点
	//循环直到链表尾端nil
	for nil != cur {
		if cur.GetValue() == node.GetValue() {
			break
		}
		pre = cur
		cur = cur.GetNext()
	}
	if nil == cur {
		return errors.New("未找到节点")
	}
	pre.next = cur.GetNext()
	node = nil //删除节点
	this.length--
	return nil
}
