/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 *
 * https://leetcode.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (52.26%)
 * Likes:    1730
 * Dislikes: 166
 * Total Accepted:    237.2K
 * Total Submissions: 453.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Find the sum of all left leaves in a given binary tree.
 *
 * Example:
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * There are two left leaves in the binary tree, with values 9 and 15
 * respectively. Return 24.
 *
 *
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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var l int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // left leaft
		l = root.Left.Val
	} else {
		l = sumOfLeftLeaves(root.Left)
	}
	r := sumOfLeftLeaves(root.Right)
	return l + r
}

// @lc code=end

