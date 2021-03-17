/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (39.20%)
 * Likes:    2743
 * Dislikes: 122
 * Total Accepted:    319.7K
 * Total Submissions: 812.2K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given the head of a sorted linked list, delete all nodes that have duplicate
 * numbers, leaving only distinct numbers from the original list. Return the
 * linked list sorted as well.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,3,4,4,5]
 * Output: [1,2,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,1,1,2,3]
 * Output: [2,3]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 300].
 * -100 <= Node.val <= 100
 * The list is guaranteed to be sorted in ascending order.
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
/*
整体思路：不从链表里删除节点，而是找能够加到链表末尾的节点

*/
func deleteDuplicates(head *ListNode) *ListNode {
	var res, tail *ListNode
	for head != nil {
		next := head // 暂存一下
		// head.Next 是空，说明到头了，head 可以附加到结果里
		// head.Next.Val != head.Val，head 也可以附加到结果里
		if head.Next == nil || head.Next.Val != head.Val {
			if tail != nil {
				// 把当前 tail 往前挪一个
				tail.Next = head
			}
			// 此时 head 节点是结果链表的最后一个
			tail = head

			// res 还没有设置过，res 就是 tail
			if res == nil {
				res = tail
			}
		}

		// head(也就是 next) 的下一个与当前值相等，就一直往后找，一直到不相等的节点
		for next != nil && next.Val == head.Val {
			next = next.Next
		}
		head = next
	}
	// 注意这个边界条件
	if prev != nil {
		prev.Next = nil
	}
	return res
}

// @lc code=end

