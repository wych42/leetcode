/*
 * @lc app=leetcode id=144 lang=golang
 */
// @lc code=start
var result []int

func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	preorderTraversalRecursive(root)
	return result
}

func preorderTraversalRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	preorderTraversalRecursive(root.Left)
	preorderTraversalRecursive(root.Right)
}

// @lc code=end