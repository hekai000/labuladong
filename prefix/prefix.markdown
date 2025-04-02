
前缀和技巧：主要适用的场景是原始数组不会被修改的情况下，频繁查询某个区间的累加和


核心代码：
type PrefixSum struct {
    // 前缀和数组
    prefix []int
}

// 输入一个数组，构造前缀和
func Constructor(nums []int) *PrefixSum {
    prefix := make([]int, len(nums)+1)
    // 计算 nums 的累加和
    for i := 1; i < len(prefix); i++ {
        prefix[i] = prefix[i-1] + nums[i-1]
    }
    return &PrefixSum{prefix}
}

// 查询闭区间 [i, j] 的累加和
func (ps *PrefixSum) Query(i, j int) int {
    return ps.prefix[j+1] - ps.prefix[i]
}

差分数组的主要适用场景是频繁对原始数组的某个区间的元素进行增减。

// 差分数组工具类
type Difference struct {
    // 差分数组
    diff []int
}

// 输入一个初始数组，区间操作将在这个数组上进行
func NewDifference(nums []int) *Difference {
    diff := make([]int, len(nums))
    // 根据初始数组构造差分数组
    diff[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        diff[i] = nums[i] - nums[i-1]
    }
    return &Difference{diff: diff}
}

// 给闭区间 [i, j] 增加 val（可以是负数）
func (d *Difference) Increment(i, j, val int) {
    d.diff[i] += val
    if j+1 < len(d.diff) {
        d.diff[j+1] -= val
    }
}

// 返回结果数组
func (d *Difference) Result() []int {
    res := make([]int, len(d.diff))
    // 根据差分数组构造结果数组
    res[0] = d.diff[0]
    for i := 1; i < len(d.diff); i++ {
        res[i] = res[i-1] + d.diff[i]
    }
    return res
}