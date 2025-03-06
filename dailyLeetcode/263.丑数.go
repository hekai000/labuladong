/*
 * @lc app=leetcode.cn id=263 lang=golang
 * @lcpr version=30008
 *
 * [263] 丑数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

// @lc code=end

/*
// @lcpr case=start
// 6\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 14\n
// @lcpr case=end

*/

