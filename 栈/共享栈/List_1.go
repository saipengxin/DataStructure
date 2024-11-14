package main

type List struct {
	maxSize int // 栈的最大长度
	top1    int // 指向栈1栈顶的位置，-1表示空栈
	top2    int // 指向栈2栈顶的位置，n表示空栈
	data    []int
}

const (
	ERROR = -1
	OK    = 1
)

func NewList(size int) *List {
	return &List{
		maxSize: size,
		top1:    -1,
		top2:    size,
		data:    make([]int, size),
	}
}

func (list *List) Push(value int, stackNumber int) int {
	if list.top1+1 == list.top2 { // 栈满
		return ERROR
	}
	if stackNumber == 1 { // 栈1
		list.top1++
		list.data[list.top1] = value
	} else if stackNumber == 2 { // 栈2
		list.top2--
		list.data[list.top2] = value
	}
	return OK
}

func (list *List) Pop(e *int, stackNumber int) int {
	if stackNumber == 1 {
		if list.top1 == -1 { // 栈满
			return ERROR
		}
		*e = list.data[list.top1]
		list.top1--
	} else if stackNumber == 2 {
		if list.top2 == list.maxSize { // 栈满
			return ERROR
		}
		*e = list.data[list.top2]
		list.top2++
	}
	return OK
}
