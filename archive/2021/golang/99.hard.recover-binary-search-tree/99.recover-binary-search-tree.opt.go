/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 *
 * https://leetcode.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Hard (42.42%)
 * Likes:    2302
 * Dislikes: 88
 * Total Accepted:    203.7K
 * Total Submissions: 479K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * You are given the root of a binary search tree (BST), where exactly two
 * nodes of the tree were swapped by mistake. Recover the tree without changing
 * its structure.
 *
 * Follow up: A solution using O(n) space is pretty straight forward. Could you
 * devise a constant space solution?
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,3,null,null,2]
 * Output: [3,1,null,null,2]
 * Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3
 * makes the BST valid.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [3,1,4,null,null,2]
 * Output: [2,1,4,null,null,3]
 * Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2
 * and 3 makes the BST valid.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [2, 1000].
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

var prev *TreeNode

func recoverTree(root *TreeNode) {
	var pair [2]*TreeNode
	prev = nil
	traversal(root, &pair)
	pair[0].Val, pair[1].Val = pair[1].Val, pair[0].Val
}

func traversal(root *TreeNode, p *[2]*TreeNode) {
	if root == nil {
		return
	}

	traversal(root.Left, p)
	if prev != nil {
		if prev.Val > root.Val { // 搜索二叉树的中序遍历，上一个节点应该永远比当前节点小
			if (*p)[0] == nil {
				// 此时 pair 数组是空的，把 root 和 prev 都放进去
				(*p)[0] = root
				(*p)[1] = prev
			} else {
				// 此时 pair 数组里有填充值，替换当前root
				(*p)[0] = root
			}
		}
	}
	prev = root
	traversal(root.Right, p)
}

// @lc code=end

