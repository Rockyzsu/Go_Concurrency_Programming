//Go单链表实现
package linklist

type ListNode struct {
	data int
	next *ListNode
}

func CreateLinkNode(value int) *ListNode {
	return &ListNode{
		data: value,
		next: nil,
	}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() int {
	return this.data
}
