package main

import "fmt"

func main() {

}

type Node struct {
	LTag  int8   // 标识
	Left  *Node  // 左孩子
	Data  string // 数据
	Right *Node  // 右孩子
	RTag  int8   // 标识
}

func InOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}
	p := tree.Left  // 找到头结点的左孩子
	for p != tree { // 如果线索二叉树遍历完成，最后会回到头结点
		for p.LTag == 0 { // 有左孩子就一直找到左孩子的最后一个结点
			p = p.Left
		}
		fmt.Println(p.Data)
		for p.RTag == 1 && p.Right != tree { // 存在后继结点就顺着后继结点一直走
			p = p.Right
			fmt.Println(p.Data)
		}
		p = p.Right // 当我们碰到真正的右子树，跳进去然后再从右子树中找到最左边的结点
	}
}
