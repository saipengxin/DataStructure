package main

import (
	"fmt"
)

type TreeNode struct {
	Left, Right *TreeNode
	Data        int
}

func main() {
	// 构造示例二叉树
	root := &TreeNode{Data: 1,
		Left:  &TreeNode{Data: 2, Left: &TreeNode{Data: 4}, Right: &TreeNode{Data: 5}},
		Right: &TreeNode{Data: 3, Left: &TreeNode{Data: 6}, Right: &TreeNode{Data: 7}},
	}

	// 执行层序遍历并打印结果
	levelOrder(root)
}

// 层序遍历二叉树
func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root} // 初始化队列，将根节点入队
	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点个数
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // 取出队列头部元素
			queue = queue[1:] // 获取到列头元素之后，我们需要将列头元素去掉，以便下一次循环的时候渠道的不是重复的数据
			fmt.Println(node.Data)
			if node.Left != nil {
				queue = append(queue, node.Left) // 左子节点入队
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // 右子节点入队
			}
		}
	}
}
