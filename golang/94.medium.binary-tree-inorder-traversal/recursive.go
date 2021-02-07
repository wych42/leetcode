/*
 * @lc app=leetcode id=94 lang=golang
 */

// @lc code=start
var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	inorderTraversalRecursive(root)
	return result
}

func inorderTraversalRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalRecursive(root.Left)
	result = append(result, root.Val)
	inorderTraversalRecursive(root.Right)
}

// @lc code=end