package graph

import "fmt"

type WeightedUndigraph struct {
	graph *WeightedDigraph
}

// 新建一个带权重的无向图，传入节点数量
func NewWeightedUndigraph(n int) *WeightedUndigraph {
	return &WeightedUndigraph{graph: NewWeightedDigraph(n)}
}

// 增，添加一条带权重的无向边
func (wg *WeightedUndigraph) addEdge(from, to, weight int) {
	wg.graph.AddEdge(from, to, weight)
	wg.graph.AddEdge(to, from, weight)
}

// 删，删除一条无向边
func (wg *WeightedUndigraph) removeEdge(from, to int) {
	wg.graph.RemoveEdge(from, to)
	wg.graph.RemoveEdge(to, from)
}

// 查，判断两个节点是否相邻
func (wg *WeightedUndigraph) hasEdge(from, to int) bool {
	return wg.graph.HasEdge(from, to)
}

// 查，返回一条边的权重
func (wg *WeightedUndigraph) weight(from, to int) int {
	return wg.graph.Weight(from, to)
}

// 查，返回某个节点的所有邻居节点
func (wg *WeightedUndigraph) neighbors(v int) []Edge {
	return wg.graph.Neighbors(v)
}

func main3() {
	graph := NewWeightedUndigraph(3)
	graph.addEdge(0, 1, 1)
	graph.addEdge(1, 2, 2)
	graph.addEdge(2, 0, 3)
	graph.addEdge(2, 1, 4)

	fmt.Println(graph.hasEdge(0, 1)) // true
	fmt.Println(graph.hasEdge(1, 0)) // true

	for _, edge := range graph.neighbors(2) {
		fmt.Printf("2 <-> %d, wight: %d\n", edge.to, edge.weight)
	}
	// 2 <-> 0, wight: 3
	// 2 <-> 1, wight: 4

	graph.removeEdge(0, 1)
	fmt.Println(graph.hasEdge(0, 1)) // false
	fmt.Println(graph.hasEdge(1, 0)) // false
}
