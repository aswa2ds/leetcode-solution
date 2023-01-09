package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	emptyHead := new(ListNode)
	emptyHead.Next = head
	pre := emptyHead
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	// pre 指向left 左边的节点
	h, t := pre.Next, pre.Next
	for i := 0; i < right-left; i++ {
		h = t.Next
		t.Next = h.Next
		h.Next = pre.Next
		pre.Next = h
	}

	return emptyHead.Next
}
