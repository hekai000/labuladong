#二分搜索模板


```
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target{
            ...
        }else if nums[mid] < target {
            left = ...
        }else if nums[mid] > target {
            right = ...
        }
    }
    return ...
}
```

搜索左侧边界

如果存在target，返回的是target的左侧边界
如果不存在target，返回的是大于target的最小索引
```

func left_bound(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            right = mid - 1
        }else if nums[mid] > target {
            right = mid - 1
        }else if nums[mid] < target {
            left = mid +1
        }
    }
    if left < 0 || left >=len(nums) {
        return -1
    }
    if nums[left] == target {
        return left
    }
    return -1
}
```


搜索右侧边界
如果存在target，返回的是target的右侧边界
如果不存在target，返回的是小于target的最大索引

```
func right_bound(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            left = mid + 1
        }else if nums[mid] > target {
            right = mid - 1
        }else if nums[mid] < target {
            left = mid +1
        }
    }
    <!-- if left - 1 < 0 || left - 1 >=len(nums) {
        return -1
    }
    if nums[left-1] == target {
        return left-1
    } -->
    if right < 0 || right >=len(nums) {
        return -1
    }
    if nums[right] == target {
        return right
    }
    return -1
}

```

|.... target left-bound....|

|.......right-bound target...|


泛化二分查找：

x, f(x), target

x是自变量
f(x)单调函数
target是目标值
```
// 函数 f 是关于自变量 x 的单调函数
func f(x int) int {
    // ...
}

// 主函数，在 f(x) == target 的约束下求 x 的最值
func solution(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    // 问自己：自变量 x 的最小值是多少？
    left := ...
    // 问自己：自变量 x 的最大值是多少？
    right := ... + 1
    
    for left < right {
        mid := left + (right - left) / 2
        if f(mid) == target {
            // 问自己：题目是求左边界还是右边界？
            // ...
        } else if f(mid) < target {
            // 问自己：怎么让 f(x) 大一点？
            // ...
        } else if f(mid) > target {
            // 问自己：怎么让 f(x) 小一点？
            // ...
        }
    }
    return left
}
```