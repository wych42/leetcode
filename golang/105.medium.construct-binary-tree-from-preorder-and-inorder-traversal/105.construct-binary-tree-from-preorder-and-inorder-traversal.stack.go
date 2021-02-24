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
	// 用 stack 存放已经访问过的 preorder 节点
	stack := make([]*TreeNode, 0)

	// 生成 root，入栈
	root := &TreeNode{Val: preorder[0]}
	preIdx++
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		/*
			pre: root-left-right
			in:  left-root-right
			栈顶的节点，与 inorder 中的当前节点不相等，说明 preorder 还在向左的过程中。下一个节点是 top 的左子节点。
			如果相等，把 inorder 指针右移，stack 出栈, 并继续判断下一个，如果下一个不想等，了，说明下一个 preorder 节点是出栈节点的右子节点
		*/
		if top.Val != inorder[inIdx] {
			// 出栈的 top 节点在 inorder 中还没遇到对应的值
			// 那么下一个 preorder 节点是 top 的左子节点
			node := &TreeNode{Val: preorder[preIdx]}
			preIdx++
			top.Left = node
			stack = append(stack, node)
		} else {
			stack = stack[:len(stack)-1]
			inIdx++

			// inorder 中的节点都访问过了，结束循环
			if inIdx == len(inorder) {
				break
			}

			// 栈里还有节点, 在 inorder 中没访问过（没找到对应的）,一个个比较，把已经访问过了出栈删掉
			if len(stack) > 0 && stack[len(stack)-1].Val == inorder[inIdx] {
				continue
			}

			// 剩下的栈顶的节点，inorder 中还没有找到对应的值(在后面), 说明栈顶节点和上一个出栈的 top 节点之间，还有inorder中没访问的右子树节点
			// preorder 中的下一个节点，就是top的右子节点
			node := &TreeNode{Val: preorder[preIdx]}
			preIdx++
			top.Right = node
			stack = append(stack, node)
		}

	}
	return root
}

// @lc code=end

