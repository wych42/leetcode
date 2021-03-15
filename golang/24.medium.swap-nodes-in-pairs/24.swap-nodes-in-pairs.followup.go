/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (53.00%)
 * Likes:    3406
 * Dislikes: 207
 * Total Accepted:    584.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4]
 * Output: [2,1,4,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1]
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 100].
 * 0 <= Node.val <= 100
 *
 *
 *
 * Follow up: Can you solve the problem without modifying the values in the
 * list's nodes? (i.e., Only nodes themselves may be changed.)
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	curr := head
	head = head.Next

	// 每次要交换 curr 和 curr.Next, 同时把 交换后的 curr 暂存到 prev 里
	for curr != nil && curr.Next != nil {
		t := curr.Next
		if prev != nil {
			// prev 指向上一次交换后，后面那个值，prev.Next = curr.Next
			prev.Next = t
			// prev 暂存 curr
			prev = curr
		} else {
			prev = curr
		}
		// curr 被交换后，Next 应该指向 curr.Next 的下一个节点
		curr.Next = curr.Next.Next

		// 实际交换: t 是原本的 curr.Next, 此时交换了， 应该指向 curr
		t.Next = curr
		curr = curr.Next
	}

	return head
}

// @lc code=end

