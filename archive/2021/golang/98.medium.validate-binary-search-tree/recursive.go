/*
 * @lc app=leetcode id=98 lang=golang
 */

// @lc code=start
// 递归
/*
root 是左子树的 maxNode，是右子树的 minNode
*/
func isValidBST(root *TreeNode) bool {
	return isValidBSTRecusive(root, nil, nil)
}

func isValidBSTRecusive(root *TreeNode, maxNode *TreeNode, minNode *TreeNode) bool {
	if root == nil {
		return true
	}
	if maxNode != nil && maxNode.Val <= root.Val {
		return false
	}
	if minNode != nil && minNode.Val >= root.Val {
		return false
	}

	return isValidBSTRecusive(root.Left, root, minNode) && isValidBSTRecusive(root.Right, maxNode, root)
}

// @lc code=end

