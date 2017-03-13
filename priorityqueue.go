package container

import (
	"container/heap"
	"sync"
)

// Node 优先队列中的基本元素
type Node struct {
	Prior int
	Value string
}

// PriorityQueue 对外暴露的优先队列
type PriorityQueue struct {
	queue *priorityQueue
}

// NewPriorityQueue 创建队列
func NewPriorityQueue(origin []*Node) *PriorityQueue {
	q := &priorityQueue{
		nodes: origin,
	}

	heap.Init(q)

	return &PriorityQueue{
		queue: q,
	}
}

// Pop 弹出优先级最高的队列
func (pq *PriorityQueue) Pop() *Node {
	node := heap.Pop(pq.queue)
	return node.(*Node)
}

// Push 压入新的元素
func (pq *PriorityQueue) Push(node *Node) {
	heap.Push(pq.queue, node)
}

// Len 获取队列长度
func (pq *PriorityQueue) Len() int {
	return pq.queue.Len()
}

// 实现heap.Interface接口实现的最小堆
type priorityQueue struct {
	nodes []*Node
	guard sync.RWMutex
}

func (q *priorityQueue) Len() int {
	q.guard.RLock()
	length := len(q.nodes)
	q.guard.RUnlock()
	return length
}

func (q *priorityQueue) Less(i, j int) bool {
	q.guard.RLock()
	isLess := q.nodes[i].Prior < q.nodes[j].Prior
	q.guard.RUnlock()
	return isLess
}

func (q *priorityQueue) Swap(i, j int) {
	q.guard.Lock()
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
	q.guard.Unlock()
}

func (q *priorityQueue) Push(x interface{}) {
	node := x.(*Node)
	q.guard.Lock()
	q.nodes = append(q.nodes, node)
	q.guard.Unlock()
}

func (q *priorityQueue) Pop() interface{} {
	q.guard.Lock()
	old := q.nodes
	n := len(old)
	node := old[n-1]
	q.nodes = old[0 : n-1]
	q.guard.Unlock()
	return node
}
