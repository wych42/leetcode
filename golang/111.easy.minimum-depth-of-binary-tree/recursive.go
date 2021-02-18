/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (35.48%)
 * Likes:    2155
 * Dislikes: 795
 * Total Accepted:    523.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: root = [2,null,3,null,4,null,5,null,6]
 * Output: 5
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 10^5].
 * -1000 <= Node.val <= 1000
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
// 递归左右子树
import "math"

func minDepth(root *TreeNode) int {
	return recursiveMinDepth(root, 1, math.MaxInt32)
}

func recursiveMinDepth(root *TreeNode, depth int, max int) int {
	if root == nil {
		return depth - 1
	}
	if depth == max {
		return depth
	}
	left := recursiveMinDepth(root.Left, depth+1, max)
	if left != depth && left < max {
		max = left
	}
	right := recursiveMinDepth(root.Right, depth+1, max)
	if left == depth { // 左子树是空
		return right
	}
	if right == depth { // 右子树是空
		return left
	}
	if left >= right {
		return right
	}
	return left
}

// @lc code=end

