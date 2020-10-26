/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (63.57%)
 * Likes:    674
 * Dislikes: 0
 * Total Accepted:    209.2K
 * Total Submissions: 328.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
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
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	size := len(queue)
	layer := make([]int, 0)

	for len(queue) > 0 {
		if size <= 0 {
			result = append(result, layer)
			layer = make([]int, 0)
			size = len(queue)
		}

		node := queue[0]
		queue = queue[1:]
		size = size - 1
		if node == nil {
			continue
		}
		layer = append(layer, node.Val)
		queue = append(queue, node.Left, node.Right)

	}
	return result
}

// @lc code=end

