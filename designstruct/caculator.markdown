```
func calculate(s string) int {
	// key 是左括号的索引，value 是对应的右括号的索引
	rightIndex := make(map[int]int)
	// 利用栈结构来找到对应的括号
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rightIndex[pop] = i
		}
	}
	return _calculate(s, 0, len(s)-1, rightIndex)
}

// 定义：返回 s[start..end] 内的表达式的计算结果
func _calculate(s string, start int, end int, rightIndex map[int]int) int {
	// 需要把字符串转成双端队列方便操作
	stk := []int{}
	// 记录算式中的数字
	num := 0
	// 记录 num 前的符号，初始化为 +
	// 修改为 byte 类型
	sign := byte('+')
	for i := start; i <= end; i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = 10*num + int(c-'0')
		}
		if c == '(' {
			// 递归计算括号内的表达式
			num = _calculate(s, i+1, rightIndex[i]-1, rightIndex)
			i = rightIndex[i]
		}
		if c == '+' || c == '-' || c == '*' || c == '/' || i == end {
			var pre int
			switch sign {
			case '+':
				stk = append(stk, num)
			case '-':
				stk = append(stk, -num)
			// 只要拿出前一个数字做对应运算即可
			case '*':
				pre = stk[len(stk)-1]
				stk[len(stk)-1] = pre * num
			case '/':
				pre = stk[len(stk)-1]
				stk[len(stk)-1] = pre / num
			}
			// 更新符号为当前符号，数字清零
			sign = c
			num = 0
		}
	}
	// 将栈中所有结果求和就是答案
	res := 0
	for len(stk) != 0 {
		res += stk[len(stk)-1]
		stk = stk[:len(stk)-1]
	}
	return res
}
```