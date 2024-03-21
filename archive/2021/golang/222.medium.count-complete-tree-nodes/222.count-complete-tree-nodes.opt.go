/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 *
 * https://leetcode.com/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Medium (49.20%)
 * Likes:    2752
 * Dislikes: 250
 * Total Accepted:    282.7K
 * Total Submissions: 572.9K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * Given the root of a complete binary tree, return the number of the nodes in
 * the tree.
 *
 * According to Wikipedia, every level, except possibly the last, is completely
 * filled in a complete binary tree, and all nodes in the last level are as far
 * left as possible. It can have between 1 and 2^h nodes inclusive at the last
 * level h.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3,4,5,6]
 * Output: 6
 *
 *
 * Example 2:
 *
 *
 * Input: root = []
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 5 * 10^4].
 * 0 <= Node.val <= 5 * 10^4
 * The tree is guaranteed to be complete.
 *
 *
 *
 * Follow up: Traversing the tree to count the number of nodes in the tree is
 * an easy solution but with O(n) complexity. Could you find a faster algorithm?
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

var completeLevel int // 完全树的最大高度
var count int         // 从右边开始缺少的个数
var isComplete bool   // 标记目前已经确认是不是完全树
var reachRight bool   // 有没有遍历完成最右树

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	completeLevel = 0
	count = 0
	isComplete = true
	reachRight = false

	traversal(root, 0)
	//fmt.Println(completeLevel, isComplete, count)

	if isComplete {
		// 是完全树，不缺
		count = 0
		completeLevel += 1
	} else {
		completeLevel += 2
	}
	return 1<<completeLevel - 1 - count
}

func traversal(root *TreeNode, level int) {
	if root == nil {
		// 从右子树开始递归，root==ni的时候，说明已经到达或者到达过最右叶子节点了
		reachRight = true
		return
	}
	if !isComplete {
		return
	}
	if !reachRight {
		// 还没到最右叶子节点，目前为止还是完全树
		completeLevel = level
	}
	if root.Left == nil && root.Right == nil { // hit leaf
		// 从右子树开始递归，说明已经到达或者到达过最右叶子节点了
		reachRight = true
		if level > completeLevel {
			// 当前高度比完全树的高度大，说明不是完全树了
			isComplete = false
			return
		}
		// 假设不是完全树，这里缺两个
		count += 2
		return
	}
	if root.Right == nil { // mising
		// 确定不是完全树，这里缺一个
		isComplete = false
		count += 1
		return
	}

	// 先访问右边
	traversal(root.Right, level+1)
	traversal(root.Left, level+1)
}

// @lc code=end

