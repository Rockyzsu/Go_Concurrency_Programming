package linklist

import "errors"

func (this *LinkList) Find(index uint) (*ListNode, error) {
	if index < 0 || index >= this.length {
		return nil, errors.New("越界")
	}

	cur := this.head.GetNext()
	for i := uint(0); i < index; i++ {
		cur = cur.GetNext()
	}

	return cur, nil
}
