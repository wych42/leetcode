/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (43.33%)
 * Likes:    1911
 * Dislikes: 380
 * Total Accepted:    252.5K
 * Total Submissions: 579.7K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given the head of a linked list and a value x, partition it such that all
 * nodes less than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,4,3,2,5,2], x = 3
 * Output: [1,2,2,4,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [2,1], x = 2
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 200].
 * -100 <= Node.val <= 100
 * -200 <= x <= 200
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
 从头开始，如果一个节点A.Val >= x，那么后续所有的节点里，<x 的都应该放在 A 之前
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := &ListNode{Next: head}
	curr := res

	var ph, pt *ListNode
	for curr.Next != nil {
		if curr.Next.Val >= x {
			if ph == nil {
				ph = curr // 后面所有比 x 小的节点，都放在这个后面
				pt = curr.Next
			}
			curr = curr.Next
			continue
		}
		if curr.Next.Val < x && ph == nil {
			// 在 x 的左边的点不用管
			curr = curr.Next
			continue
		}
		// curr.Next.Val < x && ph != nil
		ph.Next = curr.Next
		curr.Next = curr.Next.Next
		ph = ph.Next
	}
	if ph != nil {
		ph.Next = pt
	}
	return res.Next
}

// @lc code=end

