package main

type ListNode struct {
	data int
	next *ListNode
}

type List struct {
	count int
	top   *ListNode
}

const (
	ERROR = -1
	OK    = 1
)

func NewList() *List {
	return &List{
		count: 0,
		top:   nil, // 栈顶，null的时候为空栈
	}
}

func (list *List) Push(e int) int {
	node := &ListNode{
		data: e,
		next: list.top,
	}
	list.top = node
	list.count++
	return OK
}
