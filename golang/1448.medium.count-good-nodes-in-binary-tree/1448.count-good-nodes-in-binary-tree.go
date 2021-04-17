/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
 *
 * https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/
 *
 * algorithms
 * Medium (70.67%)
 * Likes:    664
 * Dislikes: 32
 * Total Accepted:    41.4K
 * Total Submissions: 58K
 * Testcase Example:  '[3,1,4,3,null,1,5]'
 *
 * Given a binary tree root, a node X in the tree is named good if in the path
 * from root to X there are no nodes with a value greater than X.
 *
 * Return the number of good nodes in the binary tree.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: root = [3,1,4,3,null,1,5]
 * Output: 4
 * Explanation: Nodes in blue are good.
 * Root Node (3) is always a good node.
 * Node 4 -> (3,4) is the maximum value in the path starting from the root.
 * Node 5 -> (3,4,5) is the maximum value in the path
 * Node 3 -> (3,1,3) is the maximum value in the path.
 *
 * Example 2:
 *
 *
 *
 *
 * Input: root = [3,3,null,4,2]
 * Output: 3
 * Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than
 * it.
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: 1
 * Explanation: Root is considered as good.
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the binary tree is in the range [1, 10^5].
 * Each node's value is between [-10^4, 10^4].
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
const min = -1 << 31

func goodNodes(root *TreeNode) int {
	/*
		从root到当前节点，其实是一个前序遍历
		在做前序遍历时，记住最大值，如果当前节点大于最大值，那么就是个好节点
		遍历完左边之后，再遍历右边，此时要记住遍历左边之前的最大值
		dfs
	*/
	ans := 0
	max := min
	traversal(root, max, &ans)
	return ans
}

func traversal(root *TreeNode, max int, count *int) {
	// capture goal
	if root == nil {
		return
	}
	if root.Val >= max {
		(*count)++
		max = root.Val
	}
	traversal(root.Left, max, count)
	traversal(root.Right, max, count)
}

// @lc code=end

