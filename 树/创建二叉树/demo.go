package main

import (
	"encoding/json"
	"fmt"
)

// Node 采用前序遍历的方式创建二叉树
type Node struct {
	Left  *Node  // 左孩子
	Data  string // 数据
	Right *Node  // 右孩子
}

func main() {
	tree := "AB#D##C##"
	root := createTree(tree)
	fmt.Println("==============================")
	t, _ := json.Marshal(root)
	fmt.Printf("%#v\n", string(t))
	fmt.Println("==============================")
}

func createTree(tree string) *Node {
	index := 0

	var create func() *Node
	create = func() *Node {
		if index > len(tree)-1 {
			return nil
		}
		data := tree[index]
		index++ // 用闭包来控制index，不用一层层的传递index下去
		var root *Node
		if data == '#' {
			return nil
		}
		root = &Node{Data: string(data)}
		root.Left = create()
		root.Right = create()
		return root
	}

	return create()
}
