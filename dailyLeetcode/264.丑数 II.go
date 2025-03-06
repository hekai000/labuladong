/*
 * @lc app=leetcode.cn id=264 lang=golang
 * @lcpr version=30008
 *
 * [264] 丑数 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	product2, product3, product5 := 1, 1, 1

	ugly := make([]int, n+1)
	p := 1

	for p <= n {
		minP := min(min(product2, product3), product5)
		ugly[p] = minP
		p++
		if minP == product2 {
			product2 = 2 * ugly[p2]
			p2++
		}
		if minP == product3 {
			product3 = 3 * ugly[p3]
			p3++
		}
		if minP == product5 {
			product5 = 5 * ugly[p5]
			p5++
		}
	}
	return ugly[n]

}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

