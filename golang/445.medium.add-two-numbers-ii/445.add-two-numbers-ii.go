/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 *
 * https://leetcode.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (56.22%)
 * Likes:    2236
 * Dislikes: 198
 * Total Accepted:    228.3K
 * Total Submissions: 404.5K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The most significant digit comes first and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the
 * lists is not allowed.
 *
 *
 *
 * Example:
 *
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	result := &ListNode{}
	var e int
	for l1 != nil || l2 != nil {
		v := e
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if v >= 10 {
			e = 1
			v = v - 10
		} else {
			e = 0
		}
		node := &ListNode{Val: v}
		node.Next = result.Next
		result.Next = node
	}
	if e != 0 {
		node := &ListNode{Val: e}
		node.Next = result.Next
		result.Next = node
	}
	return result.Next
}

func reverse(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = next
	}
	return dummy.Next
}

// @lc code=end

