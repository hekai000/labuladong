/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=30005
 *
 * [82] 删除排序链表中的重复元素 II
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
func deleteDuplicates2(head *ListNode) *ListNode {

	dummy1, dummy2 := &ListNode{Val: 101}, &ListNode{Val: 101}
	pDup, pUniq := dummy1, dummy2
	p := head

	for p != nil {
		if (p.Next != nil && p.Next.Val == p.Val) || p.Val == pDup.Val {
			pDup.Next = p
			pDup = pDup.Next
		} else {
			pUniq.Next = p
			pUniq = pUniq.Next
		}

		p = p.Next
		pUniq.Next = nil
		pDup.Next = nil
	}
	return dummy2.Next
}
func deleteDuplicates(head *ListNode) *ListNode {

	dummy1 := &ListNode{Val: 101}
	p, q := dummy1, head

	for q != nil {
		if q.Next != nil && q.Next.Val == q.Val {
			for q.Next != nil && q.Next.Val == q.Val {
				q = q.Next
			}
			q = q.Next
			if q == nil {
				p.Next = nil
			}

		} else {
			p.Next = q
			p = p.Next
			q = q.Next
		}

	}
	return dummy1.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/

