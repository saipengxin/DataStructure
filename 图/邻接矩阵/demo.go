package main

import "fmt"

type MGraph struct {
	Vertex   []int   // 顶点表
	Arc      [][]int // 邻接矩阵，边表
	numNodes int     // 顶点数
	numEdges int     // 边数
}

var visited map[int]bool

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

// DFS 连接矩阵深度优先遍历递归算法
func (m *MGraph) DFS(index int) {
	if m.numNodes == 0 {
		return
	}
	// 将对应的结点设置为已访问
	visited[index] = true

	// 对顶点进行业务操作
	fmt.Println(m.Vertex[index])

	// 执行DFS遍历，从指定的index顶点出发，一路遍历他的连通顶点
	for j, v := range m.Arc[index] {
		// 顶点被访问过不再重复访问
		if visited[j] {
			continue
		}
		if v == 1 {
			// 1 表示两个顶点之间有边，顺着边找到下一个顶点j，然后再对j执行DFS
			// 这里执行递归是因为，如果顶点0到顶点1和顶点2都存在边，那么这里判断m.Arc[0][1] = 1就会进入递归，开始对顶点1进行DFS
			// 那么顶点2就被我们漏掉了，没有访问，所以我们要再递归完成之后，再一层层的向上回退，根据边找到没有被访问过的其他顶点，
			// 我们从顶点0出发，最后应该再回到顶点0
			m.DFS(j)
		}
	}
}

// DFSTraverse 邻接矩阵的深度遍历操作，确保每一个结点都执行深度遍历
func (m *MGraph) DFSTraverse() {
	if m.numNodes == 0 {
		return
	}

	// 先将所有结点都设置为未访问
	visited = make(map[int]bool)
	for i := 0; i < len(m.Vertex); i++ {
		visited[i] = false
	}

	// 执行DFS遍历
	for index, _ := range m.Vertex {
		// 如果是连通图，第一次执行DFS就会将所有顶点都遍历完成，所以只执行一遍
		// 循环是为了防止非连通图，第一次遍历只能遍历一个连通分量，还有剩余的顶点不能遍历到
		if visited[index] {
			continue
		}
		m.DFS(index)
	}
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

	matrix.DFSTraverse()
}
