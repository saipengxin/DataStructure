package main

import "fmt"

// 静态链表 (用数组实现的链表)

const (
	MaxSize = 10
	ERROR   = -1
	OK      = 1
)

type ListNode struct {
	data int // 数据元素
	cur  int // 后继元素下标存储
}

type List struct {
	maxSize int
	length  int
	list    []ListNode
}

func InitList(size int) (list List) {
	list = List{
		maxSize: size,
		length:  0,
		list:    make([]ListNode, size),
	}

	for i := 0; i < size-1; i++ {
		list.list[i].cur = i + 1
	}
	list.list[size-1].cur = 0 // 初始化为空链表，最后一个元素的cur为0
	return list
}

// Malloc_SSL 返回插入元素的放置位置
func (list *List) Malloc_SSL() int {
	i := list.list[0].cur // 获取插入元素的放置位置
	if i != 0 {
		list.list[0].cur = list.list[i].cur // 新元素插入，备用链表的第一个元素要移动到下一个元素
	}
	return i
}

func (list *List) ListLength() int {
	return list.length
}

func (list *List) ListInsert(i int, e int) int {
	if i < 1 || i > list.ListLength()+1 {
		return ERROR
	}
	k := list.maxSize - 1
	j := list.Malloc_SSL()
	if j != 0 {
		list.list[j].data = e
		// 找到要操作位置的前一个元素
		for l := 1; l <= i-1; l++ {
			k = list.list[k].cur
		}
		list.list[j].cur = list.list[k].cur
		list.list[k].cur = j

		list.length = list.length + 1
		return OK
	}
	return ERROR
}

func (list *List) ListDelete(i int) int {
	if i < 1 || i > list.ListLength()+1 {
		return ERROR
	}
	k := list.maxSize - 1
	// 找到要操作位置的前一个元素
	for j := 1; j <= i-1; j++ {
		k = list.list[k].cur
	}
	// 跳过要删除元素的下标
	list.list[k].cur = list.list[list.list[k].cur].cur
	list.length = list.length - 1
	return OK
}

func main() {
	list := InitList(MaxSize)
	fmt.Println("==============================")
	fmt.Printf("%#v\n", list)
	fmt.Println("==============================")

	list.ListInsert(1, 10)
	fmt.Println("==============================")
	fmt.Printf("%#v\n", list)
	fmt.Println("==============================")

	list.ListInsert(1, 20)
	fmt.Println("==============================")
	fmt.Printf("%#v\n", list)
	fmt.Println("==============================")

	list.ListInsert(2, 55)
	fmt.Println("==============================")
	fmt.Printf("%#v\n", list)
	fmt.Println("==============================")

	list.ListDelete(1)
	fmt.Println("==============================")
	fmt.Printf("%#v\n", list)
	fmt.Println("==============================")
}
