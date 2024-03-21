/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (49.03%)
 * Likes:    2354
 * Dislikes: 110
 * Total Accepted:    453.7K
 * Total Submissions: 791K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given the root of a binary tree, return the postorder traversal of its
 * nodes' values.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,null,2,3]
 * Output: [3,2,1]
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
 * Output: [2,1]
 *
 *
 * Example 5:
 *
 *
 * Input: root = [1,null,2]
 * Output: [2,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of the nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * Follow up:
 *
 * Recursive solution is trivial, could you do it iteratively?
 *
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
二叉树的中序遍历: 左-右-根
- 后序遍历左子树
- 后序遍历右子树
- 访问根节点

也就是：取到一个节点后，将其暂存，遍历完左子树后，遍历右子树，再输出该节点的值。(左右根)
*/
// 非递归: 在输出根节点的值时，需要判断左右子树已经遍历完成
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	for len(stack) > 0 || root != nil {
		// 把最左侧子树入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 从最左叶子节点开始访问
		node := stack[len(stack)-1]

		// 如果没有右节点，或者右节点已经访问过，当前节点出栈
		// prev 指向当前节点
		if node.Right == nil || node.Right == prev {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			prev = node
		} else {
			// 如果有右节点且没有访问过, root 指向右节点，开始后续遍历
			root = node.Right
		}
	}
	return result
}

// @lc code=end

