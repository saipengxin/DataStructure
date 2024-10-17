package main

// 静态链表 (用数组实现的链表)

const (
	MaxSize = 1000
	ERROR   = -1
	OK      = 1
)

type ListNode struct {
	data int // 数据元素
	cur  int // 后继元素下标存储
}

var List [MaxSize]ListNode

func InitList() int {
	for i := 0; i < MaxSize-1; i++ {
		List[i].cur = i + 1
	}
	List[MaxSize-1].cur = 0 // 初始化为空链表，最后一个元素的cur为0
	return OK
}

// Malloc_SSL 返回插入元素的放置位置
func Malloc_SSL() int {
	i := List[0].cur // 获取插入元素的放置位置
	if i != 0 {
		List[0].cur = List[i].cur // 新元素插入，备用链表的第一个元素要移动到下一个元素
	}
	return i
}

func ListLength() int {
	j := 0
	i := List[MaxSize-1].cur
	for i != 0 {
		i = List[i].cur
		j++
	}
	return j
}

func ListInsert(i int, e int) int {
	if i < 1 || i > ListLength()+1 {
		return ERROR
	}
	j := Malloc_SSL()
	if j != 0 {
		List[j].data = e
		k := 0
		for i := 1; i <= i-1; i++ {
			k = List[k].cur
		}
		List[j].cur = List[k].cur
		List[k].cur = j
		return OK
	}
	return ERROR
}
