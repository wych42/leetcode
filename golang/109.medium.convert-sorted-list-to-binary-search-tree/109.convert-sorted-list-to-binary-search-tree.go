/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (50.15%)
 * Likes:    2694
 * Dislikes: 93
 * Total Accepted:    282.4K
 * Total Submissions: 562.2K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given the head of a singly linked list where elements are sorted in
 * ascending order, convert it to a height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [-10,-3,0,5,9]
 * Output: [0,-3,9,-10,null,5]
 * Explanation: One possible answer is [0,-3,9,-10,null,5], which represents
 * the shown height balanced BST.
 *
 *
 * Example 2:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [0]
 * Output: [0]
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1,3]
 * Output: [3,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in head is in the range [0, 2 * 10^4].
 * -10^5 <= Node.val <= 10^5
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
var curr *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	curr = head
	count := 1
	for head.Next != nil {
		count++
		head = head.Next
	}
	return bstRecursive(count)
}

// 一个中序遍历
func bstRecursive(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	// 用一半节点递归构建左子树
	left := bstRecursive(n / 2)

	// 构建root
	root := &TreeNode{Val: curr.Val}
	root.Left = left
	curr = curr.Next

	// 构建右子树
	root.Right = bstRecursive(n - n/2 - 1)
	return root
}

// @lc code=end

