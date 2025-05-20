package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt64 // 代表不可达的距离

// Dijkstra 算法 - 邻接矩阵实现
func DijkstraMatrix(graph [][]int, start int) []int {
	n := len(graph)            // 顶点个数
	dist := make([]int, n)     // 最短距离数组
	visited := make([]bool, n) // 是否访问过

	// 初始化：将起点到所有点的距离设置为图中原始值
	for i := 0; i < n; i++ {
		dist[i] = graph[start][i]
	}
	dist[start] = 0
	visited[start] = true

	// 遍历 n-1 个点
	for i := 0; i < n-1; i++ {
		u := -1
		minDist := INF

		// 找出当前未访问节点中距离最小的点 u
		for j := 0; j < n; j++ {
			if !visited[j] && dist[j] < minDist {
				u = j
				minDist = dist[j]
			}
		}

		if u == -1 {
			break // 所有可达节点处理完毕
		}

		visited[u] = true

		// 更新 u 的邻居节点的最短距离
		for v := 0; v < n; v++ {
			if !visited[v] && graph[u][v] != INF {
				if dist[v] > dist[u]+graph[u][v] {
					dist[v] = dist[u] + graph[u][v]
				}
			}
		}
	}

	return dist
}

func main() {
	INF := math.MaxInt64
	graph := [][]int{
		{0, 1, 1, INF, INF, INF, INF, INF, INF},
		{1, 0, 3, 7, 5, INF, INF, INF, INF},
		{1, 3, 0, INF, 1, 7, INF, INF, INF},
		{INF, 7, INF, 0, 2, INF, 3, INF, INF},
		{INF, 5, 1, 2, 0, 3, 6, 9, INF},
		{INF, INF, 7, INF, 3, 0, INF, 5, INF},
		{INF, INF, INF, 3, 6, INF, 0, 2, 7},
		{INF, INF, INF, INF, 9, 5, 2, 0, 4},
		{INF, INF, INF, INF, INF, INF, 7, 4, 0},
	}

	start := 0 // 从 A 开始
	dist := DijkstraMatrix(graph, start)

	nodeNames := []string{"V0", "V1", "V2", "V3", "V4", "V5", "V6", "V7", "V8"}
	fmt.Println("从 V0 出发的最短路径：")
	for i, d := range dist {
		fmt.Printf("V0 -> %s: %d\n", nodeNames[i], d)
	}
}
