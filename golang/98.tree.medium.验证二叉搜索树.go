package golang

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (32.63%)
 * Likes:    817
 * Dislikes: 0
 * Total Accepted:    184.4K
 * Total Submissions: 564.8K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 *
 * 假设一个二叉搜索树具有如下特征：
 *
 *
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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
func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	current := root
	var max, loop int
	for len(stack) > 0 || current != nil {
		// go to left leaf
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		node := stack[len(stack)-1]
		if loop == 0 {
			max = node.Val // most-left leaf
		} else {
			if node.Val <= max {
				return false
			}
			max = node.Val
		}
		stack = stack[:len(stack)-1]
		current = node.Right
		loop = loop + 1
	}
	return true
}

// @lc code=end
