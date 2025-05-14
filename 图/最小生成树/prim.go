package main

import (
	"fmt"
	"math"
)

// 定义一个常量表示无穷大
const INF = math.MaxInt32

// Prim 算法实现，返回最小生成树的边和总权重
func prim(graph [][]int) (mst []Edge, totalWeight int) {
	n := len(graph)
	visited := make([]bool, n)
	lowCost := make([]int, n)  // 当前到 MST 的最小边权
	adjacent := make([]int, n) // 与 lowCost 相对应的前驱节点

	// 初始化
	for i := 0; i < n; i++ {
		lowCost[i] = graph[0][i]
		adjacent[i] = 0
	}
	visited[0] = true // 从第 0 个顶点开始

	for i := 1; i < n; i++ {
		min := INF
		k := -1
		// 找出最小权重边
		for j := 0; j < n; j++ {
			if !visited[j] && lowCost[j] < min {
				min = lowCost[j]
				k = j
			}
		}
		fmt.Println("==============================")
		fmt.Printf("%#v\n", i)
		fmt.Printf("%#v\n", lowCost)
		fmt.Printf("%#v\n", adjacent)
		fmt.Printf("%#v\n", k)
		fmt.Println("==============================")
		if k == -1 {
			// 图不连通
			fmt.Println("图不连通")
			return
		}

		visited[k] = true
		totalWeight += min
		mst = append(mst, Edge{From: adjacent[k], To: k, Weight: min})

		// 更新其他点的最小权重
		for j := 0; j < n; j++ {
			if !visited[j] && graph[k][j] < lowCost[j] {
				lowCost[j] = graph[k][j]
				adjacent[j] = k
			}
		}
	}

	return
}

// 表示一条边
type Edge struct {
	From, To int
	Weight   int
}

func main() {
	// 使用邻接矩阵表示图，INF 表示无连接
	graph := [][]int{
		{0, 10, INF, INF, INF, 11, INF, INF, INF},
		{10, 0, 18, INF, INF, INF, 16, INF, 12},
		{INF, 18, 0, 22, INF, INF, INF, INF, 8},
		{INF, INF, 22, 0, 20, INF, 24, 16, 21},
		{INF, INF, INF, 20, 0, 26, INF, 7, INF},
		{11, INF, INF, INF, 26, 0, 17, INF, INF},
		{INF, 16, INF, 24, INF, 17, 0, 19, INF},
		{INF, INF, INF, 16, 7, INF, 19, 0, INF},
		{INF, 12, 8, 21, INF, INF, INF, INF, 0},
	}

	mst, total := prim(graph)
	fmt.Println("最小生成树边:")
	for _, e := range mst {
		fmt.Printf("%d -- %d (权重: %d)\n", e.From, e.To, e.Weight)
	}
	fmt.Println("最小生成树总权重:", total)
}
