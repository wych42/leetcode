/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (40.56%)
 * Likes:    2963
 * Dislikes: 144
 * Total Accepted:    312.3K
 * Total Submissions: 764.6K
 * Testcase Example:  '[1,2,3,4]'
 *
 * You are given the head of a singly linked-list. The list can be represented
 * as:
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * Reorder the list to be on the following form:
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 *
 * You may not modify the values in the list's nodes. Only nodes themselves may
 * be changed.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4]
 * Output: [1,4,2,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,3,4,5]
 * Output: [1,5,2,4,3]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [1, 5 * 10^4].
 * 1 <= Node.val <= 1000
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
func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// 用双指针，找到中间节点, 在 slow 之后的节点都是要执行插入操作的
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 把 slow 之后的节点记录下来
	nodes := make([]*ListNode, 0)
	for slow != nil {
		nodes = append(nodes, slow)
		slow = slow.Next
	}
	curr := head
	// 反向挨个插入
	for i := len(nodes) - 1; i >= 1; i-- {
		n := nodes[i]
		nodes[i-1].Next = nil
		n.Next = curr.Next
		curr.Next = n
		curr = curr.Next.Next
	}
	return
}

// @lc code=end

