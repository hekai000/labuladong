package graph

type Vertex struct {
	id        int
	neighbors []*Vertex
}

// 邻接表
var graph [][]int

// 邻接矩阵
var matrix [][]bool

// type Edge struct {
// 	to     int
// 	weight int
// }

// 邻接表
// graph[x] 存储 x 的所有邻居节点以及对应的权重
// 具体实现不一定非得这样，可以参考后面的通用实现
var graph1 [][]Edge

// 邻接矩阵
// matrix[x][y] 记录 x 指向 y 的边的权重，0 表示不相邻
var matrix1 [][]int

type Graph interface {
	// 添加一条边（带权重）
	AddEdge(from int, to int, weight int)

	// 删除一条边
	RemoveEdge(from int, to int)

	// 判断两个节点是否相邻
	HasEdge(from int, to int) bool

	// 返回一条边的权重
	Weight(from int, to int) int

	// 返回某个节点的所有邻居节点和对应权重
	Neighbors(v int) []Edge

	// 返回节点总数
	Size() int
}
