package main

type ListNode struct {
	data int
	next *ListNode
}

type List struct {
	count int
	top   *ListNode
}

func NewList() *List {
	return &List{
		count: 0,
		top:   nil,
	}
}
