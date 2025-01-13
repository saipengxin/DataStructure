package main

// Node 节点
type Node struct {
	data int
	next *Node
}

// Queue 队列
type Queue struct {
	front *Node // 头指针
	rear  *Node // 尾指针
}

func Init() *Queue {
	head := new(Node) // 头结点，空队列 头指针、尾指针都指向头结点
	q := new(Queue)
	q.front = head
	q.rear = head
	return q
}

// EnQueue 入队
func (q *Queue) EnQueue(e int) {
	node := &Node{e, nil}
	q.rear.next = node
	q.rear = node
}

// DeQueue 出队
func (q *Queue) DeQueue(e *int) {
	if q.front == q.rear {
		return // 空队列
	}
	p := q.front.next
	*e = p.data
	q.front.next = p.next
	if q.rear == p { // 如果队头就是队尾，删除队头之后将尾指针指向队列的头结点
		q.rear = q.front
	}
}
