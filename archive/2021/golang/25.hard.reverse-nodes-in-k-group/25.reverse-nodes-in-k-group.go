/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (44.75%)
 * Likes:    3490
 * Dislikes: 400
 * Total Accepted:    341.1K
 * Total Submissions: 752.3K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, reverse the nodes of a linked list k at a time and
 * return its modified list.
 *
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes, in the end, should remain as it is.
 *
 * Follow up:
 *
 *
 * Could you solve the problem in O(1) extra memory space?
 * You may not alter the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [2,1,4,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,3,4,5], k = 3
 * Output: [3,2,1,4,5]
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,2,3,4,5], k = 1
 * Output: [1,2,3,4,5]
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1], k = 1
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range sz.
 * 1 <= sz <= 5000
 * 0 <= Node.val <= 1000
 * 1 <= k <= sz
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{Next: head} // 记录结果
	dummy := res                 // 每一组翻转的头部
	tail := head                 // 每一组翻转的尾部

	fast := head
	for i := 1; i < k; i++ {
		if fast == nil { // 不够 k 个数
			return head
		}
		fast = fast.Next
	}
	slow := head
	// 现在 fast 在第k个，slow在第一个,大家一起往前走
	fastCnt := 0
	cnt := 0
	for slow != nil {
		if cnt < k { // 还不够k个，执行翻转
			// 只要下一个不是nil，fastCnt+1, 记下来下一组有几个数
			if fast != nil && fast.Next != nil {
				fast = fast.Next
				fastCnt++
			}
			// 把 slow 放到 dummy 后面，实现翻转
			// 原来 res(dummy) -> 1(head, tail) -> 2(slow) -> 3
			// 现在 res(dummy) -> 2(slow) -> 1(head, tail) -> 3
			temp := slow.Next
			slow.Next = dummy.Next
			dummy.Next = slow
			slow = temp
			// tail 的下一个也是 slow 的下一个，直接连上
			tail.Next = temp
			cnt++
		} else {
			// 一个周期结束了
			if fastCnt < k {
				// 如果剩下的不够k个，不处理，直接返回
				break
			}
			// 不然就重置计数器
			cnt = 0
			fastCnt = 0
			// dummy 指向 tail，即下一组的前一个
			dummy = tail
			// tail 指向下一组的第一个
			tail = slow
		}
	}
	return res.Next
}

// @lc code=end

