package main

import (
	"fmt"
)

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

func (list *List) Pop(e *int) int {
	if list.top == nil {
		return ERROR
	}
	*e = list.top.data
	list.top = list.top.next
	list.count--
	return OK
}

func main() {
	list := NewList()
	list.Push(1)
	list.Push(2)
	list.Push(3)

	fmt.Println("==============================")
	fmt.Printf("%#v\n", list)
	fmt.Printf("%#v\n", list.top)
	fmt.Printf("%#v\n", list.top.next)
	fmt.Println("==============================")

	e := 0
	list.Pop(&e)
	fmt.Println("==============================")
	fmt.Printf("%#v\n", list)
	fmt.Printf("%#v\n", list.top)
	fmt.Printf("%#v\n", list.top.next)
	fmt.Printf("%#v\n", e)
	fmt.Println("==============================")
}
