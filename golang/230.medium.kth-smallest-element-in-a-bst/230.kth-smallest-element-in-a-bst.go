/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (62.51%)
 * Likes:    3575
 * Dislikes: 82
 * Total Accepted:    507.7K
 * Total Submissions: 810.3K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * Given the root of a binary search tree, and an integer k, return the k^th
 * (1-indexed) smallest element in the tree.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,1,4,null,2], k = 1
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is n.
 * 1 <= k <= n <= 10^4
 * 0 <= Node.val <= 10^4
 *
 *
 *
 * Follow up: If the BST is modified often (i.e., we can do insert and delete
 * operations) and you need to find the kth smallest frequently, how would you
 * optimize?
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

var idx int
var res int
var stop bool

func kthSmallest(root *TreeNode, k int) int {
	idx = k
	stop = false
	res = -1 << 31
	traversal(root)
	return res
}

func traversal(root *TreeNode) {
	if root == nil || stop {
		return
	}
	fmt.Println(root.Val, idx)
	traversal(root.Left)
	if idx == 1 && !stop {
		res = root.Val
		stop = true
		return
	}
	idx--
	traversal(root.Right)
}

// @lc code=end

