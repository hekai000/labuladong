package graph

import (
	"fmt"
)

// 加权有向图的通用实现（邻接表）
type WeightedDigraph struct {
	graph [][]Edge
}

// 存储相邻节点及边的权重
type Edge struct {
	to     int
	weight int
}

func NewWeightedDigraph(n int) *WeightedDigraph {
	// 我们这里简单起见，建图时要传入节点总数，这其实可以优化
	// 比如把 graph 设置为 map[int][]Edge，就可以动态添加新节点了
	graph := make([][]Edge, n)
	return &WeightedDigraph{graph: graph}
}

// 增，添加一条带权重的有向边，复杂度 O(1)
func (g *WeightedDigraph) AddEdge(from, to, weight int) {
	g.graph[from] = append(g.graph[from], Edge{to: to, weight: weight})
}

// 删，删除一条有向边，复杂度 O(V)
func (g *WeightedDigraph) RemoveEdge(from, to int) {
	for i, e := range g.graph[from] {
		if e.to == to {
			g.graph[from] = append(g.graph[from][:i], g.graph[from][i+1:]...)
			break
		}
	}
}

// 查，判断两个节点是否相邻，复杂度 O(V)
func (g *WeightedDigraph) HasEdge(from, to int) bool {
	for _, e := range g.graph[from] {
		if e.to == to {
			return true
		}
	}
	return false
}

// 查，返回一条边的权重，复杂度 O(V)
func (g *WeightedDigraph) Weight(from, to int) int {
	for _, e := range g.graph[from] {
		if e.to == to {
			return e.weight
		}
	}
	panic("No such edge")
}

// 上面的 HasEdge、RemoveEdge、Weight 方法遍历 List 的行为是可以优化的
// 比如用 map[int]map[int]int 存储邻接表
// 这样就可以避免遍历 List，复杂度就能降到 O(1)

// 查，返回某个节点的所有邻居节点，复杂度 O(1)
func (g *WeightedDigraph) Neighbors(v int) []Edge {
	return g.graph[v]
}

func main() {
	graph := NewWeightedDigraph(3)
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(2, 0, 3)
	graph.AddEdge(2, 1, 4)

	fmt.Println(graph.HasEdge(0, 1)) // true
	fmt.Println(graph.HasEdge(1, 0)) // false

	for _, edge := range graph.Neighbors(2) {
		fmt.Printf("%d -> %d, weight: %d\n", 2, edge.to, edge.weight)
	}
	// 2 -> 0, weight: 3
	// 2 -> 1, weight: 4

	graph.RemoveEdge(0, 1)
	fmt.Println(graph.HasEdge(0, 1)) // false
}
