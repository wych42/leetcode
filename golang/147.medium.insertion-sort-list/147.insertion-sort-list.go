/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 *
 * https://leetcode.com/problems/insertion-sort-list/description/
 *
 * algorithms
 * Medium (44.33%)
 * Likes:    980
 * Dislikes: 679
 * Total Accepted:    227.2K
 * Total Submissions: 510.2K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Given the head of a singly linked list, sort the list using insertion sort,
 * and return the sorted list's head.
 *
 * The steps of the insertion sort algorithm:
 *
 *
 * Insertion sort iterates, consuming one input element each repetition and
 * growing a sorted output list.
 * At each iteration, insertion sort removes one element from the input data,
 * finds the location it belongs within the sorted list and inserts it
 * there.
 * It repeats until no input elements remain.
 *
 *
 * The following is a graphical example of the insertion sort algorithm. The
 * partially sorted list (black) initially contains only the first element in
 * the list. One element (red) is removed from the input data and inserted
 * in-place into the sorted list with each iteration.
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
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [1, 5000].
 * -5000 <= Node.val <= 5000
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
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	prev := dummy // 每次执行，都重置 prev 指向 dummy
	curr := head  // iter curr
	var next *ListNode
	for curr != nil {
		next = curr.Next
		for prev.Next != nil && prev.Next.Val < curr.Val {
			// 找到插入位置
			prev = prev.Next
		}
		// 把 curr 插到 prev 和 prev.Next 之间
		curr.Next = prev.Next
		prev.Next = curr
		prev = dummy
		curr = next
	}
	return dummy.Next
}

// @lc code=end

