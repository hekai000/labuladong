/*
 * @lc app=leetcode.cn id=313 lang=golang
 * @lcpr version=30008
 *
 * [313] 超级丑数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

type node struct {
	product int
	prime   int
	index   int
}
type PriorityQueue []node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].product < pq[j].product
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func nthSuperUglyNumber(n int, primes []int) int {
	pq := make(PriorityQueue, 0)
	for i := 0; i < len(primes); i++ {
		heap.Push(&pq, node{1, primes[i], 1})
	}
	ugly := make([]int, n+1)
	p := 1

	for p <= n {
		pair := heap.Pop(&pq).(node)
		product := pair.product
		prime := pair.prime
		index := pair.index

		if product != ugly[p-1] {
			ugly[p] = product
			p++
		}
		heap.Push(&pq, node{ugly[index] * prime, prime, index + 1})

	}
	return ugly[n]
}

// @lc code=end

/*
// @lcpr case=start
// 12\n[2,7,13,19]\n
// @lcpr case=end

// @lcpr case=start
// 1\n[2,3,5]\n
// @lcpr case=end

*/

