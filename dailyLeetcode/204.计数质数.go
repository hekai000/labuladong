/*
 * @lc app=leetcode.cn id=204 lang=golang
 * @lcpr version=30008
 *
 * [204] 计数质数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}

	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count

}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

