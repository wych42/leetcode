/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
 *
 * https://leetcode.com/problems/increasing-order-search-tree/description/
 *
 * algorithms
 * Easy (74.64%)
 * Likes:    1367
 * Dislikes: 531
 * Total Accepted:    122.7K
 * Total Submissions: 164.3K
 * Testcase Example:  '[5,3,6,2,4,null,8,1,null,null,null,7,9]'
 *
 * Given the root of a binary search tree, rearrange the tree in in-order so
 * that the leftmost node in the tree is now the root of the tree, and every
 * node has no left child and only one right child.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
 * Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,1,7]
 * Output: [1,null,5,null,7]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the given tree will be in the range [1, 100].
 * 0 <= Node.val <= 1000
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
func increasingBST(root *TreeNode) *TreeNode {
	ans := &TreeNode{}
	curr := ans

	var travesal func(*TreeNode)
	travesal = func(node *TreeNode) {
		if node == nil {
			return
		}
		travesal(node.Left)
		curr.Right = node
		node.Left = nil
		curr = node
		travesal(node.Right)
	}

	travesal(root)
	return ans.Right
}

// @lc code=end

