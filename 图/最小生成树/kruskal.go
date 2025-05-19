package main

import (
	"fmt"
	"sort"
)

// 边结构体
type EdgeStruct struct {
	u, v, weight int // 两个端点，权重
}

// 并查集结构体
type UnionFind struct {
	parent []int
}

// 初始化并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent}
}

// 查找根节点（带路径压缩）
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// 合并两个集合
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return false
	}
	uf.parent[rootY] = rootX
	return true
}

// Kruskal 算法主函数
func Kruskal(n int, edges []EdgeStruct) ([]EdgeStruct, int) {
	// 1. 按权重排序，遍历的边集要是从小到大排序的
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	uf := NewUnionFind(n)
	mst := []EdgeStruct{}
	totalWeight := 0

	// 2. 遍历边
	for _, edge := range edges {
		// 这段用于判断是否形成环路
		if uf.Union(edge.u, edge.v) {
			mst = append(mst, edge)
			totalWeight += edge.weight
		}
	}

	return mst, totalWeight
}

// 示例使用
func main() {
	edges := []EdgeStruct{
		{4, 7, 7},
		{2, 8, 8},
		{0, 1, 10},
		{0, 5, 11},
		{1, 8, 12},
		{3, 7, 16},
		{1, 6, 16},
		{5, 6, 17},
		{1, 2, 18},
		{6, 7, 19},
		{3, 4, 20},
		{3, 8, 21},
		{2, 3, 22},
		{3, 6, 24},
		{4, 5, 26},
	}
	n := 9 // 顶点个数

	mst, total := Kruskal(n, edges)
	fmt.Println("最小生成树的边:")
	for _, e := range mst {
		fmt.Printf("%d -- %d == %d\n", e.u, e.v, e.weight)
	}
	fmt.Printf("总权重: %d\n", total)
}
