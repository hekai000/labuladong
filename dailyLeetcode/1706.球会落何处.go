/*
 * @lc app=leetcode.cn id=1706 lang=golang
 * @lcpr version=30005
 *
 * [1706] 球会落何处
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findBall(grid [][]int) []int {

	n := len(grid[0])
	result := make([]int, n)

	for i := 0; i < n; i++ {
		col := i
		for _, row := range grid {
			dir := row[col]
			col += dir
			if col < 0 || col == n || row[col] != dir {
				col = -1
				break
			}
		}
		result[i] = col
	}
	return result

}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]\n
// @lcpr case=end

// @lcpr case=start
// [[-1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1],[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1]]\n
// @lcpr case=end

*/

