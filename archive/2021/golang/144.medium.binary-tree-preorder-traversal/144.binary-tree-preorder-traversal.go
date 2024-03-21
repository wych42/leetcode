/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (51.74%)
 * Likes:    2042
 * Dislikes: 86
 * Total Accepted:    592.5K
 * Total Submissions: 1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given the root of a binary tree, return the preorder traversal of its nodes'
 * values.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,null,2,3]
 * Output: [1,2,3]
 *
 *
 * Example 2:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: [1]
 *
 *
 * Example 4:
 *
 *
 * Input: root = [1,2]
 * Output: [1,2]
 *
 *
 * Example 5:
 *
 *
 * Input: root = [1,null,2]
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
 *
 *
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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
二叉树的前序遍历: 中-左-右
- 访问根节点
- 前序遍历左子树
- 前序遍历右子树

也就是：取到一个节点后，将其输出，遍历左子树，然后遍历右子树。(根左右)
*/
// 非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	// 根节点入栈
	stack = append(stack, root)
	result := make([]int, 0)
	for len(stack) > 0 {
		// 直接出栈，第一次循环，取到的就是根节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		// 先压入右节点：栈先入后出
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// @lc code=end