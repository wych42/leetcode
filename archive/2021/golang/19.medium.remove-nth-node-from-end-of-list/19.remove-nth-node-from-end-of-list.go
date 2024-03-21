/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (35.72%)
 * Likes:    4895
 * Dislikes: 290
 * Total Accepted:    812.1K
 * Total Submissions: 2.3M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, remove the n^th node from the end of the
 * list and return its head.
 *
 * Follow up: Could you do this in one pass?
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1], n = 1
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,2], n = 1
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is sz.
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for i := 1; i <= n; i++ {
		// fast 指针先往前走 n 步
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		// 同步往前走，走到最后一个节点；
		// 有可能，在 fast 先走的时候，已经走到尽头 == nil 了，那就直接中止
		fast = fast.Next
		slow = slow.Next
	}
	if slow == head && fast == nil {
		// slow 没动，fast 就已经走到头了，说明删除的是 head 自己
		return head.Next
	}
	// 否则删除的就是 slow.Next
	slow.Next = slow.Next.Next
	return head
}

// @lc code=end

