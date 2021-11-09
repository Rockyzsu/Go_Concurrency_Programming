package linklist

import "errors"

//Insert 节点插入,i是节点索引,从0开始;node是需要插入的节点
func (this *LinkList) Insert(i uint, node *ListNode) error {
	if i < 0 || node == nil || i > this.length {
		return errors.New("节点为空或越界")
	}
	//从head依次循环定位到索引i
	curNode := (*this).head
	for j := uint(0); j < i; j++ {
		curNode = curNode.GetNext()
	}
	node.next = curNode.GetNext()
	curNode.next = node
	this.length++
	return nil
}
