/*
 * @lc app=leetcode.cn id=1287 lang=golang
 * @lcpr version=30005
 *
 * [1287] 有序数组中出现次数超过25%的元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// func findSpecialInteger(arr []int) int {
// 	cur := arr[0]
// 	count := 0
// 	length := len(arr)
// 	for i := 0; i < length; i++ {
// 		if arr[i] == cur {
// 			count++
// 			if count*4 > length {
// 				return cur
// 			}
// 		} else {
// 			cur = arr[i]
// 			count = 1
// 		}
// 	}
// 	return -1
// }

func findSpecialInteger(arr []int) int {
	//假设arr[0], arr[span],arr[2*span]..不含x，则x的次数最多为span-1-1+1=span-1次，与假设不符
	//故arr[0], arr[span],arr[2*span]肯定包含x
	n := len(arr)
	span := n/4 + 1
	for i := 0; i < n; i += span {
		start := binarySearch(arr, arr[i]) //>=arr[i]的位置，也就是arr[i]的最左边界
		end := binarySearch(arr, arr[i]+1) //表示>=arr[i]+1的位置，也就是end-1是arr[i]最右边界
		if end-start >= span {             //区间个数，最右边界-最左边界+1=end -1 -start + 1=end-start
			return arr[i]
		}
	}
	return -1
}

func binarySearch(arr []int, target int) int {
	length := len(arr)
	left := 0
	right := length - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	//闭区间最终形态是...RL...，R+1=L,L是等于target的最左边界
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,6,6,6,6,7,10]\n
// @lcpr case=end

*/

