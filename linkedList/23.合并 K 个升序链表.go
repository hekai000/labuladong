/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30005
 *
 * [23] 合并 K 个升序链表
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func mergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{Val: -1}
	p := dummy
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}

	}

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	start := 0
	end := len(lists) - 1
	return mergeKLists2(lists, start, end)

}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummyNode := &ListNode{}
	cur := dummyNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next

		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummyNode.Next
}

func mergeKLists2(lists []*ListNode, start int, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := start + (end-start)/2
	left := mergeKLists2(lists, start, mid)
	right := mergeKLists2(lists, mid+1, end)
	return mergeTwoLists(left, right)
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/

