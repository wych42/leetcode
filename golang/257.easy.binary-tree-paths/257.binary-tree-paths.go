/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (46.36%)
 * Likes:    2404
 * Dislikes: 127
 * Total Accepted:    379.6K
 * Total Submissions: 707.3K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * Output: ["1->2->5", "1->3"]
 *
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	l := binaryTreePaths(root.Left)
	r := binaryTreePaths(root.Right)
	res := make([]string, 0)
	rv := strconv.Itoa(root.Val) + "->"
	for _, s := range l {
		res = append(res, rv+s)
	}
	for _, s := range r {
		res = append(res, rv+s)
	}
	return res
}

// @lc code=end

