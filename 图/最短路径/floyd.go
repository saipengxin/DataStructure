package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt32

func floydWarshall(graph [][]int) ([][]int, [][]int) {
	n := len(graph)
	// 初始化距离矩阵和路径矩阵
	dist := make([][]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = graph[i][j]
			if graph[i][j] != INF && i != j {
				next[i][j] = j
			} else {
				next[i][j] = -1
			}
		}
	}

	// Floyd-Warshall 三重循环
	// 每次都尝试使用k作为中间节点看看能不能优化路径长度
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != INF && dist[k][j] != INF && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}

	return dist, next
}

// 用 next 表还原路径
func constructPath(u, v int, next [][]int) []int {
	if next[u][v] == -1 {
		return nil
	}
	path := []int{u}
	for u != v {
		u = next[u][v]
		path = append(path, u)
	}
	return path
}

func main() {
	// 示例图：有向图，4 个节点
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

	dist, next := floydWarshall(graph)

	fmt.Println("最短路径长度矩阵:")
	for _, row := range dist {
		for _, val := range row {
			if val == INF {
				fmt.Print("INF\t")
			} else {
				fmt.Printf("%d\t", val)
			}
		}
		fmt.Println()
	}

	fmt.Println("\n路径示例（从 0 到 3 的路径）:")
	path := constructPath(0, 3, next)
	if path == nil {
		fmt.Println("不可达")
	} else {
		fmt.Println(path)
	}
}
