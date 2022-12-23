package traceback

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 节点可能为负数
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return [][]int{{root.Val}}
	}

	childrenPathSums := append(pathSum(root.Left, targetSum-root.Val), pathSum(root.Right, targetSum-root.Val)...)

	var res [][]int
	for _, path := range childrenPathSums {
		res = append(res, append([]int{root.Val}, path...))
	}

	return res
}
