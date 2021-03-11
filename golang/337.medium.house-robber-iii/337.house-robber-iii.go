/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 *
 * https://leetcode.com/problems/house-robber-iii/description/
 *
 * algorithms
 * Medium (51.83%)
 * Likes:    3815
 * Dislikes: 65
 * Total Accepted:    203.7K
 * Total Submissions: 392.2K
 * Testcase Example:  '[3,2,3,null,3,null,1]'
 *
 * The thief has found himself a new place for his thievery again. There is
 * only one entrance to this area, called root.
 *
 * Besides the root, each house has one and only one parent house. After a
 * tour, the smart thief realized that all houses in this place form a binary
 * tree. It will automatically contact the police if two directly-linked houses
 * were broken into on the same night.
 *
 * Given the root of the binary tree, return the maximum amount of money the
 * thief can rob without alerting the police.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,2,3,null,3,null,1]
 * Output: 7
 * Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [3,4,5,1,3,null,1]
 * Output: 9
 * Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * 0 <= Node.val <= 10^4
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
对于一个 root，有两个选项：
- root.Val + 左边没抢的 + 右边没抢的
- max(左边抢了的，左边没抢的) + max(右边没抢的，右边抢了的); 因为不涉及 root，所以这里没问题

一个后序遍历。

问题是叶子节点：
- 只有一个节点时，就抢，所以叶子节点默认设置为 robbed, 此时 notRobbed 为 0.
*/
type pair struct {
	robbed    int
	notRobbed int
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	p := dfs(root)
	return max(p.robbed, p.notRobbed)
}

func dfs(root *TreeNode) pair {
	if root == nil {
		return pair{}
	}
	if root.Left == nil && root.Right == nil {
		return pair{robbed: root.Val}
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	p := pair{}
	p.robbed = root.Val + l.notRobbed + r.notRobbed
	p.notRobbed = max(l.robbed, l.notRobbed) + max(r.robbed, r.notRobbed)
	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

