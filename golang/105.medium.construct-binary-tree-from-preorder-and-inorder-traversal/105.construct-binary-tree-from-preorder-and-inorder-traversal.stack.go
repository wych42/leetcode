/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (41.74%)
 * Likes:    4756
 * Dislikes: 123
 * Total Accepted:    461.3K
 * Total Submissions: 890.8K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given two integer arrays preorder and inorder where preorder is the preorder
 * traversal of a binary tree and inorder is the inorder traversal of the same
 * tree, construct and return the binary tree.
 *
 *
 * Example 1:
 *
 *
 * Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * Output: [3,9,20,null,null,15,7]
 *
 *
 * Example 2:
 *
 *
 * Input: preorder = [-1], inorder = [-1]
 * Output: [-1]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= preorder.length <= 3000
 * inorder.length == preorder.length
 * -3000 <= preorder[i], inorder[i] <= 3000
 * preorder and inorder consist of unique values.
 * Each value of inorder also appears in preorder.
 * preorder is guaranteed to be the preorder traversal of the tree.
 * inorder is guaranteed to be the inorder traversal of the tree.
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	preIdx, inIdx := 0, 0
	stack := make([]*TreeNode, 0)
	root := &TreeNode{Val: preorder[0]}
	preIdx++
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top.Val != inorder[inIdx] {
			node := &TreeNode{Val: preorder[preIdx]}
			preIdx++
			top.Left = node
			stack = append(stack, node)

		} else {
			stack = stack[:len(stack)-1]
			inIdx++

			if inIdx == len(inorder) {
				break
			}

			if len(stack) > 0 && stack[len(stack)-1].Val == inorder[inIdx] {
				continue
			}

			node := &TreeNode{Val: preorder[preIdx]}
			preIdx++
			top.Right = node
			stack = append(stack, node)
		}

	}
	return root
}

// @lc code=end

