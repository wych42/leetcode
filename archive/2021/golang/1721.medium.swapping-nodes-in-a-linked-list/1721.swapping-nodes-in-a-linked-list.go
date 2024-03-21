/*
 * @lc app=leetcode id=1721 lang=golang
 *
 * [1721] Swapping Nodes in a Linked List
 *
 * https://leetcode.com/problems/swapping-nodes-in-a-linked-list/description/
 *
 * algorithms
 * Medium (64.56%)
 * Likes:    446
 * Dislikes: 33
 * Total Accepted:    42.2K
 * Total Submissions: 62.1K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * You are given the head of a linked list, and an integer k.
 *
 * Return the head of the linked list after swapping the values of the k^th
 * node from the beginning and the k^th node from the end (the list is
 * 1-indexed).
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [1,4,3,2,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
 * Output: [7,9,6,6,8,7,3,0,9,5]
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1], k = 1
 * Output: [1]
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1,2], k = 1
 * Output: [2,1]
 *
 *
 * Example 5:
 *
 *
 * Input: head = [1,2,3], k = 2
 * Output: [1,2,3]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is n.
 * 1 <= k <= n <= 10^5
 * 0 <= Node.val <= 100
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
func swapNodes(head *ListNode, k int) *ListNode {
	fast := head
	for c := 1; c < k && fast != nil; c++ {
		fast = fast.Next // kth node
	}
	prev := fast

	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Val, prev.Val = prev.Val, slow.Val
	return head
}

// @lc code=end

