/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (31.72%)
 * Likes:    2166
 * Dislikes: 1135
 * Total Accepted:    361.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, rotate the list to the right by k
 * places.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [4,5,1,2,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [0,1,2], k = 4
 * Output: [2,0,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 500].
 * -100 <= Node.val <= 100
 * 0 <= k <= 2 * 10^9
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	count := 1
	fast := head
	// 找到长度； fast 在末尾
	for fast.Next != nil {
		fast = fast.Next
		count++
	}
	// 取余数, 避免重复执行
	k = k % count
	if k == 0 { // 是 0 就直接返回了
		return head
	}

	// slow 挪到要翻转的前一个(新的 head 前一个)
	slow := head
	for i := 1; i < count-k; i++ {
		slow = slow.Next
	}
	fast.Next = head

	head = slow.Next
	slow.Next = nil
	return head
}

// @lc code=end

