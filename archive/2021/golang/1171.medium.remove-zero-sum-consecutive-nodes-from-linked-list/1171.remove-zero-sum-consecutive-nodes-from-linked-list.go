/*
 * @lc app=leetcode id=1171 lang=golang
 *
 * [1171] Remove Zero Sum Consecutive Nodes from Linked List
 *
 * https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/description/
 *
 * algorithms
 * Medium (41.43%)
 * Likes:    822
 * Dislikes: 52
 * Total Accepted:    24K
 * Total Submissions: 57.7K
 * Testcase Example:  '[1,2,-3,3,1]'
 *
 * Given the head of a linked list, we repeatedly delete consecutive sequences
 * of nodes that sum to 0 until there are no such sequences.
 *
 * After doing so, return the head of the final linked list.  You may return
 * any such answer.
 *
 *
 * (Note that in the examples below, all sequences are serializations of
 * ListNode objects.)
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,-3,3,1]
 * Output: [3,1]
 * Note: The answer [1,2,1] would also be accepted.
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,3,-3,4]
 * Output: [1,2,4]
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,2,3,-3,-2]
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The given linked list will contain between 1 and 1000 nodes.
 * Each node in the linked list has -1000 <= node.val <= 1000.
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
func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy // 要记住前一个

	for head != nil {
		sum := head.Val
		curr := head.Next
		for curr != nil && sum != 0 { // 暴力回溯
			sum += curr.Val
			curr = curr.Next
		}
		if sum == 0 {
			prev.Next = curr // head .. curr.Prev 之和是0，所以 prev.Next 指向 curr
			head = curr
		} else {
			prev = head
			head = head.Next
		}
	}
	return dummy.Next
}

// @lc code=end

