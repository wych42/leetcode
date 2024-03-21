/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (51.71%)
 * Likes:    4754
 * Dislikes: 123
 * Total Accepted:    461K
 * Total Submissions: 890.5K
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
// preorder 的第一个元素是 root，在 inorder 中，该元素的左右分别是左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 因为要在循环中找元素的位置，先生成两个 map
	inorderM := make(map[int]int)
	for i, e := range inorder {
		inorderM[e] = i
	}
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, inorderM)
}

// 递归方法
// preorder 和 inorder 都传递原列表
// 传每次调用时，应该包含 preorder 中的元素区间
// 传每次调用时，还剩余的 inorder 的元素区间
// 传递 inorderM，方便快速的找元素
func build(preorder []int, inorder []int, preStart, preStop int, inStart, inStop int, inorderM map[int]int) *TreeNode {
	if preStart > preStop || inStart > inStop {
		return nil
	}
	root := &TreeNode{Val: preorder[preStart]}
	inIdx := inorderM[preorder[preStart]]
	leftCount := inIdx - inStart
	root.Left = build(
		preorder,
		inorder,
		preStart+1,         // 跳过已经被使用的元素
		preStart+leftCount, // 左子树包含 [preStart+1:preStart+leftCount]
		inStart,            // 左子树对应的 inorder 从 inStart 开始
		inIdx-1,            // 左子树对应 inorder 包含 [inStart, inIdx-1], 这里要跳过 inIdx
		inorderM,
	)
	root.Right = build(
		preorder,
		inorder,
		preStart+leftCount+1, // 跳过已经被使用的元素和左子树对应的元素
		preStop,              // 右子树包含 [preStart+leftCount+1:preStop]
		inIdx+1,              // 右子树对应的 inorder 从 inIdx+1 开始, 跳过 inIdx
		inStop,               // 右子树对应 inorder 包含 [inIdx+1, inStop]
		inorderM,
	)
	return root
}

// @lc code=end

