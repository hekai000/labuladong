/*
 * @lc app=leetcode.cn id=2080 lang=golang
 * @lcpr version=30005
 *
 * [2080] 区间内查询数字的频率
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type RangeFreqQuery struct {
	freqmap map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	freqmap := map[int][]int{}
	for index, value := range arr {
		freqmap[value] = append(freqmap[value], index)
	}

	return RangeFreqQuery{
		freqmap: freqmap,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {

	if _, ok := this.freqmap[value]; !ok {
		return 0
	}
	l := lowerBounds(this.freqmap[value], left)
	r := lowerBounds(this.freqmap[value], right+1)
	return r - l
}

func lowerBounds(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left

}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
// @lc code=end

/*
// @lcpr case=start
// ["RangeFreqQuery", "query", "query"][[[12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56]], [1, 2, 4], [0, 11, 33]]\n
// @lcpr case=end

*/

