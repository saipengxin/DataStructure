package main

import "fmt"

type MGraph struct {
	Vertex   []int   // 顶点表
	Arc      [][]int // 邻接矩阵，边表
	numNodes int     // 顶点数
	numEdges int     // 边数
}

// CreateMGraph 构建一个无向图
// n: 顶点集合
// edges: 边集合
func CreateMGraph(n []int, edges [][2]int) *MGraph {
	m := new(MGraph)
	num := len(n)
	m.numNodes = num
	m.numEdges = len(edges)
	m.Vertex = n
	m.Arc = make([][]int, num)
	for i := 0; i < num; i++ {
		m.Arc[i] = make([]int, num)
	}

	for _, edge := range edges {
		if edge[0] > num || edge[0] < 0 || edge[1] > num || edge[1] < 0 {
			continue
		}
		m.Arc[edge[0]][edge[1]] = 1
		m.Arc[edge[1]][edge[0]] = 1 // 无向图，设置对称
	}
	return m
}

func main() {
	n := []int{0, 1, 2, 3, 4}
	edges := [][2]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{3, 4},
	}

	matrix := CreateMGraph(n, edges)
	fmt.Println("Adjacency Matrix:")
	for _, row := range matrix.Arc {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
