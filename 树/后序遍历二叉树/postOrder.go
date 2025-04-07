package main

import "fmt"

// 后序遍历算法

func main() {}

type TreeNode struct {
	Left, Right *TreeNode
	Data        int
}

// PostOrderTraverse 后序遍历算法
func PostOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	PostOrderTraverse(tree.Left)  // 遍历左子树
	PostOrderTraverse(tree.Right) // 遍历右子树
	fmt.Println(tree.Data)        // 打印结点数据
}
