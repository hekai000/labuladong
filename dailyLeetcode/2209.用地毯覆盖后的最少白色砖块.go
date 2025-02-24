/*
 * @lc app=leetcode.cn id=2209 lang=golang
 * @lcpr version=30005
 *
 * [2209] 用地毯覆盖后的最少白色砖块
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	//d[i][j] 表示前i个砖块使用j条地毯最少的白色砖块
	//d[0][j] j[0,numCarpets]为0，
	//d[i][0] d[i-1][0]+floor[i]
	const INF = 0x3f3f3f3f
	n := len(floor)
	d := make([][]int, n+1)

	for i := range d {
		d[i] = make([]int, numCarpets+1)
		for j := range d[i] {
			d[i][j] = INF
		}
	}

	for j := 0; j <= numCarpets; j++ {
		d[0][j] = 0
	}

	for i := 1; i <= n; i++ {
		d[i][0] = d[i-1][0]
		if floor[i-1] == '1' {
			d[i][0]++
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= numCarpets; j++ {
			d[i][j] = d[i-1][j]
			if floor[i-1] == '1' {
				d[i][j]++
			}
			d[i][j] = min(d[i][j], d[max(0, i-carpetLen)][j-1])
		}
	}
	return d[n][numCarpets]

}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

/*
// @lcpr case=start
// "10110101"\n2\n2\n
// @lcpr case=end

// @lcpr case=start
// "11111"\n2\n3\n
// @lcpr case=end

*/

