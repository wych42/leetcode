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
type pair struct {
	withRoot    int
	withoutRoot int
}

func rob(root *TreeNode) int {
	/*
	   子问题有两个：带root，和不带root
	   带root：就不能抢左右子节点，只能向下走
	   不带root：可以抢左右子节点，也可以跳过

	   用一个后序遍历，对每一个树，都找到它的左右子树 [带root，不带root]
	   那么这棵树：
	     带root：root+左不带root+右不带root
	     不带root：max(左边不带root，左边带root) + max(右边同理)

	   对于叶子节点：带root=root.Val, 不带 root = 0
	   对于空节点: [0, 0]
	*/

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	p := dfs(root)
	return max(p.withRoot, p.withoutRoot)
}

func dfs(root *TreeNode) pair {
	if root == nil {
		return pair{}
	}
	p := pair{withRoot: root.Val}
	p := pair{}
	l := dfs(root.Left)
	r := dfs(root.Right)
	p.withRoot = root.Val + l.withoutRoot + r.withoutRoot
	p.withoutRoot = max(l.withRoot, l.withoutRoot) + max(r.withRoot, r.withoutRoot)
	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

