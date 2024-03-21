/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (49.03%)
 * Likes:    4215
 * Dislikes: 102
 * Total Accepted:    774.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
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
 * return its level order traversal as:
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
// 用队列保存待访问的节点; 用临时变量存储当前层的长度；处理完一层清空临时变量
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	// 每一层的数据
	layer := make([]int, 0)

	// 待访问的队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	// 当前层的大小
	size := len(queue)
	for len(queue) > 0 {
		if size <= 0 { // 一层访问完了
			result = append(result, layer)
			// 重置层数据
			layer = make([]int, 0)
			size = len(queue)
		}

		// 出队列
		node := queue[0]
		queue = queue[1:]
		size = size - 1
		if node == nil {
			continue
		}
		layer = append(layer, node.Val)
		// 子节点加入
		queue = append(queue, node.Left, node.Right)
	}
	return result
}

// @lc code=end

