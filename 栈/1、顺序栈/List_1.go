package main

type List struct {
	maxSize int // 栈的最大长度
	top     int // 指向栈顶的位置，-1表示空栈
	data    []int
}

const (
	ERROR = -1
	OK    = 1
)

func (l *List) Init(size int) *List {
	return &List{
		maxSize: size,
		top:     -1,
		data:    make([]int, size),
	}
}

func (l *List) Push(v int) int {
	// 数组下标等于长度-1
	if l.top == l.maxSize-1 { // 栈满
		return ERROR
	}
	l.top++
	l.data[l.top] = v
	return OK
}

// 删除栈顶元素，并用e返回他的值
func (l *List) Pop(e *int) int {
	if l.top == -1 { // 空栈
		return ERROR
	}
	*e = l.data[l.top]
	l.top--
	return OK
}
