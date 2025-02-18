package graph

func allPathsSourceTarget(graph [][]int) [][]int {
	// 记录所有路径
	var res [][]int
	var path []int
	traversegraph(graph, 0, &path, &res)
	return res
}

// 图的遍历框架
func traversegraph(graph [][]int, s int, path *[]int, res *[][]int) {
	// 添加节点 s 到路径
	*path = append(*path, s)

	n := len(graph)
	if s == n-1 {
		// 到达终点
		*res = append(*res, append([]int(nil), *path...))
		*path = (*path)[:len(*path)-1]
		return
	}

	// 递归每个相邻节点
	for _, v := range graph[s] {
		traversegraph(graph, v, path, res)
	}

	// 从路径移出节点 s
	*path = (*path)[:len(*path)-1]
}
