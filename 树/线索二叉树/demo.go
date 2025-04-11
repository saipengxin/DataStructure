package main

type Node struct {
	LTag  int8   // 标识
	Left  *Node  // 左孩子
	Data  string // 数据
	Right *Node  // 右孩子
	RTag  int8   // 标识
}

func InTraverse(tree *Node) {
	pre := &Node{} // 前序结点,采用闭包的方式保证执行期间只初始化一次

	var create func(*Node)
	create = func(tree *Node) {
		if tree == nil {
			return
		}
		create(tree.Left)     // 遍历左子树
		if tree.Left == nil { // 左孩子为空，设置为前序线索
			tree.LTag = 1
			tree.Left = pre
		}
		if pre.Right == nil { // 如果前驱结点的右孩子为空，那当前结点其实就是前驱结点的后继了
			pre.RTag = 1
			pre.Right = tree
		}
		pre = tree         // 始终让pre指向前驱结点
		create(tree.Right) // 遍历右子树
	}
	create(tree)
}
