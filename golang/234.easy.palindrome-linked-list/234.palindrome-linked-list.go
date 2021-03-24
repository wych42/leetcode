/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (40.41%)
 * Likes:    4854
 * Dislikes: 430
 * Total Accepted:    584.6K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,2,1]'
 *
 * Given the head of a singly linked list, return true if it is a
 * palindrome.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,2,1]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [1, 10^5].
 * 0 <= Node.val <= 9
 *
 *
 *
 * Follow up: Could you do it in O(n) time and O(1) space?
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	dummy := &ListNode{}
	var temp *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		temp = slow.Next
		slow.Next = dummy.Next
		dummy.Next = slow
		slow = temp
	}
	l2 := slow.Next
	if fast != nil {
		slow.Next = dummy.Next
		dummy.Next = slow
	}
	l1 := dummy.Next

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// @lc code=end

