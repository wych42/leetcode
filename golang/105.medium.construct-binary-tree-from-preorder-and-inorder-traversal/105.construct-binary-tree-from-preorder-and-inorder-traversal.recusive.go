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
/*
preorder: root, left, rigth. 第一个元素是 Root
inorder: left, root, rigth. 第一个元素是最左节点

从 preorder 中不断的生成 node 节点，直到节点的值等于 inorder 中的值

*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	stack := make([]*TreeNode, 0)
	nodeSet := make(map[int]struct{})

	var root *TreeNode
	pre, in := 0, 0
	for pre < len(preorder) {
		var node *TreeNode
		for {
			node = &TreeNode{Val: preorder[pre]}
			if root == nil { // 第一个元素是 root
				root = node
			}
			if len(stack) > 0 {
				topNode := stack[len(stack)-1]
				if _, ok := nodeSet[topNode.Val]; ok {
					delete(nodeSet, topNode.Val)
					stack[len(stack)-1].Right = node
					stack = stack[:len(stack)-1]
				} else {
					stack[len(stack)-1].Left = node
				}
			}
			stack = append(stack, node)
			if pre >= len(preorder) || preorder[pre] == inorder[in] { // 元素就是最左节点
				pre++
				break
			}
			pre++
		}
		node = nil
		for len(stack) > 0 && in < len(inorder) && stack[len(stack)-1].Val == inorder[in] {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			in++
		}
		if node != nil {
			nodeSet[node.Val] = struct{}{}
			stack = append(stack, node)
		}
	}
	return root
}

// @lc code=end

