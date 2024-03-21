/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (25.90%)
 * Likes:    5423
 * Dislikes: 645
 * Total Accepted:    898.2K
 * Total Submissions: 3.1M
 * Testcase Example:  '[2,1,3]'
 *
 * Given the root of a binary tree, determine if it is a valid binary search
 * tree (BST).
 *
 * A valid BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is
 * 4.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * -2^31 <= Node.val <= 2^31 - 1
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
/*
 左子树都要比根小；右子树都要比根大。
 左右子树都要是搜索树。
 一个后续遍历:先看左边，再看右边
 需要一个临时变量存储最大值
*/
func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	current := root
	var max, loop int
	for len(stack) > 0 || current != nil {
		// 左侧节点入栈
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		node := stack[len(stack)-1]
		if loop == 0 {
			max = node.Val // 最左叶子节点出栈，这是当前的最大值，但是后续所有的值都要大于它
		} else {
			if node.Val <= max {
				return false
			}
			max = node.Val // 最大值更新为当前节点
		}
		stack = stack[:len(stack)-1] // 已经比较过的节点可以去除
		current = node.Right         // 遍历右子树
		loop = loop + 1
	}
	return true

}

// @lc code=end

