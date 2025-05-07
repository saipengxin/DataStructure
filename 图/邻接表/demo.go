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

func CreateALGraph() {

}

func main() {}
