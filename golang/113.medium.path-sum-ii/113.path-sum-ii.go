/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (49.03%)
 * Likes:    2555
 * Dislikes: 85
 * Total Accepted:    395.8K
 * Total Submissions: 805.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given the root of a binary tree and an integer targetSum, return all
 * root-to-leaf paths where each path's sum equals targetSum.
 *
 * A leaf is a node with no children.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * Output: [[5,4,11,2],[5,8,4,5]]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,2,3], targetSum = 5
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1,2], targetSum = 0
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 5000].
 * -1000 <= Node.val <= 1000
 * -1000 <= targetSum <= 1000
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	paths := make([][]int, 0)
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		paths = append(paths, []int{root.Val})
		return paths
	}
	if root.Left != nil {
		lps := pathSum(root.Left, targetSum)
		for _, p := range lps {
			if len(p) == 0 {
				continue
			}
			paths = append(paths, append([]int{root.Val}, p...))
		}
	}
	if root.Right != nil {
		rps := pathSum(root.Right, targetSum)
		for _, p := range rps {
			if len(p) == 0 {
				continue
			}
			paths = append(paths, append([]int{root.Val}, p...))
		}
	}
	return paths
}

// @lc code=end

