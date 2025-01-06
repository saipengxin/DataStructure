package main

// 栈的实现
type ListNode struct {
	data interface{}
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

func (list *List) Push(e interface{}) int {
	node := &ListNode{
		data: e,
		next: list.top,
	}
	list.top = node
	list.count++
	return OK
}

func (list *List) Pop(e *interface{}) int {
	if list.top == nil {
		return ERROR
	}
	*e = list.top.data
	list.top = list.top.next
	list.count--
	return OK
}
