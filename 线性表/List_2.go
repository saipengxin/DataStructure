package main

import (
	"fmt"
	"math/rand/v2"
)

// 线性表链式存储结构

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

// ListInsert 在链表L的第i个位置插入e
func (L *List) ListInsert(i int, e int) int {
	j := 1       // 计数器
	p := &Node{} // 空节点，用于从头开始遍历数据
	p = L.head   // 空节点指向链表的头节点
	for p != nil && j < i {
		p = p.next
		j++ // j累加
	}
	if p == nil || j > i {
		return ERROR
	}
	// 设置一个空节点存数据
	s := &Node{}
	s.data = e
	// 这两行不能写反，否则数据会被覆盖
	s.next = p.next
	p.next = s
	return OK
}

// ListDelete 删除链表L的第i个位置，并用e返回数据
func (L *List) ListDelete(i int, e *int) int {
	j := 1       // 计数器
	p := &Node{} // 空节点，用于从头开始遍历数据
	p = L.head   // p指向链表的头节点 （注意我们链表的开头有一个不存数据的头节点，用于统一操作使用）
	for p != nil && j < i {
		p = p.next
		j++ // j累加
	}
	if p == nil || j > i {
		return ERROR
	}
	q := p.next
	p.next = q.next
	*e = q.data
	return OK
}

// CreateListHead 初始化链表，随机生成n个值（头插法）
func (L *List) CreateListHead(n int) {
	L.head.next = nil
	for i := 0; i < n; i++ {
		a := rand.IntN(100)
		p := &Node{
			data: a,
		}
		p.next = L.head.next
		L.head.next = p // 数据插到表头
	}
}

// CreateListTail 初始化链表，随机生成n个值（尾插法）
func (L *List) CreateListTail(n int) {
	L.head.next = nil
	h := L.head // h为指向链表尾部的指针
	for i := 0; i < n; i++ {
		a := rand.IntN(100)
		p := &Node{
			data: a,
		}
		h.next = p
		h = p // 保证h指向尾部
	}
	h.next = nil // 遍历生成完成后，尾结点指向null
}

// ClearList 链表重制为空
func (L *List) ClearList() {
	L.head.next = nil
}

// Print 打印链表数据
func (L *List) Print() {
	p := L.head.next
	for true {
		fmt.Print(p.data, " ")
		p = p.next
		if p == nil {
			break
		}
	}
}

func main() {}
