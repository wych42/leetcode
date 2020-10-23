/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	for len(stack) > 0 || root != nil {
		// push all left into stack
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		/* first loop: pop left leaf, append value, point root to parent.
		   each loop after:
			 - if node has right, visit, current node still in stack
			 - if node's right visited, visit node itself.
		*/
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == prev {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			prev = node
		} else {
			root = node.Right
		}
	}
	return result
}

// @lc code=end