package main

// Node 定义单链表节点结构体
type Node struct {
	data int   // 节点数据存储
	next *Node // 指向下一个节点的指针
}

// 定义链表结构体
type List struct {
	head *Node // 指向头结点的指针
}

const (
	ERROR = -1 // 获取失败的返回结果
	OK    = 1  // 操作成功返回结果
)

// GetElem 找到链表L中第i个位置的元素，并用e返回 （注意：链表存在头结点）
func (L *List) GetElem(i int, e *int) int {
	j := 1          // 计数器
	p := &Node{}    // 空节点，用于从头开始遍历数据
	p = L.head.next // 空节点指向链表的第一个元素节点
	for p != nil && j < i {
		p = p.next
		j++ // j累加
	}
	if p == nil || j > i {
		return ERROR
	}
	*e = p.data
	return OK
}
