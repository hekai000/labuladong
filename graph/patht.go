package graph

import "fmt"

type Node struct {
	val      int
	children []*Node
}

// 多叉树的遍历框架，寻找从根节点到目标节点的路径
var path []*Node

func mttraverse(root *Node, targetNode *Node) {
	// base case
	if root == nil {
		return
	}
	// 前序位置
	path = append(path, root)
	if root.val == targetNode.val {
		fmt.Printf("find path: ")
		for _, node := range path {
			fmt.Printf("%d ", node.val)
		}
		fmt.Println()
	}
	for _, child := range root.children {
		mttraverse(child, targetNode)
	}
	// 后序位置
	path = path[:len(path)-1]
}
