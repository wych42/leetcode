/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 *
 * https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (35.43%)
 * Likes:    5175
 * Dislikes: 375
 * Total Accepted:    471.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3]'
 *
 * A path in a binary tree is a sequence of nodes where each pair of adjacent
 * nodes in the sequence has an edge connecting them. A node can only appear in
 * the sequence at most once. Note that the path does not need to pass through
 * the root.
 *
 * The path sum of a path is the sum of the node's values in the path.
 *
 * Given the root of a binary tree, return the maximum path sum of any path.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3]
 * Output: 6
 * Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 =
 * 6.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [-10,9,20,null,null,15,7]
 * Output: 42
 * Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 +
 * 7 = 42.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 3 * 10^4].
 * -1000 <= Node.val <= 1000
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

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// m 是保存最大值的临时变量
	m := -1 << 31
	maxSubTree(root, &m)
	return m
}

func maxSubTree(root *TreeNode, m *int) int {
	if root == nil {
		// root是叶子节点时，计算左右子树的最大值会走到这里
		return 0
	}
	// 先找到左子树的最大值
	l := maxSubTree(root.Left, m)
	// 再找到右子树的最大值
	r := maxSubTree(root.Right, m)
	// 比较 root 值，左侧最大值+root，右侧最大值+root
	//      root
	//     .    .
	//   l        r
	// 此处保留下来的最大值，符合相邻节点之间有边连接的条件
	// 即单臂的路径: root, l-root, r-root
	// **这里不需要再单独比较 l, r的值，因为他们中较大的那个已经在 *m 里
	cm := max(root.Val, max(l+root.Val, r+root.Val))
	// 比较 cm 和 l+r+root.Val，即三节点组成的小子树的值: l-root-r
	// 将较大的存到 max value 指针里
	*m = max(*m, max(cm, l+r+root.Val))
	return cm
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
·