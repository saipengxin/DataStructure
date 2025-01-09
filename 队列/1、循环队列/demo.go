package main

type Queue struct {
	data    []int
	front   int // 头指针
	rear    int // 尾指针
	maxSize int // 队列最大长度
}

const (
	ERROR = -1
	OK    = 1
)

// Init 初始化一个队列
func Init(maxSize int) *Queue {
	q := new(Queue)
	q.data = make([]int, maxSize)
	q.front = 0
	q.rear = 0
	q.maxSize = maxSize
	return q
}

// QueueLength 队列长度
func (q *Queue) QueueLength() int {
	return (q.rear - q.front + q.maxSize) % q.maxSize
}

// EnQueue 入队
func (q *Queue) EnQueue(e int) int {
	if (q.rear+1)%q.maxSize == q.front {
		return ERROR // 队列满报错
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % q.maxSize // 尾指针有可能要移动到前面的位置，所以不能用 q.rear + 1 来处理移动尾指针
	return OK
}

// DeQueue 出队，并用e返回它的值
func (q *Queue) DeQueue(e *int) int {
	if q.rear == q.front {
		return ERROR // 队列空报错
	}
	*e = q.data[q.front]
	q.front = (q.front + 1) % q.maxSize
	return OK
}
