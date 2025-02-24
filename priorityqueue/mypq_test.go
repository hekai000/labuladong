package priorityqueue

import (
	"fmt"
	"testing"
)

func TestMyPriorityQueue(T *testing.T) {
	pq := NewMyPriorityQueue(3, func(x, y interface{}) int {
		a := x.(int)
		b := y.(int)
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	pq.Push(3)
	pq.Push(1)
	pq.Push(4)
	pq.Push(1)
	pq.Push(5)
	pq.Push(9)

	// 1 1 3 4 5 9
	for !pq.IsEmpty() {
		item, _ := pq.Pop()
		fmt.Println(item)
	}
}
