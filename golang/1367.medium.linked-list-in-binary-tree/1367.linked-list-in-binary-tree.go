/*
 * @lc app=leetcode id=1367 lang=golang
 *
 * [1367] Linked List in Binary Tree
 *
 * https://leetcode.com/problems/linked-list-in-binary-tree/description/
 *
 * algorithms
 * Medium (41.21%)
 * Likes:    769
 * Dislikes: 28
 * Total Accepted:    28.5K
 * Total Submissions: 69.5K
 * Testcase Example:  '[4,2,8]\n[1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]'
 *
 * Given a binary tree root and a linked list with head as the first node.
 *
 * Return True if all the elements in the linked list starting from the head
 * correspond to some downward path connected in the binary tree otherwise
 * return False.
 *
 * In this context downward path means a path that starts at some node and goes
 * downwards.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: head = [4,2,8], root =
 * [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
 * Output: true
 * Explanation: Nodes in blue form a subpath in the binary Tree.
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: head = [1,4,2,6], root =
 * [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,4,2,6,8], root =
 * [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
 * Output: false
 * Explanation: There is no path in the binary tree that contains all the
 * elements of the linked list from head.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree will be in the range [1, 2500].
 * The number of nodes in the list will be in the range [1, 100].
 * 1 <= Node.val <= 100 for each node in the linked list and binary tree.
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	// 对树的每一个节点，都用一个前序遍历来判断
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			if isSubPathFromNode(head, node) {
				return true
			}
			queue = append(queue, root.Left, root.Right)
		}
	}
	return false
}
func isSubPathFromNode(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil || root.Val != head.Val {
		return false
	}
	return isSubPathFromNode(head.Next, root.Left) || isSubPathFromNode(head.Next, root.Right)
}

// @lc code=end

