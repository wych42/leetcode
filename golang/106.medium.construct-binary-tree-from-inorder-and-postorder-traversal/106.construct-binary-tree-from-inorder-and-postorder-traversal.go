/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (49.56%)
 * Likes:    2455
 * Dislikes: 47
 * Total Accepted:    282.4K
 * Total Submissions: 568.9K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given two integer arrays inorder and postorder where inorder is the inorder
 * traversal of a binary tree and postorder is the postorder traversal of the
 * same tree, construct and return the binary tree.
 *
 *
 * Example 1:
 *
 *
 * Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * Output: [3,9,20,null,null,15,7]
 *
 *
 * Example 2:
 *
 *
 * Input: inorder = [-1], postorder = [-1]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= inorder.length <= 3000
 * postorder.length == inorder.length
 * -3000 <= inorder[i], postorder[i] <= 3000
 * inorder and postorder consist of unique values.
 * Each value of postorder also appears in inorder.
 * inorder is guaranteed to be the inorder traversal of the tree.
 * postorder is guaranteed to be the postorder traversal of the tree.
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	/*
	 inorder: left-root-right
	 postorder: left-right-root

	 那么找到跟节点: postorder 的最后一个元素
	 在 inorder 中找到这个根节点 idx，它的左边就是左子树；右边就是右子树
	 同时, postorder 中, [0:idx] 也是左子树, [idx:len-1] 就是右子树
	*/
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(postorder) == 1 && len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	inIdx := 0
	for ; inIdx < len(inorder) && inorder[inIdx] != root.Val; inIdx++ {
	}
	if inIdx > 0 {
		// inIdx == 0, 说明没有左子树
		root.Left = buildTree(inorder[0:inIdx], postorder[0:inIdx])
	}
	if inIdx < len(inorder)-1 {
		//inIdx == len(inorder) - 1, 说明没有左子树
		root.Right = buildTree(inorder[inIdx+1:], postorder[inIdx:len(postorder)-1])
	}
	return root
}

// @lc code=end

