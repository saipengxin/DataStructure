package main

import (
	"container/list"
	"fmt"
)

// 顶点表结构
type VertexNode struct {
	data      int       // 数据域
	firstedge *EdgeNode // 指针域，指向边表的头指针
}

type EdgeNode struct {
	adjvex int       // 邻接点对应下标
	weight int       // 权值，非网图可以省略
	next   *EdgeNode // 下一个邻接点指针
}

type GraphAdjList struct {
	numNodes int
	numEdges int
	adjList  []VertexNode
}

var visited map[int]bool

// 创建图（邻接表）
// 参数：顶点数量，边列表（每条边表示为 [from, to, weight]）
func CreateALGraph(numNodes int, edges [][3]int) *GraphAdjList {
	graph := &GraphAdjList{
		numNodes: numNodes,
		numEdges: len(edges),
		adjList:  make([]VertexNode, numNodes),
	}

	// 初始化顶点 data（你也可以换成其他信息）
	for i := 0; i < numNodes; i++ {
		graph.adjList[i].data = i
		graph.adjList[i].firstedge = nil
	}

	// 添加边
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		weight := edge[2]

		// 插入边（头插法）
		edgeNode := &EdgeNode{
			adjvex: to,
			weight: weight,
			next:   graph.adjList[from].firstedge,
		}
		graph.adjList[from].firstedge = edgeNode

		// 如果是无向图，还需要添加反向边（注释掉下面这段可以变为有向图）
		edgeNodeUndir := &EdgeNode{
			adjvex: from,
			weight: weight,
			next:   graph.adjList[to].firstedge,
		}
		graph.adjList[to].firstedge = edgeNodeUndir
	}

	return graph
}

// 深度优先遍历
func (m *GraphAdjList) DFS(index int) {
	if m.numNodes == 0 {
		return
	}
	// 将对应的结点设置为已访问
	visited[index] = true

	// 对顶点进行业务操作
	fmt.Println(m.adjList[index].data)

	for p := m.adjList[index].firstedge; p != nil; p = p.next {
		if !visited[p.adjvex] {
			m.DFS(p.adjvex)
		}
	}
}

// DFSTraverse 邻接矩阵的深度遍历操作，确保每一个结点都执行深度遍历
func (m *GraphAdjList) DFSTraverse() {
	if m.numNodes == 0 {
		return
	}

	// 先将所有结点都设置为未访问
	visited = make(map[int]bool)
	for i := 0; i < len(m.adjList); i++ {
		visited[i] = false
	}

	// 执行DFS遍历
	for index, _ := range m.adjList {
		if visited[index] {
			continue
		}
		m.DFS(index)
	}
}

// ============================

func (m *GraphAdjList) BFS(index int) {
	if m.numNodes == 0 {
		return
	}
	queue := list.New()
	queue.PushBack(index)
	visited[index] = true
	for queue.Len() > 0 {
		e := queue.Remove(queue.Front()).(int)

		fmt.Println(m.adjList[e].data)

		// 访问当前结点的所有邻接结点，并放入队列
		for p := m.adjList[e].firstedge; p != nil; p = p.next {
			if !visited[p.adjvex] {
				queue.PushBack(p.adjvex)
				visited[p.adjvex] = true
			}
		}
	}
}

func (m *GraphAdjList) BFSTraverse() {
	if m.numNodes == 0 {
		return
	}

	// 先将所有结点都设置为未访问
	visited = make(map[int]bool)
	for i := 0; i < len(m.adjList); i++ {
		visited[i] = false
	}

	// 保证所有连通分量都能被遍历
	for index, _ := range m.adjList {
		if visited[index] {
			continue
		}
		m.BFS(index)
	}
}

func main() {
	edges := [][3]int{
		{0, 1, 1},
		{0, 2, 1},
		{1, 2, 1},
		{3, 4, 1},
	}

	matrix := CreateALGraph(5, edges)
	fmt.Println("Adjacency Matrix:")
	matrix.DFSTraverse()
}
