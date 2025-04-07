package main

import "fmt"

// 前序遍历二叉树
func main() {

}

type TreeNode struct {
	Left, Right *TreeNode
	Data        int
}

// PreOrderTraverse 前序遍历算法
func PreOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Data)       //打印结点数据
	PreOrderTraverse(tree.Left)  // 遍历左子树
	PreOrderTraverse(tree.Right) // 遍历右子树
}
