/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 *
 * https://leetcode.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (46.19%)
 * Likes:    3998
 * Dislikes: 171
 * Total Accepted:    344.1K
 * Total Submissions: 740K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Given the head of a linked list, return the list after sorting it in
 * ascending order.
 *
 * Follow up: Can you sort the linked list in O(n logn) time and O(1) memory
 * (i.e. constant space)?
 *
 *
 * Example 1:
 *
 *
 * Input: head = [4,2,1,3]
 * Output: [1,2,3,4]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [-1,5,3,4,0]
 * Output: [-1,0,3,4,5]
 *
 *
 * Example 3:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 5 * 10^4].
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
// merge sort, 递归
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	temp := slow.Next
	slow.Next = nil
	return mergeSortList(sortList(head), sortList(temp))
}

func mergeSortList(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for a != nil && b != nil {
		if a.Val < b.Val {
			curr.Next = a
			a = a.Next
		} else {
			curr.Next = b
			b = b.Next
		}
		curr = curr.Next
	}
	if a != nil {
		curr.Next = a
	}
	if b != nil {
		curr.Next = b
	}
	return dummy.Next
}

// @lc code=end

