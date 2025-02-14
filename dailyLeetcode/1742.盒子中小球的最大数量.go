/*
 * @lc app=leetcode.cn id=1742 lang=golang
 * @lcpr version=30004
 *
 * [1742] 盒子中小球的最大数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countBalls(lowLimit int, highLimit int) int {
	result := map[int]int{}
	maxres := 0
	for i := lowLimit; i <= highLimit; i++ {
		index := 0
		for j := i; j > 0; j /= 10 {
			index += j % 10
		}
		result[index]++
		if maxres < result[index] {
			maxres = result[index]
		}

	}
	return maxres

}

// @lc code=end

/*
// @lcpr case=start
// 1\n10\n
// @lcpr case=end

// @lcpr case=start
// 5\n15\n
// @lcpr case=end

// @lcpr case=start
// 19\n28\n
// @lcpr case=end

*/

