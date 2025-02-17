package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	traverse(root.Right)
}

type MTreeNode struct {
	val      int
	children []*MTreeNode
}

func traverseMTree(root *MTreeNode) {
	if root == nil {
		return
	}
	//前序位置
	for _, child := range root.children {
		traverseMTree(child)
	}
	//后序位置
}
