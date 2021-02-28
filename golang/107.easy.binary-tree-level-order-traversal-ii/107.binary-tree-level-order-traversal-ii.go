/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (55.14%)
 * Likes:    2023
 * Dislikes: 237
 * Total Accepted:    408.6K
 * Total Submissions: 740.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given the root of a binary tree, return the bottom-up level order traversal
 * of its nodes' values. (i.e., from left to right, level by level from leaf to
 * root).
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: [[15,7],[9,20],[3]]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1]
 * Output: [[1]]
 *
 *
 * Example 3:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 2000].
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
func levelOrderBottom(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	size := len(queue)

	result := make([][]int, 0)
	level := make([]int, 0)
	for len(queue) > 0 {
		if size <= 0 {
			result = append(result, level)
			size = len(queue)
			level = make([]int, 0)
		}
		node := queue[0]
		queue = queue[1:]
		size--
		if node == nil {
			continue
		}
		level = append(level, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	l := len(result)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[l-i-1] = result[l-i-1], result[i]
	}
	return result
}

// @lc code=end

