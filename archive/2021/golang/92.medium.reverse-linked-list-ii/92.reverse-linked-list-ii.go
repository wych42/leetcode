/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (40.52%)
 * Likes:    3431
 * Dislikes: 178
 * Total Accepted:    336.6K
 * Total Submissions: 826.8K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Given the head of a singly linked list and two integers left and right where
 * left <= right, reverse the nodes of the list from position left to position
 * right, and return the reversed list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], left = 2, right = 4
 * Output: [1,4,3,2,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [5], left = 1, right = 1
 * Output: [5]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is n.
 * 1 <= n <= 500
 * -500 <= Node.val <= 500
 * 1 <= left <= right <= n
 *
 *
 *
 * Follow up: Could you do it in one pass?
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	count := 0
	for count < left-1 {
		curr = curr.Next
		count++
	}
	lprev := curr    // left node 的前一个
	curr = curr.Next // left node
	count++
	// lhead := curr
	for count < right {
		if curr.Next == nil {
			break
		}
		// 每一个node都插到 lprev 后面
		temp := curr.Next
		curr.Next = temp.Next

		temp.Next = lprev.Next
		lprev.Next = temp

		count++
	}
	return dummy.Next
}

// @lc code=end

