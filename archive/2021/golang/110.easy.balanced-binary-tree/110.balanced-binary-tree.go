/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (44.68%)
 * Likes:    3211
 * Dislikes: 211
 * Total Accepted:    542.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the left and right subtrees of every node differ in
 * height by no more than 1.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,2,2,3,3,null,null,4,4]
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: root = []
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 5000].
 * -10^4 <= Node.val <= 10^4
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
func isBalanced(root *TreeNode) bool {
	_, balanced := height(root)
	return balanced
}

func height(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	if root.Left == nil && root.Right == nil {
		return 1, true
	}
	var left, right int
	var balanced bool
	left, balanced = height(root.Left)
	if !balanced {
		return 0, false
	}
	right, balanced = height(root.Right)
	if !balanced {
		return 0, false
	}
	if left > right {
		return left + 1, left-right <= 1
	}
	return right + 1, right-left <= 1
}

// @lc code=end

