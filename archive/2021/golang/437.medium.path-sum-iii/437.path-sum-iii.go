/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 *
 * https://leetcode.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Medium (48.10%)
 * Likes:    4873
 * Dislikes: 319
 * Total Accepted:    258.8K
 * Total Submissions: 537.2K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go
 * downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range
 * -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 * ⁠     10
 * ⁠    /  \
 * ⁠   5   -3
 * ⁠  / \    \
 * ⁠ 3   2   11
 * ⁠/ \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
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
 暴力解法：
 计算从root出发，满足条件的 path 数量
 再把每一个节点当作 root，挨个计算
*/
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	l := pathSum(root.Left, sum)
	r := pathSum(root.Right, sum)
	rt := rootPathSum(root, sum)
	return l + r + rt
}

func rootPathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	l := rootPathSum(root.Left, sum-root.Val)
	r := rootPathSum(root.Right, sum-root.Val)
	if root.Val == sum {
		return l + r + 1
	}
	return l + r
}

// @lc code=end

