package priorityqueue

import "fmt"

type SimpleMinPQ struct {
	heap []int
	size int
}

func (pq *SimpleMinPQ) parent(node int) int {
	return (node - 1) / 2
}

func (pq *SimpleMinPQ) left(node int) int {
	return 2*node + 1
}

func (pq *SimpleMinPQ) right(node int) int {
	return 2*node + 2
}

func (pq *SimpleMinPQ) swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *SimpleMinPQ) peek() int {
	return pq.heap[0]
}

func (pq *SimpleMinPQ) push(x int) {
	pq.heap[pq.size] = x
	pq.swim(pq.size)
	pq.size++
}

func (pq *SimpleMinPQ) pop() int {

	res := pq.heap[0]
	pq.heap[0] = pq.heap[pq.size-1]
	pq.size--
	pq.sink(0)
	return res
}

func (pq *SimpleMinPQ) swim(node int) {
	for node > 0 && pq.heap[pq.parent(node)] > pq.heap[node] {
		pq.swap(node, pq.parent(node))
		node = pq.parent(node)
	}
}

func (pq *SimpleMinPQ) sink(node int) {
	for pq.left(node) < pq.size || pq.right(node) < pq.size {
		min := node
		if pq.left(node) < pq.size && pq.heap[pq.left(node)] < pq.heap[min] {
			min = pq.left(node)
		}
		if pq.right(node) < pq.size && pq.heap[pq.right(node)] < pq.heap[min] {
			min = pq.right(node)
		}
		if min == node {
			break
		}
		pq.swap(node, min)
		node = min
	}
}

func main3() {
	pq := SimpleMinPQ{
		heap: make([]int, 5),
		size: 0,
	}

	pq.push(3)
	pq.push(2)
	pq.push(1)
	pq.push(5)
	pq.push(4)

	fmt.Println(pq.pop()) // 1
	fmt.Println(pq.pop()) // 2
	fmt.Println(pq.pop()) // 3
	fmt.Println(pq.pop()) // 4
	fmt.Println(pq.pop()) // 5
}
