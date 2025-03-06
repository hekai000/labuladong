/*
 * @lc app=leetcode.cn id=1201 lang=golang
 * @lcpr version=30008
 *
 * [1201] 丑数 III
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func nthUglyNumber(n int, a int, b int, c int) int {
	// 题目说本题结果在 [1, 2 * 10^9] 范围内，
	// 所以就按照这个范围初始化两端都闭的搜索区间
	left, right := 1, int(2e9)
	// 搜索左侧边界的二分搜索
	for left <= right {
		mid := left + (right-left)/2
		if f(mid, a, b, c) < n {
			// [1..mid] 中符合条件的元素个数不足 n，所以目标在右半边
			left = mid + 1
		} else {
			// [1..mid] 中符合条件的元素个数大于 n，所以目标在左半边
			right = mid - 1
		}
	}
	return left
}

// 计算最大公因数（辗转相除/欧几里得算法）
func gcd(a, b int) int {
	if a < b {
		// 保证 a > b
		return gcd(b, a)
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 最小公倍数
func lcm(a, b int) int {
	// 最小公倍数就是乘积除以最大公因数
	return a * b / gcd(a, b)
}

// 计算 [1..num] 之间有多少个能够被 a 或 b 或 c 整除的数字
func f(num, a, b, c int) int {
	setA, setB, setC := num/a, num/b, num/c
	setAB := num / lcm(a, b)
	setAC := num / lcm(a, c)
	setBC := num / lcm(b, c)
	setABC := num / lcm(lcm(a, b), c)
	// 集合论定理：A + B + C - A ∩ B - A ∩ C - B ∩ C + A ∩ B ∩ C
	return setA + setB + setC - setAB - setAC - setBC + setABC
}

// @lc code=end

/*
// @lcpr case=start
// 3\n2\n3\n5\n
// @lcpr case=end

// @lcpr case=start
// 4\n2\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// 5\n2\n11\n13\n
// @lcpr case=end

*/

