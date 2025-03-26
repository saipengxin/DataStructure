package main

// 孩子结点的结构
type ChildNode struct {
	child int        // 孩子结点的下标
	next  *ChildNode // 指针域，用来指向下一个孩子结点
}

// 树节点
type TreeNode struct {
	data       string     // 数据域，存储结点的数据信息
	firstChild *ChildNode // 指向孩子链表的头指针
}

type Tree struct {
	root []TreeNode // 树
	n    int        // 节点数
}
