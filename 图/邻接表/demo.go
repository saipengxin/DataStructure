package main

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

func main() {}
