/*
 * @lc app=leetcode.cn id=1552 lang=golang
 * @lcpr version=30005
 *
 * [1552] 两球之间的磁力
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

/*
**
思路：

position[i]表示球的位置
m个球
最小化磁力：表示放置m个球后，position[i] - position[j]中最小值
最大化最小磁力：使得position[i] - position[j]的值尽可能大，大到一定程度后就放不下m个球了
假设答案是ans，那么小于ans的值肯定是合法的；大于ans的值就不合法
为了简化问题，可以首先对position排序，排序后左边界就是position[0], 右边界是position[len-1]
然后通过二分查找的方法找到这个ans
验证这个ans合法，则左边界右移，否则右边界左移

怎么验证这个ans呢？
可以通过模拟放置小球的方法：
第一个小球开始放在0的位置，然后遍历position找到第一个距离第一个球>ans的位置，并且放置小球，cnt数加一，然后将这个位置
作为下一次比较的左边界值，以此类推，最终能够放置cnt个球，看这cnt是否大于规定的m个球，如果大于则
说明能够放置；否则不能放置。
**
*/
func maxDistance(position []int, m int) int {

	sort.Ints(position)
	left := 1
	right := position[len(position)-1] - position[0]
	ans := -1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if check(mid, position, m) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return ans
}

func check(x int, position []int, m int) bool {
	pre, cnt := position[0], 1

	for i := 1; i < len(position); i++ {
		if position[i]-pre >= x {
			pre = position[i]
			cnt++
		}
	}
	return cnt >= m
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [5,4,3,2,1,1000000000]\n2\n
// @lcpr case=end

*/

