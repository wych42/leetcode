/*
 * @lc app=leetcode id=145 lang=golang
 */

// 递归
// @lc code=start
var result []int

func postorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	postorderTraversalRecursive(root)
	return result
}

func postorderTraversalRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversalRecursive(root.Left)
	postorderTraversalRecursive(root.Right)
	result = append(result, root.Val)
}

// @lc code=end

