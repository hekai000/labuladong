package tree

import "fmt"

type State struct {
	node  *TreeNode
	depth int
}

func levelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			fmt.Printf("depth = %d, val = %d\n", depth, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++

	}
}

func levelOrderTraverse3(root *TreeNode) {
	if root == nil {
		return
	}
	q := []State{{root, 1}}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		// 访问 cur 节点，同时知道它的路径权重和
		fmt.Printf("depth = %d, val = %d\n", cur.depth, cur.node.Val)

		if cur.node.Left != nil {
			q = append(q, State{cur.node.Left, cur.depth + 1})
		}
		if cur.node.Right != nil {
			q = append(q, State{cur.node.Right, cur.depth + 1})
		}

	}
}

//最短距离： BFS
//所有路径：DFS

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return depth
}

//所有路径

func allPaths(root *TreeNode) [][]int {
	result := [][]int{}

	path := []int{}

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			result = append(result, path)
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return result

}

// 基础写法
func levelOrderTraverseMTree1(root *MTreeNode) {
	if root == nil {
		return
	}
	q := []*MTreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Println(cur.val)

		q = append(q, cur.children...)
	}
}

// 最常见的写法，带depth
func levelOrderTraverseMTree2(root *MTreeNode) {
	if root == nil {
		return
	}
	q := []*MTreeNode{}
	depth := 1
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			// 访问 cur 节点，同时知道它所在的层数
			fmt.Printf("depth = %d, val = %d\n", depth, cur.val)
			q = append(q, cur.children...)
		}
		depth++
	}

}

type MState struct {
	node  *MTreeNode
	depth int
}

// 能够适配不同权重边的写法
func levelOrderTraverseMTree3(root *MTreeNode) {
	if root == nil {
		return
	}
	q := []MState{}
	q = append(q, MState{root, 1})
	for len(q) > 0 {
		state := q[0]
		q = q[1:]
		cur := state.node
		depth := state.depth
		// 访问 cur 节点，同时知道它所在的层数
		fmt.Printf("depth = %d, val = %d\n", depth, cur.val)
		for _, child := range cur.children {
			q = append(q, MState{child, depth + 1})
		}
	}
}
