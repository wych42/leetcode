/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (42.47%)
 * Likes:    2867
 * Dislikes: 200
 * Total Accepted:    221.4K
 * Total Submissions: 520.3K
 * Testcase Example:  '3'
 *
 * Given an integer n, return all the structurally unique BST's (binary search
 * trees), which has exactly n nodes of unique values from 1 to n. Return the
 * answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3
 * Output:
 * [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: [[1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 8
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
func generateTrees(n int) []*TreeNode {
	/*
		节点取值区间是 [1:n]
		每一个节点都可以成为 root. 取 i 为当前的 root
		[1:i) 为左子树；(i:n] 为右子树
		递归
	*/
	return generateTreesByRange(1, n)

}

func generateTreesByRange(s, e int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if s > e {
		trees = append(trees, nil)
		return trees
	}
	if s == e {
		trees = append(trees, &TreeNode{Val: s})
		return trees
	}
	for i := s; i <= e; i++ {
		lTrees := generateTreesByRange(s, i-1)
		rTrees := generateTreesByRange(i+1, e)
		for _, lt := range lTrees {
			for _, rt := range rTrees {
				root := &TreeNode{Val: i}
				root.Left = lt
				root.Right = rt
				trees = append(trees, root)
			}
		}
	}
	return trees
}

// @lc code=end

