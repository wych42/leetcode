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
func recoverTree(root *TreeNode) {
	nodes := make([]*TreeNode, 0)
	traversal(root, &nodes)
	if len(nodes) == 0 {
		return
	}
	for _, node := range nodes {
		fmt.Println(node.Val)
	}
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].Val > nodes[j].Val {
				nodes[i].Val, nodes[j].Val = nodes[j.Val], nodes[i].Val
			}
		}
	}
}

func traversal(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	traversal(root.Left, nodes)
	*nodes = append(*nodes, root)
	traversal(root.Right, nodes)
}

// @lc code=end

