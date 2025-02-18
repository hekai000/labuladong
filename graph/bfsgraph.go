package graph

import (
	"container/list"
	"fmt"
)

// 多叉树的层序遍历
// 多叉树的层序遍历,第一种写法是不记录遍历步数的：
// func levelOrderTraverse(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	q := list.New()
// 	q.PushBack(root)
// 	for q.Len() > 0 {
// 		e := q.Front()
// 		cur := e.Value.(*Node)
// 		q.Remove(e)
// 		// 访问 cur 节点
// 		fmt.Println(cur.val)

// 		// 把 cur 的所有子节点加入队列
// 		for _, child := range cur.children {
// 			q.PushBack(child)
// 		}
// 	}
// }

// Graph structure's BFS traversal starting from node s
func bfs1(graph Graph, s int) {
	visited := make([]bool, graph.Size())
	q := list.New()
	q.PushBack(s)
	visited[s] = true

	for q.Len() > 0 {
		e := q.Front()
		cur := e.Value.(int)
		q.Remove(e)
		fmt.Println("visit", cur)
		for _, edge := range graph.Neighbors(cur) {
			if !visited[edge.to] {
				q.PushBack(edge.to)
				visited[edge.to] = true
			}
		}
	}
}

// 二种能够记录遍历步数的写法：
// 多叉树的层序遍历
// func levelOrderTraverse(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	q := []*Node{root}
// 	// 记录当前遍历到的层数（根节点视为第 1 层）
// 	depth := 1

// 	for len(q) > 0 {
// 		sz := len(q)
// 		for i := 0; i < sz; i++ {
// 			cur := q[0]
// 			q = q[1:]
// 			// 访问 cur 节点，同时知道它所在的层数
// 			fmt.Printf("depth = %d, val = %d\n", depth, cur.val)

// 			for _, child := range cur.children {
// 				q = append(q, child)
// 			}
// 		}
// 		depth++
// 	}
// }

// 从 s 开始 BFS 遍历图的所有节点，且记录遍历的步数
func bfs2(graph Graph, s int) {
	visited := make([]bool, graph.Size())
	q := []int{s}
	visited[s] = true
	// 记录从 s 开始走到当前节点的步数
	step := 0
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			fmt.Printf("visit %d at step %d\n", cur, step)
			for _, e := range graph.Neighbors(cur) {
				if !visited[e.to] {
					q = append(q, e.to)
					visited[e.to] = true
				}
			}
		}
		step++
	}
}

// 第三种能够适配不同权重边的写法：
// 多叉树的层序遍历
// 每个节点自行维护 State 类，记录深度等信息
// type State struct {
// 	node  *Node
// 	depth int
// }

// func levelOrderTraverse(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	q := []*State{{root, 1}}
// 	for len(q) > 0 {
// 		state := q[0]
// 		q = q[1:]
// 		cur := state.node
// 		depth := state.depth
// 		// 访问 cur 节点，同时知道它所在的层数
// 		fmt.Printf("depth = %d, val = %d\n", depth, cur.val)

// 		for _, child := range cur.children {
// 			q = append(q, &State{child, depth + 1})
// 		}
// 	}
// }

// 图结构的 BFS 遍历，从节点 s 开始进行 BFS，且记录路径的权重和
// 每个节点自行维护 State 类，记录从 s 走来的权重和
type State struct {
	// 当前节点 ID
	node int
	// 从起点 s 到当前节点的权重和
	weight int
}

func bfs3(graph Graph, s int) {
	visited := make([]bool, graph.Size())
	q := []*State{{s, 0}}
	visited[s] = true
	for len(q) > 0 {
		state := q[0]
		q = q[1:]
		cur := state.node
		weight := state.weight
		fmt.Printf("visit %d with path weight %d\n", cur, weight)
		for _, e := range graph.Neighbors(cur) {
			if !visited[e.to] {
				q = append(q, &State{e.to, weight + e.weight})
				visited[e.to] = true
			}
		}
	}
}
