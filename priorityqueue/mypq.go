package priorityqueue

import "errors"

type MyPriorityQueue struct {
	heap       []interface{}
	size       int
	comparator func(x, y interface{}) int
}

func NewMyPriorityQueue(capacity int, comparator func(x, y interface{}) int) *MyPriorityQueue {
	return &MyPriorityQueue{
		heap:       make([]interface{}, capacity),
		size:       0,
		comparator: comparator,
	}
}

func (pq *MyPriorityQueue) Size() int {
	return pq.size
}

func (pq *MyPriorityQueue) IsEmpty() bool {
	return pq.size == 0
}

func (pq *MyPriorityQueue) Parent(node int) int {
	return (node - 1) / 2
}

func (pq *MyPriorityQueue) Left(node int) int {
	return node*2 + 1
}

func (pq *MyPriorityQueue) Right(node int) int {
	return node*2 + 2
}

func (pq *MyPriorityQueue) Swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *MyPriorityQueue) Peek() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("pq underflow")
	}
	return pq.heap[0], nil
}

func (pq *MyPriorityQueue) Push(x interface{}) {
	if pq.size == len(pq.heap) {
		pq.resize(2 * len(pq.heap))
	}
	pq.heap[pq.size] = x
	pq.swim(pq.size)
	pq.size++
}

func (pq *MyPriorityQueue) resize(capacity int) {
	newHeap := make([]interface{}, capacity)
	for i := 0; i < capacity; i++ {
		newHeap[i] = pq.heap[i]
	}
	pq.heap = newHeap
}

func (pq *MyPriorityQueue) swim(node int) {
	for node > 0 && pq.comparator(pq.heap[pq.Parent(node)], pq.heap[node]) > 0 {
		pq.Swap(node, pq.Parent(node))
		node = pq.Parent(node)
	}
}

func (pq *MyPriorityQueue) sink(node int) {
	for pq.Left(node) < pq.size {
		minNode := node
		if pq.Left(node) < pq.size && (pq.comparator(pq.heap[pq.Left(node)], pq.heap[minNode]) < 0) {
			minNode = pq.Left(node)
		}
		if pq.Right(node) < pq.size && (pq.comparator(pq.heap[pq.Right(node)], pq.heap[minNode]) < 0) {
			minNode = pq.Right(node)
		}
		if minNode == node {
			break
		}
		pq.Swap(minNode, node)
		node = minNode
	}
}

func (pq *MyPriorityQueue) Pop() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, errors.New("pq underflow")
	}
	res := pq.heap[0]
	pq.Swap(0, pq.size-1)
	pq.size--
	pq.sink(0)
	if pq.size > 0 && pq.size == len(pq.heap)/4 {
		pq.resize(len(pq.heap) / 2)
	}

	return res, nil
}
