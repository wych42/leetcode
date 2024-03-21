/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (42.61%)
 * Likes:    6725
 * Dislikes: 349
 * Total Accepted:    839.9K
 * Total Submissions: 2M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * You are given an array of k linked-lists lists, each linked-list is sorted
 * in ascending order.
 *
 * Merge all the linked-lists into one sorted linked-list and return it.
 *
 *
 * Example 1:
 *
 *
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * merging them into one sorted list:
 * 1->1->2->3->4->4->5->6
 *
 *
 * Example 2:
 *
 *
 * Input: lists = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: lists = [[]]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] is sorted in ascending order.
 * The sum of lists[i].length won't exceed 10^4.
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
const max = 1 << 31

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	head := &ListNode{}
	curr := head

	pm := make(map[int]*ListNode)
	for idx, l := range lists {
		if l == nil {
			continue
		}
		pm[idx] = l
	}

	for len(pm) > 1 {
		min := max
		minIdx := -1
		for idx, l := range pm {
			if l != nil && l.Val < min {
				min = l.Val
				minIdx = idx
			}
		}
		if minIdx == -1 { // no value
			break
		}
		curr.Next = pm[minIdx]
		curr = curr.Next
		if pm[minIdx].Next == nil {
			delete(pm, minIdx)
		} else {
			pm[minIdx] = pm[minIdx].Next
		}
	}
	for _, l := range pm {
		curr.Next = l
	}
	return head.Next
}

// @lc code=end

