package graph

import "fmt"

func traverse(graph WeightedDigraph, s int, visited []bool) {
	if s < 0 || s >= len(graph.graph) {
		return
	}
	if visited[s] {
		return
	}
	visited[s] = true
	fmt.Println("visit", s)
	for _, e := range graph.Neighbors(s) {
		traverse(graph, e.to, visited)
	}
}

// 下面的算法代码可以遍历图的所有路径，寻找从 src 到 dest 的所有路径

// onPath 和 path 记录当前递归路径上的节点

var onPath []bool = make([]bool, graph.Size())
var path2 []int

func traverse2(graph Graph, src int, dest int) {
	// base case
	if src < 0 || src >= graph.Size() {
		return
	}
	if onPath[src] {
		// 防止死循环（成环）
		return
	}
	// 前序位置
	onPath[src] = true
	path2 = append(path2, src)
	if src == dest {
		fmt.Println("find path:", path2)
	}
	for _, e := range graph.Neighbors(src) {
		traverse2(graph, e.to, dest)
	}
	// 后序位置
	path2 = path2[:len(path2)-1]
	onPath[src] = false
}
