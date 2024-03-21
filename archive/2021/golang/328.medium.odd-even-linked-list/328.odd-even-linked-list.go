/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 *
 * https://leetcode.com/problems/odd-even-linked-list/description/
 *
 * algorithms
 * Medium (57.03%)
 * Likes:    2929
 * Dislikes: 335
 * Total Accepted:    359.7K
 * Total Submissions: 629.1K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given the head of a singly linked list, group all the nodes with odd indices
 * together followed by the nodes with even indices, and return the reordered
 * list.
 *
 * The first node is considered odd, and the second node is even, and so on.
 *
 * Note that the relative order inside both the even and odd groups should
 * remain as it was in the input.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5]
 * Output: [1,3,5,2,4]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [2,1,3,5,6,4,7]
 * Output: [2,3,6,7,1,5,4]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the linked list is in the range [0, 10^4].
 * -10^6 <= Node.val <= 10^6
 *
 *
 *
 * Follow up: Could you solve it in O(1) space complexity and O(nodes) time
 * complexity?
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	var eh, ehCurr *ListNode // even header && even current
	var prev *ListNode       // prev of current
	for curr != nil && curr.Next != nil {
		prev = curr
		even := curr.Next          // 当前的下一个是 even
		curr.Next = curr.Next.Next // 跳到下一个 odd
		curr = curr.Next
		if eh != nil { // even 加到even 列表里
			ehCurr.Next = even
			ehCurr = ehCurr.Next
		} else {
			eh = even
			ehCurr = eh
		}
	}
	ehCurr.Next = nil // 注意回路
	if curr != nil {  // 一共有奇数
		curr.Next = eh
	} else { // 一共有偶数个
		prev.Next = eh
	}
	return head
}

// @lc code=end

