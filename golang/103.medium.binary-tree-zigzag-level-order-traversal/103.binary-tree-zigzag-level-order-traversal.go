/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (42.20%)
 * Likes:    3098
 * Dislikes: 122
 * Total Accepted:    475.1K
 * Total Submissions: 949.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the zigzag level order traversal of its nodes'
 * values. (ie, from left to right, then right to left for the next level and
 * alternate between).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its zigzag level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
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
// 在 levelOrder 的基础上，把队列换成栈，节点先入后出，入栈下一层节点时，与当前层入栈的顺序相反
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	// 待访问的栈
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	toggleOrder := false
	for len(stack) > 0 {
		// 每一层的数据
		layer := make([]int, 0)
		// 当前层的长度
		length := len(stack)
		for i := 0; i < length; i++ {
			node := stack[length-1-i]
			if node == nil {
				continue
			}
			layer = append(layer, node.Val)
			if toggleOrder {
				stack = append(stack, node.Right, node.Left)
			} else {
				stack = append(stack, node.Left, node.Right)
			}
		}
		stack = stack[length:]
		if len(layer) > 0 {
			result = append(result, layer)
		}
		toggleOrder = !toggleOrder
	}
	return result
}

// @lc code=end

