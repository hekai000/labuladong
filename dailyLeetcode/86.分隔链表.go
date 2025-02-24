/*
 * @lc app=leetcode.cn id=86 lang=golang
 * @lcpr version=30005
 *
 * [86] 分隔链表
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
func partition(head *ListNode, x int) *ListNode {
	dummyNode1 := &ListNode{} //little
	dummyNode2 := &ListNode{} //big
	cur1 := dummyNode1
	cur2 := dummyNode2
	p := head

	for p != nil {
		if p.Val < x {
			cur1.Next = p

			cur1 = cur1.Next
		} else {
			cur2.Next = p

			cur2 = cur2.Next
		}
		//断开连接
		temp := p.Next
		p.Next = nil
		p = temp

	}
	cur1.Next = dummyNode2.Next
	return dummyNode1.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,4,3,2,5,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n2\n
// @lcpr case=end

*/

