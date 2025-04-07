package main

import "fmt"

// 后序遍历算法

func main() {}

type TreeNode struct {
	Left, Right *TreeNode
	Data        int
}

// InOrderTraverse 中序遍历算法
func InOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	InOrderTraverse(tree.Left)  // 遍历左子树
	fmt.Println(tree.Data)      //打印结点数据
	InOrderTraverse(tree.Right) // 遍历右子树
}
