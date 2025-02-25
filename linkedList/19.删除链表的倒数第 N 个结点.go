/*
 * @lc app=leetcode.cn id=19 lang=golang
 * @lcpr version=30005
 *
 * [19] 删除链表的倒数第 N 个结点
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

func findNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	p2 := head

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	x := findNthFromEnd(dummyNode, n+1)
	x.Next = x.Next.Next

	return dummyNode.Next

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

*/

