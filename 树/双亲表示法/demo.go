package main

type Node struct {
	data   int // 结点存储的数据
	parent int // 结点双亲所在数组下标
}

type Tree struct {
	TNodes []Node // 树
	n      int    // 节点数
	r      int    // 根节点的位置
}
