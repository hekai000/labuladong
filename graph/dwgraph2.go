package graph

import "fmt"

// 加权有向图的通用实现（邻接矩阵）
type Edge2 struct {
	To     int
	Weight int
}

type WeightedDigraph2 struct {
	// 邻接矩阵，matrix[from][to] 存储从节点 from 到节点 to 的边的权重
	// 0 表示没有连接
	matrix [][]int
}

func NewWeightedDigraph2(n int) *WeightedDigraph2 {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return &WeightedDigraph2{matrix}
}

// 增，添加一条带权重的有向边，复杂度 O(1)
func (g *WeightedDigraph2) AddEdge(from, to, weight int) {
	g.matrix[from][to] = weight
}

// 删，删除一条有向边，复杂度 O(1)
func (g *WeightedDigraph2) RemoveEdge(from, to int) {
	g.matrix[from][to] = 0
}

// 查，判断两个节点是否相邻，复杂度 O(1)
func (g *WeightedDigraph2) HasEdge(from, to int) bool {
	return g.matrix[from][to] != 0
}

// 查，返回一条边的权重，复杂度 O(1)
func (g *WeightedDigraph2) Weight(from, to int) int {
	return g.matrix[from][to]
}

// 查，返回某个节点的所有邻居节点，复杂度 O(V)
func (g *WeightedDigraph2) Neighbors(v int) []Edge2 {
	res := []Edge2{}
	for i := 0; i < len(g.matrix[v]); i++ {
		if g.matrix[v][i] > 0 {
			res = append(res, Edge2{To: i, Weight: g.matrix[v][i]})
		}
	}
	return res
}

func main2() {
	graph := NewWeightedDigraph2(3)
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(2, 0, 3)
	graph.AddEdge(2, 1, 4)

	fmt.Println(graph.HasEdge(0, 1)) // true
	fmt.Println(graph.HasEdge(1, 0)) // false

	for _, Edge2 := range graph.Neighbors(2) {
		fmt.Printf("%d -> %d, weight: %d\n", 2, Edge2.To, Edge2.Weight)
	}
	// 2 -> 0, weight: 3
	// 2 -> 1, weight: 4

	graph.RemoveEdge(0, 1)
	fmt.Println(graph.HasEdge(0, 1)) // false
}
