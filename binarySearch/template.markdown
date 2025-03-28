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
    if left < 0 || left >=len(nums) {
        return -1
    }
    if nums[left-1] == target {
        return left-1
    }
    return -1
}

```

|.... target left-bound....|

|.......right-bound target...|