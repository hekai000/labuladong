#单调栈模板

## 代码

```golang
func calculateGreaterElement(nums []int) []int {
    n := len(nums)
    // 存放答案的数组
    res := make([]int, n)
    s := make([]int, 0)
    // 倒着往栈里放
    for i := n - 1; i >= 0; i-- {
        // 判定个子高矮
        for len(s) != 0 && s[len(s) - 1] <= nums[i] {
            // 矮个起开，反正也被挡着了。。。
            s = s[:len(s)-1]
        }
        // nums[i] 身后的更大元素
        if len(s) == 0 {
            res[i] = -1
        } else {
            res[i] = s[len(s) - 1]
        }
        s = append(s, nums[i])
    }
    return res
}
```